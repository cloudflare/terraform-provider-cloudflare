package structural

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
)

// ListTransformer creates a transformer for list operations
//
// Example YAML configuration:
//   list_transforms:
//     - attribute: ip_addresses
//       type: wrap_single
//     - attribute: allowed_methods
//       type: flatten
//     - attribute: ports
//       type: first_element
//     - attribute: servers
//       type: join
//       separator: ","
//
// Transforms:
//   resource "example" "test" {
//     ip_addresses = "192.168.1.1"
//     allowed_methods = [["GET", "POST"], ["PUT", "DELETE"]]
//     ports = [8080, 8081, 8082]
//     servers = ["server1", "server2", "server3"]
//   }
//
// Into:
//   resource "example" "test" {
//     ip_addresses = ["192.168.1.1"]
//     allowed_methods = ["GET", "POST", "PUT", "DELETE"]
//     ports = 8080
//     servers = "server1,server2,server3"
//   }
func ListTransformer(transforms []basic.ListTransform) basic.TransformFunc {
	if len(transforms) == 0 {
		return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
			return nil
		}
	}

	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		
		for _, transform := range transforms {
			attr := body.GetAttribute(transform.Attribute)
			if attr == nil {
				continue
			}
			
			switch transform.Type {
			case "string_to_object":
				err := transformStringListToObjectList(body, transform)
				if err != nil {
					return fmt.Errorf("failed to transform string list to object list for %s: %w", transform.Attribute, err)
				}
				
			case "wrap_strings":
				err := wrapStringList(body, transform)
				if err != nil {
					return fmt.Errorf("failed to wrap string list for %s: %w", transform.Attribute, err)
				}
				
			case "id_list_to_object_list":
				err := transformIDListToObjectList(body, transform)
				if err != nil {
					return fmt.Errorf("failed to transform ID list for %s: %w", transform.Attribute, err)
				}
				
			default:
				// Unknown transformation type - skip
				continue
			}
		}
		
		return nil
	}
}

// transformStringListToObjectList converts ["item1", "item2"] to [{ value = "item1" }, { value = "item2" }]
func transformStringListToObjectList(body *hclwrite.Body, transform basic.ListTransform) error {
	attr := body.GetAttribute(transform.Attribute)
	if attr == nil {
		return nil
	}
	
	// Parse the existing list
	tokens := attr.Expr().BuildTokens(nil)
	tokenStr := strings.TrimSpace(string(tokens.Bytes()))
	
	// Simple parsing for string lists
	if !strings.HasPrefix(tokenStr, "[") || !strings.HasSuffix(tokenStr, "]") {
		return nil // Not a list literal
	}
	
	// Extract items from the list
	listContent := strings.TrimSpace(tokenStr[1 : len(tokenStr)-1])
	if listContent == "" {
		// Empty list - keep as is
		return nil
	}
	
	// Parse items (simple approach for string literals)
	items := parseListItems(listContent)
	
	if len(items) == 0 {
		return nil
	}
	
	// Build new HCL expression string
	var result strings.Builder
	result.WriteString("[\n")
	
	for i, item := range items {
		if transform.ObjectTemplate != nil {
			result.WriteString("    { ")
			first := true
			for key, templateValue := range transform.ObjectTemplate {
				if !first {
					result.WriteString(", ")
				}
				// Replace ${item} with actual item value
				value := strings.ReplaceAll(templateValue, "${item}", item)
				result.WriteString(fmt.Sprintf("%s = \"%s\"", key, value))
				first = false
			}
			result.WriteString(" }")
			if i < len(items)-1 {
				result.WriteString(",")
			}
			result.WriteString("\n")
		}
	}
	
	result.WriteString("  ]")
	
	// Parse the new expression and set it
	newExpr := result.String()
	// We need to create a temporary file to parse the expression
	tempConfig := fmt.Sprintf("x = %s", newExpr)
	tempFile, _ := hclwrite.ParseConfig([]byte(tempConfig), "", hcl.InitialPos)
	if tempFile != nil && tempFile.Body() != nil {
		if tempAttr := tempFile.Body().GetAttribute("x"); tempAttr != nil {
			body.SetAttributeRaw(transform.Attribute, tempAttr.Expr().BuildTokens(nil))
		}
	}
	
	return nil
}

// wrapStringList wraps each string in a list with a key
func wrapStringList(body *hclwrite.Body, transform basic.ListTransform) error {
	attr := body.GetAttribute(transform.Attribute)
	if attr == nil {
		return nil
	}
	
	// Parse the existing list
	tokens := attr.Expr().BuildTokens(nil)
	tokenStr := strings.TrimSpace(string(tokens.Bytes()))
	
	if !strings.HasPrefix(tokenStr, "[") || !strings.HasSuffix(tokenStr, "]") {
		return nil // Not a list literal
	}
	
	// Extract items
	listContent := strings.TrimSpace(tokenStr[1 : len(tokenStr)-1])
	if listContent == "" {
		return nil
	}
	
	items := parseListItems(listContent)
	
	if len(items) == 0 || transform.WrapperKey == "" {
		return nil
	}
	
	// Build new HCL expression string
	var result strings.Builder
	result.WriteString("[\n")
	
	for i, item := range items {
		result.WriteString(fmt.Sprintf("    { %s = \"%s\" }", transform.WrapperKey, item))
		if i < len(items)-1 {
			result.WriteString(",")
		}
		result.WriteString("\n")
	}
	
	result.WriteString("  ]")
	
	// Parse the new expression and set it
	newExpr := result.String()
	// We need to create a temporary file to parse the expression
	tempConfig := fmt.Sprintf("x = %s", newExpr)
	tempFile, _ := hclwrite.ParseConfig([]byte(tempConfig), "", hcl.InitialPos)
	if tempFile != nil && tempFile.Body() != nil {
		if tempAttr := tempFile.Body().GetAttribute("x"); tempAttr != nil {
			body.SetAttributeRaw(transform.Attribute, tempAttr.Expr().BuildTokens(nil))
		}
	}
	
	return nil
}

