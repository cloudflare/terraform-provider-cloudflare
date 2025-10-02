package config

import (
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/common"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// SplitObjectTransformer splits an object attribute into multiple attributes
// This is useful when an API changes from a nested object to flat attributes
//
// Example YAML configuration:
//   structural_changes:
//     - type: split_object
//       source: address
//       parameters:
//         attributes:
//           - street
//           - city
//           - state
//           - zip
//         prefix: addr_
//
// Transforms:
//   address = {
//     street = "123 Main St"
//     city = "San Francisco"
//     state = "CA"
//     zip = "94102"
//   }
//
// Into:
//   addr_street = "123 Main St"
//   addr_city = "San Francisco"
//   addr_state = "CA"
//   addr_zip = "94102"
func SplitObjectTransformer(source string, attributes []string, prefix string) common.TransformerFunc {
	return func(block *hclwrite.Block, ctx *common.TransformContext) error {
		if source == "" {
			return nil
		}
		
		body := block.Body()
		attr := body.GetAttribute(source)
		
		if attr == nil {
			return nil // Source attribute doesn't exist
		}
		
		// Parse the object expression
		expr := attr.Expr()
		tokens := expr.BuildTokens(nil)
		
		// Extract key-value pairs from the object
		values := parseObjectExpression(tokens)
		
		if len(values) == 0 {
			// Not an object, nothing to split
			return nil
		}
		
		// If attributes list is empty, extract all fields
		if len(attributes) == 0 {
			// Extract all fields from the object
			for key, value := range values {
				newAttrName := key
				if prefix != "" {
					newAttrName = prefix + key
				}
				body.SetAttributeRaw(newAttrName, value)
			}
			
			// Remove the original attribute after successful extraction
			body.RemoveAttribute(source)
		} else {
			// Extract only specified fields
			hasAnyField := false
			for _, attrName := range attributes {
				if value, exists := values[attrName]; exists {
					newAttrName := attrName
					if prefix != "" {
						newAttrName = prefix + attrName
					}
					body.SetAttributeRaw(newAttrName, value)
					hasAnyField = true
				}
			}
			
			// Only remove original if we extracted at least one field
			// or if all requested fields were attempted (partial extraction)
			if hasAnyField || len(attributes) > 0 {
				// Check if there are remaining fields
				remainingFields := make(map[string]hclwrite.Tokens)
				for key, value := range values {
					found := false
					for _, attrName := range attributes {
						if key == attrName {
							found = true
							break
						}
					}
					if !found {
						remainingFields[key] = value
					}
				}
				
				if len(remainingFields) > 0 {
					// Keep original with remaining fields
					newTokens := buildObjectTokens(remainingFields)
					body.SetAttributeRaw(source, newTokens)
				} else {
					// Remove original if all fields were extracted
					body.RemoveAttribute(source)
				}
			}
		}
		
		return nil
	}
}

// SplitObjectWithMapping splits an object and renames the attributes
//
// Example YAML:
//   structural_changes:
//     - type: split_object_with_mapping
//       source: config
//       parameters:
//         attribute_map:
//           host: hostname
//           port: service_port
//           proto: protocol
//
// Transforms:
//   config = {
//     host = "example.com"
//     port = 443
//     proto = "https"
//   }
//
// Into:
//   hostname = "example.com"
//   service_port = 443
//   protocol = "https"
func SplitObjectWithMapping(source string, attributeMap map[string]string) common.TransformerFunc {
	return func(block *hclwrite.Block, ctx *common.TransformContext) error {
		if source == "" {
			return nil
		}
		
		body := block.Body()
		attr := body.GetAttribute(source)
		
		if attr == nil {
			return nil
		}
		
		// Parse the object expression
		expr := attr.Expr()
		tokens := expr.BuildTokens(nil)
		
		values := parseObjectExpression(tokens)
		
		if len(values) == 0 {
			// Not an object, nothing to split
			return nil
		}
		
		// If mapping is empty or nil, extract all fields with original names
		if len(attributeMap) == 0 {
			for key, value := range values {
				body.SetAttributeRaw(key, value)
			}
			body.RemoveAttribute(source)
			return nil
		}
		
		// Extract and rename specified fields
		extractedAny := false
		for oldName, newName := range attributeMap {
			if value, exists := values[oldName]; exists {
				body.SetAttributeRaw(newName, value)
				extractedAny = true
			}
		}
		
		// Also extract unmapped fields with original names
		for key, value := range values {
			// Check if this field was mapped
			mapped := false
			for oldName := range attributeMap {
				if key == oldName {
					mapped = true
					break
				}
			}
			if !mapped {
				// This field wasn't mapped, extract with original name
				body.SetAttributeRaw(key, value)
			}
		}
		
		// Remove the original attribute if we extracted anything
		if extractedAny || len(values) > 0 {
			body.RemoveAttribute(source)
		}
		
		return nil
	}
}

// parseObjectExpression parses an object expression into key-value pairs
func parseObjectExpression(tokens hclwrite.Tokens) map[string]hclwrite.Tokens {
	values := make(map[string]hclwrite.Tokens)
	
	var currentKey string
	var currentValue hclwrite.Tokens
	var braceDepth, bracketDepth, parenDepth int
	var inObject, equalSeen, collectingValue bool
	
	for i, token := range tokens {
		switch token.Type {
		case hclsyntax.TokenOBrace:
			braceDepth++
			if braceDepth == 1 && i == 0 {
				// This is the opening brace of our object
				inObject = true
				continue
			}
			if collectingValue {
				currentValue = append(currentValue, token)
			}
			
		case hclsyntax.TokenCBrace:
			braceDepth--
			if braceDepth == 0 && inObject {
				// End of our object
				if currentKey != "" && len(currentValue) > 0 {
					values[currentKey] = trimTokenWhitespace(currentValue)
				}
				return values
			}
			if collectingValue {
				currentValue = append(currentValue, token)
			}
			
		case hclsyntax.TokenOBrack:
			bracketDepth++
			if collectingValue {
				currentValue = append(currentValue, token)
			}
			
		case hclsyntax.TokenCBrack:
			bracketDepth--
			if collectingValue {
				currentValue = append(currentValue, token)
			}
			
		case hclsyntax.TokenOParen:
			parenDepth++
			if collectingValue {
				currentValue = append(currentValue, token)
			}
			
		case hclsyntax.TokenCParen:
			parenDepth--
			if collectingValue {
				currentValue = append(currentValue, token)
			}
			
		case hclsyntax.TokenIdent:
			if inObject && braceDepth == 1 && !equalSeen && bracketDepth == 0 && parenDepth == 0 {
				// Check if this is actually a field name (not whitespace)
				text := string(token.Bytes)
				if strings.TrimSpace(text) != "" {
					currentKey = text
				}
			} else if collectingValue {
				currentValue = append(currentValue, token)
			}
			
		case hclsyntax.TokenEqual:
			if inObject && braceDepth == 1 && currentKey != "" && !equalSeen {
				equalSeen = true
				collectingValue = true
			} else if collectingValue {
				currentValue = append(currentValue, token)
			}
			
		case hclsyntax.TokenComma, hclsyntax.TokenNewline:
			if inObject && braceDepth == 1 && bracketDepth == 0 && parenDepth == 0 && currentKey != "" && equalSeen {
				// End of current field
				if len(currentValue) > 0 {
					values[currentKey] = trimWhitespace(currentValue)
				}
				currentKey = ""
				currentValue = nil
				equalSeen = false
				collectingValue = false
			} else if collectingValue {
				currentValue = append(currentValue, token)
			}
			
		default:
			if collectingValue {
				currentValue = append(currentValue, token)
			}
		}
	}
	
	// Handle last field if no trailing comma/newline
	if currentKey != "" && len(currentValue) > 0 {
		values[currentKey] = trimWhitespace(currentValue)
	}
	
	return values
}

// buildObjectTokens builds object tokens from a map
func buildObjectTokens(fields map[string]hclwrite.Tokens) hclwrite.Tokens {
	var tokens hclwrite.Tokens
	
	// Opening brace
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenOBrace,
		Bytes: []byte{'{'},
	})
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte{'\n'},
	})
	
	// Add each field
	for key, value := range fields {
		// Add indentation
		tokens = append(tokens, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte("    "),
		})
		
		// Add key
		tokens = append(tokens, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(key),
		})
		
		// Add equals
		tokens = append(tokens, &hclwrite.Token{
			Type:  hclsyntax.TokenEqual,
			Bytes: []byte{' ', '=', ' '},
		})
		
		// Add value
		tokens = append(tokens, value...)
		
		// Add newline
		tokens = append(tokens, &hclwrite.Token{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte{'\n'},
		})
	}
	
	// Closing brace
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("  "),
	})
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenCBrace,
		Bytes: []byte{'}'},
	})
	
	return tokens
}

// trimTokenWhitespace removes leading/trailing whitespace tokens
func trimTokenWhitespace(tokens hclwrite.Tokens) hclwrite.Tokens {
	// Trim leading whitespace/newlines
	start := 0
	for start < len(tokens) {
		if tokens[start].Type == hclsyntax.TokenNewline ||
		   (tokens[start].Type == hclsyntax.TokenIdent && strings.TrimSpace(string(tokens[start].Bytes)) == "") {
			start++
		} else {
			break
		}
	}
	
	// Trim trailing whitespace/newlines
	end := len(tokens)
	for end > start {
		if tokens[end-1].Type == hclsyntax.TokenNewline ||
		   (tokens[end-1].Type == hclsyntax.TokenIdent && strings.TrimSpace(string(tokens[end-1].Bytes)) == "") {
			end--
		} else {
			break
		}
	}
	
	if start >= end {
		return hclwrite.Tokens{}
	}
	
	return tokens[start:end]
}