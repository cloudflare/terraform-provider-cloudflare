package basic

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// ExtendedTypeConverter creates a transformer for extended type conversions
//
// Example YAML configuration:
//   type_conversions:
//     - attribute: port
//       from_type: string
//       to_type: number
//     - attribute: enabled
//       from_type: string
//       to_type: bool
//     - attribute: servers
//       from_type: list
//       to_type: single
//     - attribute: ip_address
//       from_type: string
//       to_type: list
//
// Transforms:
//   resource "example" "test" {
//     port = "8080"
//     enabled = "true"
//     servers = ["server1", "server2"]
//     ip_address = "192.168.1.1"
//   }
//
// Into:
//   resource "example" "test" {
//     port = 8080
//     enabled = true
//     servers = "server1"  # Takes first element
//     ip_address = ["192.168.1.1"]
//   }
func ExtendedTypeConverter(conversions []TypeConversion) TransformFunc {
	if len(conversions) == 0 {
		return func(block *hclwrite.Block, ctx *TransformContext) error {
			return nil
		}
	}

	return func(block *hclwrite.Block, ctx *TransformContext) error {
		body := block.Body()
		
		for _, conversion := range conversions {
			attr := body.GetAttribute(conversion.Attribute)
			if attr == nil {
				continue
			}
			
			// Get the current value as string
			tokens := attr.Expr().BuildTokens(nil)
			currentValue := strings.TrimSpace(string(tokens.Bytes()))
			
			// Perform conversion based on type
			var newValue interface{}
			var err error
			
			switch conversion.ToType {
			case "string":
				newValue, err = convertToString(currentValue, conversion.FromType)
			case "number", "int", "float":
				newValue, err = convertToNumber(currentValue, conversion.FromType)
			case "bool", "boolean":
				newValue, err = convertToBool(currentValue, conversion.FromType)
			case "list":
				newValue, err = convertToList(currentValue, conversion.FromType)
			case "single":
				newValue, err = convertToSingle(currentValue, conversion.FromType)
			default:
				continue // Unknown conversion type
			}
			
			if err != nil {
				// Log error but continue with other conversions
				if ctx.Diagnostics == nil {
					ctx.Diagnostics = []string{}
				}
				ctx.Diagnostics = append(ctx.Diagnostics, 
					fmt.Sprintf("Failed to convert %s from %s to %s: %v", 
						conversion.Attribute, conversion.FromType, conversion.ToType, err))
				continue
			}
			
			// Set the new value
			setConvertedValue(body, conversion.Attribute, newValue, conversion.ToType)
		}
		
		return nil
	}
}

// TypeConversion defines a type conversion
type TypeConversion struct {
	Attribute string `yaml:"attribute"`
	FromType  string `yaml:"from_type"`
	ToType    string `yaml:"to_type"`
}

// convertToString converts various types to string
func convertToString(value string, fromType string) (interface{}, error) {
	// Remove quotes if present
	value = strings.Trim(value, `"'`)
	
	switch fromType {
	case "number", "int", "float":
		// Number to string - just wrap in quotes
		return value, nil
		
	case "bool", "boolean":
		// Boolean to string
		if value == "true" || value == "false" {
			return value, nil
		}
		return "", fmt.Errorf("invalid boolean value: %s", value)
		
	case "list":
		// List to string - take first element or join
		if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
			listContent := strings.TrimSpace(value[1:len(value)-1])
			if listContent == "" {
				return "", nil
			}
			// Parse first element
			elements := parseListElements(listContent)
			if len(elements) > 0 {
				return strings.Trim(elements[0], `"'`), nil
			}
		}
		return value, nil
		
	default:
		return value, nil
	}
}

// convertToNumber converts various types to number
func convertToNumber(value string, fromType string) (interface{}, error) {
	// Remove quotes if present
	value = strings.Trim(value, `"'`)
	
	switch fromType {
	case "string":
		// Try to parse as int first
		if i, err := strconv.ParseInt(value, 10, 64); err == nil {
			return int(i), nil
		}
		// Try float
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			return f, nil
		}
		return 0, fmt.Errorf("cannot convert string '%s' to number", value)
		
	case "bool", "boolean":
		// Boolean to number (0 or 1)
		if value == "true" {
			return 1, nil
		} else if value == "false" {
			return 0, nil
		}
		return 0, fmt.Errorf("invalid boolean value: %s", value)
		
	default:
		// Try direct conversion
		if i, err := strconv.ParseInt(value, 10, 64); err == nil {
			return int(i), nil
		}
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			return f, nil
		}
		return 0, fmt.Errorf("cannot convert '%s' to number", value)
	}
}

