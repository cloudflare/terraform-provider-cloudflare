package conditional

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
)

// ConditionalTransformer creates a transformer that applies conditional logic
//
// Example YAML configuration:
//   conditional_transforms:
//     - condition:
//         attribute: tier
//         operator: equals
//         value: premium
//       then:
//         set_attributes:
//           advanced_features: true
//           max_retries: 5
//         rename_attributes:
//           basic_setting: premium_setting
//       else:
//         set_attributes:
//           advanced_features: false
//           max_retries: 3
//
// Transforms (when tier = "premium"):
//   resource "example" "test" {
//     tier = "premium"
//     basic_setting = "value"
//   }
//
// Into:
//   resource "example" "test" {
//     tier = "premium"
//     premium_setting = "value"
//     advanced_features = true
//     max_retries = 5
//   }
//
// Supported operators: exists, not_exists, equals, not_equals, contains, 
// starts_with, ends_with, is_empty, is_not_empty
func ConditionalTransformer(transforms []basic.ConditionalTransform) basic.TransformFunc {
	if len(transforms) == 0 {
		return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
			return nil
		}
	}

	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		
		for _, transform := range transforms {
			// Evaluate condition
			conditionMet := evaluateCondition(body, transform.Condition)
			
			// Apply appropriate actions
			var actions *basic.TransformActions
			if conditionMet {
				actions = &transform.Then
			} else if transform.Else != nil {
				actions = transform.Else
			} else {
				continue // No else clause and condition not met
			}
			
			// Apply the actions
			if err := applyActions(body, actions); err != nil {
				return fmt.Errorf("failed to apply conditional actions: %w", err)
			}
		}
		
		return nil
	}
}

// evaluateCondition checks if a condition is met
func evaluateCondition(body *hclwrite.Body, condition basic.TransformCondition) bool {
	attr := body.GetAttribute(condition.Attribute)
	
	switch condition.Operator {
	case "exists":
		return attr != nil
		
	case "not_exists":
		return attr == nil
		
	case "equals", "==":
		if attr == nil {
			return false
		}
		attrValue := getAttributeValue(attr)
		return attrValue == condition.Value
		
	case "not_equals", "!=":
		if attr == nil {
			return condition.Value != ""
		}
		attrValue := getAttributeValue(attr)
		return attrValue != condition.Value
		
	case "contains":
		if attr == nil {
			return false
		}
		attrValue := getAttributeValue(attr)
		return strings.Contains(attrValue, condition.Value)
		
	case "starts_with":
		if attr == nil {
			return false
		}
		attrValue := getAttributeValue(attr)
		return strings.HasPrefix(attrValue, condition.Value)
		
	case "ends_with":
		if attr == nil {
			return false
		}
		attrValue := getAttributeValue(attr)
		return strings.HasSuffix(attrValue, condition.Value)
		
	case "is_empty":
		if attr == nil {
			return true
		}
		attrValue := getAttributeValue(attr)
		return attrValue == "" || attrValue == "[]" || attrValue == "{}"
		
	case "is_not_empty":
		if attr == nil {
			return false
		}
		attrValue := getAttributeValue(attr)
		return attrValue != "" && attrValue != "[]" && attrValue != "{}"
		
	default:
		// Unknown operator, default to false
		return false
	}
}

// getAttributeValue extracts the string value of an attribute
func getAttributeValue(attr *hclwrite.Attribute) string {
	if attr == nil {
		return ""
	}
	
	tokens := attr.Expr().BuildTokens(nil)
	value := strings.TrimSpace(string(tokens.Bytes()))
	
	// Remove quotes if present
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
		value = value[1:len(value)-1]
	}
	
	return value
}

// applyActions applies transform actions to a body
func applyActions(body *hclwrite.Body, actions *basic.TransformActions) error {
	if actions == nil {
		return nil
	}
	
	// Remove attributes
	for _, attrName := range actions.RemoveAttributes {
		body.RemoveAttribute(attrName)
	}
	
	// Set attributes
	for attrName, value := range actions.SetAttributes {
		// Parse the value to determine its type
		if value == "true" || value == "false" {
			if value == "true" {
				body.SetAttributeValue(attrName, cty.BoolVal(true))
			} else {
				body.SetAttributeValue(attrName, cty.BoolVal(false))
			}
		} else if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
			// List - create as raw expression
			body.SetAttributeRaw(attrName, parseRawTokens(value))
		} else if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
			// Object - create as raw expression
			body.SetAttributeRaw(attrName, parseRawTokens(value))
		} else {
			// String or number
			body.SetAttributeValue(attrName, cty.StringVal(value))
		}
	}
	
	// Rename attributes
	for oldName, newName := range actions.RenameAttributes {
		if attr := body.GetAttribute(oldName); attr != nil {
			// Copy to new name
			body.SetAttributeRaw(newName, attr.Expr().BuildTokens(nil))
			// Remove old
			body.RemoveAttribute(oldName)
		}
	}
	
	return nil
}

