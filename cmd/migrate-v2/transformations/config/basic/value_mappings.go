package basic

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// ValueMapper creates a transformer that maps attribute values
//
// Example YAML configuration:
//   value_mappings:
//     - attribute: type
//       mappings:
//         firewall: ruleset
//         access_app: zero_trust_app
//         old_type: new_type
//     - attribute: status
//       mappings:
//         on: enabled
//         off: disabled
//         1: active
//         0: inactive
//
// Transforms:
//   resource "example" "test" {
//     type = "firewall"
//     status = "on"
//     other_type = "access_app"  # Not mapped (different attribute)
//   }
//
// Into:
//   resource "example" "test" {
//     type = "ruleset"
//     status = "enabled"
//     other_type = "access_app"  # Unchanged
//   }
func ValueMapper(mappings []ValueMapping) TransformFunc {
	if len(mappings) == 0 {
		return func(block *hclwrite.Block, ctx *TransformContext) error {
			return nil
		}
	}

	return func(block *hclwrite.Block, ctx *TransformContext) error {
		body := block.Body()
		
		for _, mapping := range mappings {
			// Get the attribute
			attr := body.GetAttribute(mapping.Attribute)
			if attr == nil {
				continue
			}
			
			// Get the current value as a string
			currentValue, err := getAttributeValueAsString(attr)
			if err != nil {
				// Skip if we can't parse the value
				continue
			}
			
			// Apply value transformation based on type
			var newValue string
			transformed := false
			
			switch mapping.Type {
			case "boolean_to_string":
				// Handle boolean to string conversion
				if currentValue == "true" && mapping.TrueValue != "" {
					newValue = mapping.TrueValue
					transformed = true
				} else if currentValue == "false" && mapping.FalseValue != "" {
					newValue = mapping.FalseValue
					transformed = true
				}
				
			default:
				// Use explicit mappings
				if mapping.Mappings != nil {
					if mappedValue, ok := mapping.Mappings[currentValue]; ok {
						newValue = mappedValue
						transformed = true
					}
				}
			}
			
			// Apply the transformation if we found a mapping
			if transformed {
				body.SetAttributeValue(mapping.Attribute, cty.StringVal(newValue))
			}
			
			// Handle rename if specified
			if mapping.RenameTo != "" && mapping.RenameTo != mapping.Attribute {
				// Get the current value (possibly transformed)
				attr = body.GetAttribute(mapping.Attribute)
				if attr != nil {
					// Copy to new name
					body.SetAttributeRaw(mapping.RenameTo, attr.Expr().BuildTokens(nil))
					// Remove old attribute
					body.RemoveAttribute(mapping.Attribute)
				}
			}
		}
		
		return nil
	}
}

// ValueMapperForState creates a transformer for state value mappings
func ValueMapperForState(mappings []ValueMapping) func(map[string]interface{}) error {
	return func(state map[string]interface{}) error {
		if state == nil {
			return nil
		}
		
		// Get attributes map
		attributes, ok := state["attributes"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("state does not contain attributes map")
		}
		
		for _, mapping := range mappings {
			// Get current value
			currentValue, exists := attributes[mapping.Attribute]
			if !exists {
				continue
			}
			
			// Convert to string for mapping
			currentStr := fmt.Sprintf("%v", currentValue)
			
			// Apply transformation
			var newValue interface{}
			transformed := false
			
			switch mapping.Type {
			case "boolean_to_string":
				if currentStr == "true" && mapping.TrueValue != "" {
					newValue = mapping.TrueValue
					transformed = true
				} else if currentStr == "false" && mapping.FalseValue != "" {
					newValue = mapping.FalseValue
					transformed = true
				}
				
			default:
				if mapping.Mappings != nil {
					if mappedValue, ok := mapping.Mappings[currentStr]; ok {
						newValue = mappedValue
						transformed = true
					}
				}
			}
			
			// Apply transformation
			if transformed {
				attributes[mapping.Attribute] = newValue
			}
			
			// Handle rename
			if mapping.RenameTo != "" && mapping.RenameTo != mapping.Attribute {
				attributes[mapping.RenameTo] = attributes[mapping.Attribute]
				delete(attributes, mapping.Attribute)
			}
		}
		
		return nil
	}
}

// Helper function to get attribute value as string
func getAttributeValueAsString(attr *hclwrite.Attribute) (string, error) {
	if attr == nil {
		return "", fmt.Errorf("attribute is nil")
	}
	
	// Get the tokens
	tokens := attr.Expr().BuildTokens(nil)
	if len(tokens) == 0 {
		return "", fmt.Errorf("no tokens in attribute")
	}
	
	// Convert tokens to string and clean up
	value := string(tokens.Bytes())
	value = strings.TrimSpace(value)
	
	// Remove quotes if present
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
		value = strings.Trim(value, `"`)
	}
	
	return value, nil
}