// convertToBool converts various types to boolean
func convertToBool(value string, fromType string) (interface{}, error) {
	// Remove quotes if present
	value = strings.Trim(value, `"'`)
	
	switch fromType {
	case "string":
		// String to boolean
		switch strings.ToLower(value) {
		case "true", "yes", "on", "enabled", "1":
			return true, nil
		case "false", "no", "off", "disabled", "0", "":
			return false, nil
		default:
			return false, fmt.Errorf("cannot convert string '%s' to boolean", value)
		}
		
	case "number", "int", "float":
		// Number to boolean
		if value == "0" || value == "0.0" {
			return false, nil
		}
		return true, nil
		
	default:
		// Try standard boolean parsing
		if value == "true" {
			return true, nil
		} else if value == "false" {
			return false, nil
		}
		return false, fmt.Errorf("cannot convert '%s' to boolean", value)
	}
}

// convertToList converts single value to list
func convertToList(value string, fromType string) (interface{}, error) {
	// If already a list, return as is
	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		return value, nil
	}
	
	// Wrap single value in list
	if strings.Contains(value, `"`) {
		// Already quoted
		return fmt.Sprintf("[%s]", value), nil
	} else if fromType == "string" {
		// Add quotes for string
		return fmt.Sprintf(`["%s"]`, value), nil
	} else {
		// Number or boolean - no quotes
		return fmt.Sprintf("[%s]", value), nil
	}
}

// convertToSingle extracts single value from list
func convertToSingle(value string, fromType string) (interface{}, error) {
	if fromType != "list" {
		return value, nil
	}
	
	// Extract from list
	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		listContent := strings.TrimSpace(value[1:len(value)-1])
		if listContent == "" {
			return "", fmt.Errorf("cannot convert empty list to single value")
		}
		
		elements := parseListElements(listContent)
		if len(elements) > 0 {
			// Get first element and determine its type
			elem := strings.Trim(elements[0], `"'`)
			
			// Try to parse as number first
			if i, err := strconv.ParseInt(elem, 10, 64); err == nil {
				return int(i), nil
			}
			if f, err := strconv.ParseFloat(elem, 64); err == nil {
				return f, nil
			}
			// Try boolean
			if elem == "true" {
				return true, nil
			}
			if elem == "false" {
				return false, nil
			}
			// Default to string
			return elem, nil
		}
	}
	
	return value, nil
}

// parseListElements parses comma-separated list elements
func parseListElements(content string) []string {
	if content == "" {
		return []string{}
	}
	
	var elements []string
	var current strings.Builder
	inQuotes := false
	quoteChar := '"'
	
	for _, ch := range content {
		switch ch {
		case '"', '\'':
			if !inQuotes {
				inQuotes = true
				quoteChar = ch
			} else if ch == quoteChar {
				inQuotes = false
			}
			current.WriteRune(ch)
		case ',':
			if !inQuotes {
				if elem := strings.TrimSpace(current.String()); elem != "" {
					elements = append(elements, elem)
				}
				current.Reset()
			} else {
				current.WriteRune(ch)
			}
		default:
			current.WriteRune(ch)
		}
	}
	
	// Add last element
	if elem := strings.TrimSpace(current.String()); elem != "" {
		elements = append(elements, elem)
	}
	
	if elements == nil {
		return []string{}
	}
	
	return elements
}