// transformIDListToObjectList converts ["id1", "id2"] to [{ pool = "id1" }, { pool = "id2" }]
func transformIDListToObjectList(body *hclwrite.Body, transform basic.ListTransform) error {
	attr := body.GetAttribute(transform.Attribute)
	if attr == nil {
		return nil
	}
	
	// Parse the existing list
	tokens := attr.Expr().BuildTokens(nil)
	tokenStr := strings.TrimSpace(string(tokens.Bytes()))
	
	if !strings.HasPrefix(tokenStr, "[") || !strings.HasSuffix(tokenStr, "]") {
		return nil // Not a list literal
	}
	
	// Extract items
	listContent := strings.TrimSpace(tokenStr[1 : len(tokenStr)-1])
	if listContent == "" {
		return nil
	}
	
	items := parseListItems(listContent)
	
	if len(items) == 0 || transform.ObjectKey == "" {
		return nil
	}
	
	// Build new HCL expression string
	var result strings.Builder
	result.WriteString("[\n")
	
	for i, item := range items {
		result.WriteString(fmt.Sprintf("    { %s = \"%s\" }", transform.ObjectKey, item))
		if i < len(items)-1 {
			result.WriteString(",")
		}
		result.WriteString("\n")
	}
	
	result.WriteString("  ]")
	
	// Parse the new expression and set it
	newExpr := result.String()
	// We need to create a temporary file to parse the expression
	tempConfig := fmt.Sprintf("x = %s", newExpr)
	tempFile, _ := hclwrite.ParseConfig([]byte(tempConfig), "", hcl.InitialPos)
	if tempFile != nil && tempFile.Body() != nil {
		if tempAttr := tempFile.Body().GetAttribute("x"); tempAttr != nil {
			body.SetAttributeRaw(transform.Attribute, tempAttr.Expr().BuildTokens(nil))
		}
	}
	
	return nil
}

// parseListItems parses comma-separated items from a list string
func parseListItems(listContent string) []string {
	// Initialize to empty slice instead of nil
	items := []string{}
	var current strings.Builder
	inQuotes := false
	quoteChar := '"'
	
	for i, ch := range listContent {
		switch ch {
		case '"', '\'':
			if !inQuotes {
				inQuotes = true
				quoteChar = ch
			} else if ch == quoteChar {
				inQuotes = false
			}
		case ',':
			if !inQuotes {
				item := strings.TrimSpace(current.String())
				if item != "" {
					// Remove surrounding quotes if present
					item = strings.Trim(item, `"'`)
					items = append(items, item)
				}
				current.Reset()
				continue
			}
		}
		
		// Don't include the quotes in the value
		if ch != '"' && ch != '\'' {
			current.WriteRune(ch)
		}
		
		// Handle last item
		if i == len(listContent)-1 {
			item := strings.TrimSpace(current.String())
			if item != "" {
				items = append(items, item)
			}
		}
	}
	
	return items
}

// ListTransformerForState transforms lists in state
func ListTransformerForState(transforms []basic.ListTransform) func(map[string]interface{}) error {
	return func(state map[string]interface{}) error {
		if state == nil {
			return nil
		}
		
		attributes, ok := state["attributes"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("state does not contain attributes map")
		}
		
		for _, transform := range transforms {
			value, exists := attributes[transform.Attribute]
			if !exists {
				continue
			}
			
			// Handle different list types
			switch v := value.(type) {
			case []interface{}:
				transformed := transformStateList(v, transform)
				if transformed != nil {
					attributes[transform.Attribute] = transformed
				}
				
			case []string:
				// Convert to []interface{} first
				interfaceList := make([]interface{}, len(v))
				for i, s := range v {
					interfaceList[i] = s
				}
				transformed := transformStateList(interfaceList, transform)
				if transformed != nil {
					attributes[transform.Attribute] = transformed
				}
			}
		}
		
		return nil
	}
}

// transformStateList transforms a list in state based on the transform type
func transformStateList(list []interface{}, transform basic.ListTransform) []interface{} {
	var result []interface{}
	
	switch transform.Type {
	case "string_to_object":
		for _, item := range list {
			itemStr := fmt.Sprintf("%v", item)
			
			if transform.ObjectTemplate != nil {
				objMap := make(map[string]interface{})
				for key, templateValue := range transform.ObjectTemplate {
					// Replace ${item} with actual item value
					value := strings.ReplaceAll(templateValue, "${item}", itemStr)
					objMap[key] = value
				}
				result = append(result, objMap)
			}
		}
		
	case "wrap_strings":
		for _, item := range list {
			itemStr := fmt.Sprintf("%v", item)
			
			if transform.WrapperKey != "" {
				objMap := map[string]interface{}{
					transform.WrapperKey: itemStr,
				}
				result = append(result, objMap)
			}
		}
		
	case "id_list_to_object_list":
		for _, item := range list {
			itemStr := fmt.Sprintf("%v", item)
			
			if transform.ObjectKey != "" {
				objMap := map[string]interface{}{
					transform.ObjectKey: itemStr,
				}
				result = append(result, objMap)
			}
		}
		
	default:
		return list // Return unchanged for unknown types
	}
	
	if len(result) > 0 {
		return result
	}
	
	return list
}