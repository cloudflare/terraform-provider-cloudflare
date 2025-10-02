package structural

import (
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// ListToBlocksConverter converts a list attribute to multiple blocks
// This is the inverse of BlocksToListConverter
//
// Example YAML configuration:
//   lists_to_blocks:
//     - destinations  # Convert list attribute to multiple blocks
//
// Transforms:
//   destinations = [
//     { uri = "https://app1.com" },
//     { uri = "https://app2.com" }
//   ]
//
// Into:
//   destinations {
//     uri = "https://app1.com"
//   }
//   destinations {
//     uri = "https://app2.com"
//   }
func ListToBlocksConverter(attrName string) basic.TransformerFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		if attrName == "" {
			return nil
		}
		
		body := block.Body()
		attr := body.GetAttribute(attrName)
		
		if attr == nil {
			return nil // Attribute doesn't exist, nothing to convert
		}
		
		// Get the attribute expression
		expr := attr.Expr()
		tokens := expr.BuildTokens(nil)
		
		// Check if this is actually a list (starts with '[')
		isList := false
		for _, token := range tokens {
			if token.Type == hclsyntax.TokenOBrack {
				isList = true
				break
			}
			// Skip whitespace
			if token.Type != hclsyntax.TokenNewline && 
			   !(token.Type == hclsyntax.TokenIdent && strings.TrimSpace(string(token.Bytes)) == "") {
				// Non-whitespace before bracket means it's not a list
				break
			}
		}
		
		if !isList {
			// Not a list, nothing to convert
			return nil
		}
		
		// Parse the list structure
		items := parseListExpression(tokens)
		
		// Always remove the list attribute (even if empty)
		body.RemoveAttribute(attrName)
		
		// Create a block for each item
		for _, itemTokens := range items {
			newBlock := body.AppendNewBlock(attrName, nil)
			
			// Parse the item as an object
			parsed := parseObjectFromTokens(itemTokens)
			
			if len(parsed) > 0 {
				// Add all fields from the object to the block
				for fieldName, fieldTokens := range parsed {
					newBlock.Body().SetAttributeRaw(fieldName, fieldTokens)
				}
			} else {
				// If it's not an object, wrap it in a value attribute
				newBlock.Body().SetAttributeRaw("value", itemTokens)
			}
		}
		
		return nil
	}
}

// parseListExpression parses a list expression into individual items
func parseListExpression(tokens hclwrite.Tokens) []hclwrite.Tokens {
	var result []hclwrite.Tokens
	var currentItem hclwrite.Tokens
	var bracketDepth, braceDepth, parenDepth int
	var inList bool
	
	for i, token := range tokens {
		switch token.Type {
		case hclsyntax.TokenOBrack:
			bracketDepth++
			if bracketDepth == 1 && i == 0 {
				// Start of our list
				inList = true
				continue
			}
			if inList && bracketDepth > 1 {
				currentItem = append(currentItem, token)
			}
			
		case hclsyntax.TokenCBrack:
			bracketDepth--
			if bracketDepth == 0 && inList {
				// End of our list
				if len(currentItem) > 0 {
					result = append(result, trimWhitespace(currentItem))
				}
				return result
			}
			if inList && bracketDepth > 0 {
				currentItem = append(currentItem, token)
			}
			
		case hclsyntax.TokenOBrace:
			braceDepth++
			if inList {
				currentItem = append(currentItem, token)
			}
			
		case hclsyntax.TokenCBrace:
			braceDepth--
			if inList {
				currentItem = append(currentItem, token)
			}
			
		case hclsyntax.TokenOParen:
			parenDepth++
			if inList {
				currentItem = append(currentItem, token)
			}
			
		case hclsyntax.TokenCParen:
			parenDepth--
			if inList {
				currentItem = append(currentItem, token)
			}
			
		case hclsyntax.TokenComma:
			if inList && bracketDepth == 1 && braceDepth == 0 && parenDepth == 0 {
				// End of current item
				if len(currentItem) > 0 {
					result = append(result, trimWhitespace(currentItem))
					currentItem = nil
				}
			} else if inList {
				currentItem = append(currentItem, token)
			}
			
		default:
			if inList {
				currentItem = append(currentItem, token)
			}
		}
	}
	
	// Handle last item if no trailing comma
	if len(currentItem) > 0 {
		result = append(result, trimWhitespace(currentItem))
	}
	
	return result
}

