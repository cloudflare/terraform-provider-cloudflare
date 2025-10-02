package structural

import (
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// FlattenNestedTransformer flattens nested object structures
func FlattenNestedTransformer(source string, separator string, maxDepth int) basic.TransformerFunc {
	// Don't default separator - allow empty separator for concatenation
	if maxDepth <= 0 {
		return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
			// maxDepth of 0 means don't flatten
			return nil
		}
	}
	
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		attr := body.GetAttribute(source)
		
		if attr == nil {
			return nil
		}
		
		// Get the expression tokens
		tokens := attr.Expr().BuildTokens(nil)
		
		// Parse and flatten the object
		flattened := parseAndFlattenObject(tokens, source, separator, 0, maxDepth)
		
		if len(flattened) > 0 {
			// Remove the original attribute
			body.RemoveAttribute(source)
			
			// Add flattened attributes
			for key, value := range flattened {
				body.SetAttributeRaw(key, value)
			}
		}
		
		return nil
	}
}

// FlattenWithPrefix flattens and adds a prefix to all resulting attributes
func FlattenWithPrefix(source string, prefix string, separator string) basic.TransformerFunc {
	if separator == "" {
		separator = "_"
	}
	
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		attr := body.GetAttribute(source)
		
		if attr == nil {
			return nil
		}
		
		// Get the expression tokens
		tokens := attr.Expr().BuildTokens(nil)
		
		// Parse and flatten with prefix
		flattened := parseAndFlattenObject(tokens, prefix, separator, 0, 3)
		
		if len(flattened) > 0 {
			// Remove the original attribute
			body.RemoveAttribute(source)
			
			// Add flattened attributes
			for key, value := range flattened {
				body.SetAttributeRaw(key, value)
			}
		}
		
		return nil
	}
}

// NestedRestructureTransformer restructures nested attributes
func NestedRestructureTransformer(source string, paths map[string]string) basic.TransformerFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		attr := body.GetAttribute(source)
		
		if attr == nil {
			return nil
		}
		
		// Get the expression tokens
		tokens := attr.Expr().BuildTokens(nil)
		
		// If no paths specified, flatten with default naming
		if len(paths) == 0 {
			flattened := parseAndFlattenObject(tokens, source, "_", 0, 3)
			
			if len(flattened) > 0 {
				body.RemoveAttribute(source)
				
				for key, value := range flattened {
					body.SetAttributeRaw(key, value)
				}
			}
			return nil
		}
		
		// Parse the object to extract specific paths
		parsed := parseObjectToMap(tokens)
		
		if len(parsed) > 0 {
			body.RemoveAttribute(source)
			
			// Extract and set values for each path
			for path, targetAttr := range paths {
				if value := getValueAtPath(parsed, path); value != nil {
					body.SetAttributeRaw(targetAttr, value)
				}
			}
			
			// Also add any unmapped fields with default names
			flattenUnmappedFields(body, parsed, source, "_", paths)
		}
		
		return nil
	}
}

// parseAndFlattenObject parses an object from tokens and flattens it
func parseAndFlattenObject(tokens hclwrite.Tokens, prefix string, separator string, depth int, maxDepth int) map[string]hclwrite.Tokens {
	result := make(map[string]hclwrite.Tokens)
	
	if depth >= maxDepth {
		if prefix != "" {
			// Remove trailing separator
			key := strings.TrimSuffix(prefix, separator)
			result[key] = tokens
		}
		return result
	}
	
	// Parse the object
	parsed := parseObjectToMap(tokens)
	if len(parsed) == 0 {
		// Not an object, return as-is
		if prefix != "" {
			key := strings.TrimSuffix(prefix, separator)
			result[key] = tokens
		}
		return result
	}
	
	// Process each field
	for fieldName, fieldTokens := range parsed {
		newPrefix := fieldName
		if prefix != "" {
			// Concatenate with separator (which might be empty)
			newPrefix = prefix + separator + fieldName
		}
		
		// Try to flatten further
		nested := parseAndFlattenObject(fieldTokens, newPrefix, separator, depth+1, maxDepth)
		if len(nested) == 0 {
			// Can't flatten further
			result[newPrefix] = fieldTokens
		} else {
			// Add nested results
			for k, v := range nested {
				result[k] = v
			}
		}
	}
	
	return result
}