// Helper functions for parsing values
func parseBoolValue(value string) interface{} {
	return value == "true"
}

func parseStringValue(value string) interface{} {
	// Try to parse as number
	if _, err := fmt.Sscanf(value, "%d", new(int)); err == nil {
		var n int
		fmt.Sscanf(value, "%d", &n)
		return n
	}
	if _, err := fmt.Sscanf(value, "%f", new(float64)); err == nil {
		var f float64
		fmt.Sscanf(value, "%f", &f)
		return f
	}
	return value
}

func parseRawTokens(value string) hclwrite.Tokens {
	// Create a temporary HCL file to parse the expression
	tempConfig := fmt.Sprintf("x = %s", value)
	tempFile, _ := hclwrite.ParseConfig([]byte(tempConfig), "", hcl.InitialPos)
	if tempFile != nil && tempFile.Body() != nil {
		if tempAttr := tempFile.Body().GetAttribute("x"); tempAttr != nil {
			return tempAttr.Expr().BuildTokens(nil)
		}
	}
	// Fallback to empty tokens
	return hclwrite.Tokens{}
}

// ConditionalTransformerForState applies conditional transformations to state
func ConditionalTransformerForState(transforms []basic.ConditionalTransform) func(map[string]interface{}) error {
	return func(state map[string]interface{}) error {
		if state == nil {
			return nil
		}
		
		attributes, ok := state["attributes"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("state does not contain attributes map")
		}
		
		for _, transform := range transforms {
			// Evaluate condition on state
			conditionMet := evaluateStateCondition(attributes, transform.Condition)
			
			// Apply appropriate actions
			var actions *basic.TransformActions
			if conditionMet {
				actions = &transform.Then
			} else if transform.Else != nil {
				actions = transform.Else
			} else {
				continue
			}
			
			// Apply actions to state
			applyStateActions(attributes, actions)
		}
		
		return nil
	}
}

// evaluateStateCondition evaluates a condition on state attributes
func evaluateStateCondition(attributes map[string]interface{}, condition basic.TransformCondition) bool {
	value, exists := attributes[condition.Attribute]
	
	switch condition.Operator {
	case "exists":
		return exists
		
	case "not_exists":
		return !exists
		
	case "equals", "==":
		if !exists {
			return false
		}
		return fmt.Sprintf("%v", value) == condition.Value
		
	case "not_equals", "!=":
		if !exists {
			return condition.Value != ""
		}
		return fmt.Sprintf("%v", value) != condition.Value
		
	case "contains":
		if !exists {
			return false
		}
		strValue := fmt.Sprintf("%v", value)
		return strings.Contains(strValue, condition.Value)
		
	case "is_empty":
		if !exists {
			return true
		}
		switch v := value.(type) {
		case string:
			return v == ""
		case []interface{}:
			return len(v) == 0
		case map[string]interface{}:
			return len(v) == 0
		default:
			return false
		}
		
	case "is_not_empty":
		if !exists {
			return false
		}
		switch v := value.(type) {
		case string:
			return v != ""
		case []interface{}:
			return len(v) > 0
		case map[string]interface{}:
			return len(v) > 0
		default:
			return true
		}
		
	default:
		return false
	}
}

// applyStateActions applies actions to state attributes
func applyStateActions(attributes map[string]interface{}, actions *basic.TransformActions) {
	if actions == nil {
		return
	}
	
	// Remove attributes
	for _, attrName := range actions.RemoveAttributes {
		delete(attributes, attrName)
	}
	
	// Set attributes
	for attrName, value := range actions.SetAttributes {
		attributes[attrName] = parseStateValue(value)
	}
	
	// Rename attributes
	for oldName, newName := range actions.RenameAttributes {
		if val, exists := attributes[oldName]; exists {
			attributes[newName] = val
			delete(attributes, oldName)
		}
	}
}

// parseStateValue parses a string value for state
func parseStateValue(value string) interface{} {
	// Boolean
	if value == "true" {
		return true
	}
	if value == "false" {
		return false
	}
	
	// Number
	if _, err := fmt.Sscanf(value, "%d", new(int)); err == nil {
		var num int
		fmt.Sscanf(value, "%d", &num)
		return num
	}
	
	// Default to string
	return value
}