// parseObjectFromTokens parses an object from tokens
func parseObjectFromTokens(tokens hclwrite.Tokens) map[string]hclwrite.Tokens {
	result := make(map[string]hclwrite.Tokens)
	
	var currentKey string
	var currentValue hclwrite.Tokens
	var braceDepth, bracketDepth, parenDepth int
	var inObject, equalSeen, collectingValue bool
	
	// First check if this is an object
	hasOpenBrace := false
	for _, token := range tokens {
		if token.Type == hclsyntax.TokenOBrace {
			hasOpenBrace = true
			break
		}
		// Skip whitespace
		if token.Type != hclsyntax.TokenNewline && 
		   !(token.Type == hclsyntax.TokenIdent && strings.TrimSpace(string(token.Bytes)) == "") {
			// Non-whitespace before brace means it's not an object
			break
		}
	}
	
	if !hasOpenBrace {
		return result
	}
	
	for _, token := range tokens {
		switch token.Type {
		case hclsyntax.TokenOBrace:
			braceDepth++
			if braceDepth == 1 && !inObject {
				// Start of our object
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
					result[currentKey] = trimWhitespace(currentValue)
				}
				return result
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
					result[currentKey] = trimWhitespace(currentValue)
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
		result[currentKey] = trimWhitespace(currentValue)
	}
	
	return result
}

// trimWhitespace removes leading/trailing whitespace tokens
func trimWhitespace(tokens hclwrite.Tokens) hclwrite.Tokens {
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

// ListToBlocksWithMapping converts a list attribute to blocks with field mapping
// This variant allows renaming fields during the conversion
//
// Example usage:
//   mapping := map[string]string{
//     "original_field": "new_field",
//     "another_field": "renamed_field"
//   }
//   transformer := ListToBlocksWithMapping("items", mapping)
//
// Transforms:
//   items = [
//     { original_field = "value1", another_field = "value2" },
//     { original_field = "value3", another_field = "value4" }
//   ]
//
// Into:
//   items {
//     new_field = "value1"
//     renamed_field = "value2"
//   }
//   items {
//     new_field = "value3"
//     renamed_field = "value4"
//   }
func ListToBlocksWithMapping(attrName string, fieldMapping map[string]string) basic.TransformerFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		if attrName == "" {
			return nil
		}
		
		body := block.Body()
		attr := body.GetAttribute(attrName)
		
		if attr == nil {
			return nil // Attribute doesn't exist, nothing to convert
		}
		
		// Get the attribute expression
		expr := attr.Expr()
		tokens := expr.BuildTokens(nil)
		
		// Check if this is actually a list (starts with '[')
		isList := false
		for _, token := range tokens {
			if token.Type == hclsyntax.TokenOBrack {
				isList = true
				break
			}
			// Skip whitespace
			if token.Type != hclsyntax.TokenNewline && 
			   !(token.Type == hclsyntax.TokenIdent && strings.TrimSpace(string(token.Bytes)) == "") {
				// Non-whitespace before bracket means it's not a list
				break
			}
		}
		
		if !isList {
			// Not a list, nothing to convert
			return nil
		}
		
		// Parse the list structure
		items := parseListExpression(tokens)
		
		// Always remove the list attribute (even if empty)
		body.RemoveAttribute(attrName)
		
		// Create a block for each item
		for _, itemTokens := range items {
			newBlock := body.AppendNewBlock(attrName, nil)
			
			// Parse the item as an object
			parsed := parseObjectFromTokens(itemTokens)
			
			if len(parsed) > 0 {
				// Add fields with mapping
				for fieldName, fieldTokens := range parsed {
					// Check if this field should be renamed
					targetName := fieldName
					if mappedName, ok := fieldMapping[fieldName]; ok {
						targetName = mappedName
					}
					newBlock.Body().SetAttributeRaw(targetName, fieldTokens)
				}
			} else {
				// If it's not an object, wrap it in a value attribute
				newBlock.Body().SetAttributeRaw("value", itemTokens)
			}
		}
		
		return nil
	}
}