// parseObjectToMap parses object tokens into a map of field names to token values
func parseObjectToMap(tokens hclwrite.Tokens) map[string]hclwrite.Tokens {
	result := make(map[string]hclwrite.Tokens)
	
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
					result[currentKey] = trimTokens(currentValue)
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
				// This is a field name
				currentKey = string(token.Bytes)
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
					result[currentKey] = trimTokens(currentValue)
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
	
	// Handle case where there's no trailing comma/newline
	if currentKey != "" && len(currentValue) > 0 {
		result[currentKey] = trimTokens(currentValue)
	}
	
	return result
}

// trimTokens removes leading/trailing whitespace tokens
func trimTokens(tokens hclwrite.Tokens) hclwrite.Tokens {
	// Trim leading whitespace/newlines
	start := 0
	for start < len(tokens) && isWhitespaceToken(tokens[start]) {
		start++
	}
	
	// Trim trailing whitespace/newlines
	end := len(tokens)
	for end > start && isWhitespaceToken(tokens[end-1]) {
		end--
	}
	
	if start >= end {
		return hclwrite.Tokens{}
	}
	
	return tokens[start:end]
}

// isWhitespaceToken checks if a token is whitespace or newline
func isWhitespaceToken(token *hclwrite.Token) bool {
	return token.Type == hclsyntax.TokenNewline || 
		(token.Type == hclsyntax.TokenIdent && strings.TrimSpace(string(token.Bytes)) == "")
}

// getValueAtPath retrieves a value from parsed object using dot-notation path
func getValueAtPath(parsed map[string]hclwrite.Tokens, path string) hclwrite.Tokens {
	parts := strings.Split(path, ".")
	if len(parts) == 0 {
		return nil
	}
	
	// Get the first part
	firstPart := parts[0]
	value, ok := parsed[firstPart]
	if !ok {
		return nil
	}
	
	// If this is the last part, return the value
	if len(parts) == 1 {
		return value
	}
	
	// Otherwise, parse this value as an object and recurse
	subParsed := parseObjectToMap(value)
	if len(subParsed) == 0 {
		return nil
	}
	
	remainingPath := strings.Join(parts[1:], ".")
	return getValueAtPath(subParsed, remainingPath)
}

// flattenUnmappedFields adds unmapped fields with default naming
func flattenUnmappedFields(body *hclwrite.Body, parsed map[string]hclwrite.Tokens, prefix string, separator string, mappedPaths map[string]string) {
	// For each field in the parsed object, check if any of its paths were mapped
	for fieldName, fieldTokens := range parsed {
		// Check if this is a nested object
		subParsed := parseObjectToMap(fieldTokens)
		
		if len(subParsed) > 0 {
			// It's a nested object - check each nested field
			for nestedFieldName, nestedTokens := range subParsed {
				fullPath := fieldName + "." + nestedFieldName
				
				// Check if this specific path was mapped
				mapped := false
				for mappedPath := range mappedPaths {
					if mappedPath == fullPath {
						mapped = true
						break
					}
				}
				
				if !mapped {
					// This nested field wasn't mapped, add it with default naming
					attrName := prefix + separator + fieldName + separator + nestedFieldName
					
					// Check if we need to flatten further
					deepParsed := parseObjectToMap(nestedTokens)
					if len(deepParsed) > 0 {
						// Recursively flatten deeper levels
						flattened := parseAndFlattenObject(nestedTokens, attrName, separator, 2, 3)
						for k, v := range flattened {
							body.SetAttributeRaw(k, v)
						}
					} else {
						body.SetAttributeRaw(attrName, nestedTokens)
					}
				}
			}
		} else {
			// It's a simple field - check if the field itself was mapped
			mapped := false
			for mappedPath := range mappedPaths {
				if mappedPath == fieldName {
					mapped = true
					break
				}
			}
			
			if !mapped {
				// This field wasn't mapped, add it with default naming
				attrName := prefix + separator + fieldName
				body.SetAttributeRaw(attrName, fieldTokens)
			}
		}
	}
}