// setConvertedValue sets the converted value with proper type
func setConvertedValue(body *hclwrite.Body, attribute string, value interface{}, toType string) {
	switch toType {
	case "string":
		// Set as string literal
		if str, ok := value.(string); ok {
			body.SetAttributeValue(attribute, cty.StringVal(str))
		}
		
	case "number", "int", "float":
		// Set as number
		switch v := value.(type) {
		case int:
			body.SetAttributeValue(attribute, cty.NumberIntVal(int64(v)))
		case int64:
			body.SetAttributeValue(attribute, cty.NumberIntVal(v))
		case float64:
			body.SetAttributeValue(attribute, cty.NumberFloatVal(v))
		}
		
	case "bool", "boolean":
		// Set as boolean
		if b, ok := value.(bool); ok {
			body.SetAttributeValue(attribute, cty.BoolVal(b))
		}
		
	case "list":
		// Set as raw expression
		if str, ok := value.(string); ok {
			tempConfig := fmt.Sprintf("x = %s", str)
			tempFile, _ := hclwrite.ParseConfig([]byte(tempConfig), "", hcl.InitialPos)
			if tempFile != nil && tempFile.Body() != nil {
				if tempAttr := tempFile.Body().GetAttribute("x"); tempAttr != nil {
					body.SetAttributeRaw(attribute, tempAttr.Expr().BuildTokens(nil))
				}
			}
		}
		
	case "single":
		// Set based on actual type
		switch v := value.(type) {
		case string:
			body.SetAttributeValue(attribute, cty.StringVal(v))
		case int:
			body.SetAttributeValue(attribute, cty.NumberIntVal(int64(v)))
		case float64:
			body.SetAttributeValue(attribute, cty.NumberFloatVal(v))
		case bool:
			body.SetAttributeValue(attribute, cty.BoolVal(v))
		}
	}
}

// ExtendedTypeConverterForState applies extended type conversions to state
func ExtendedTypeConverterForState(conversions []TypeConversion) func(map[string]interface{}) error {
	return func(state map[string]interface{}) error {
		if state == nil {
			return nil
		}
		
		attributes, ok := state["attributes"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("state does not contain attributes map")
		}
		
		for _, conversion := range conversions {
			value, exists := attributes[conversion.Attribute]
			if !exists {
				continue
			}
			
			var newValue interface{}
			var err error
			
			switch conversion.ToType {
			case "string":
				newValue, err = convertStateToString(value, conversion.FromType)
			case "number", "int", "float":
				newValue, err = convertStateToNumber(value, conversion.FromType)
			case "bool", "boolean":
				newValue, err = convertStateToBool(value, conversion.FromType)
			case "list":
				newValue, err = convertStateToList(value, conversion.FromType)
			case "single":
				newValue, err = convertStateToSingle(value, conversion.FromType)
			}
			
			if err == nil && newValue != nil {
				attributes[conversion.Attribute] = newValue
			}
		}
		
		return nil
	}
}

// State conversion helpers

func convertStateToString(value interface{}, fromType string) (interface{}, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case int, int64, float64:
		return fmt.Sprintf("%v", v), nil
	case bool:
		if v {
			return "true", nil
		}
		return "false", nil
	case []interface{}:
		// Take first element
		if len(v) > 0 {
			return fmt.Sprintf("%v", v[0]), nil
		}
		return "", nil
	default:
		return fmt.Sprintf("%v", value), nil
	}
}

func convertStateToNumber(value interface{}, fromType string) (interface{}, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case float64:
		return v, nil
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return int(i), nil
		}
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f, nil
		}
		return 0, fmt.Errorf("cannot convert string to number: %s", v)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to number", value)
	}
}

func convertStateToBool(value interface{}, fromType string) (interface{}, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case string:
		switch strings.ToLower(v) {
		case "true", "yes", "on", "enabled", "1":
			return true, nil
		case "false", "no", "off", "disabled", "0", "":
			return false, nil
		default:
			return false, fmt.Errorf("cannot convert string to boolean: %s", v)
		}
	case int, int64:
		return v != 0, nil
	case float64:
		return v != 0.0, nil
	default:
		return false, fmt.Errorf("cannot convert %T to boolean", value)
	}
}

func convertStateToList(value interface{}, fromType string) (interface{}, error) {
	switch v := value.(type) {
	case []interface{}:
		return v, nil
	default:
		// Wrap in list
		return []interface{}{value}, nil
	}
}

func convertStateToSingle(value interface{}, fromType string) (interface{}, error) {
	switch v := value.(type) {
	case []interface{}:
		if len(v) > 0 {
			return v[0], nil
		}
		return nil, fmt.Errorf("cannot convert empty list to single value")
	default:
		return value, nil
	}
}