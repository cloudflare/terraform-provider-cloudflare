package structural

import (
	"fmt"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// ArrayToObjectArrayConverter transforms array attributes from simple values to objects
// For example: ["id1", "id2"] becomes [{ id = "id1" }, { id = "id2" }]
func ArrayToObjectArrayConverter(attributeName string, wrapperField string) basic.TransformerFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		attr := body.GetAttribute(attributeName)
		
		if attr == nil {
			return nil
		}
		
		// Get the expression as a string
		tokens := attr.Expr().BuildTokens(nil)
		exprStr := string(tokens.Bytes())
		
		// Transform the expression
		transformed := transformArrayToObjectArray(exprStr, wrapperField)
		
		// Set the new expression
		body.SetAttributeRaw(attributeName, hclwrite.Tokens{
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(transformed)},
		})
		
		return nil
	}
}

// StringArrayToObjectArrayWithRename is similar but also handles resource renaming
// For example: cloudflare_access_policy.foo becomes cloudflare_zero_trust_access_policy.foo
func StringArrayToObjectArrayWithRename(attributeName string, wrapperField string, oldPrefix string, newPrefix string) basic.TransformerFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		attr := body.GetAttribute(attributeName)
		
		if attr == nil {
			return nil
		}
		
		// Get the expression as a string
		tokens := attr.Expr().BuildTokens(nil)
		exprStr := string(tokens.Bytes())
		
		// Apply rename if needed
		if oldPrefix != "" && newPrefix != "" {
			exprStr = strings.ReplaceAll(exprStr, oldPrefix, newPrefix)
		}
		
		// Transform the expression
		transformed := transformArrayToObjectArray(exprStr, wrapperField)
		
		// Set the new expression
		body.SetAttributeRaw(attributeName, hclwrite.Tokens{
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(transformed)},
		})
		
		return nil
	}
}

// transformArrayToObjectArray transforms array expression from strings to objects
func transformArrayToObjectArray(expr string, wrapperField string) string {
	// Remove surrounding brackets and whitespace
	expr = strings.TrimSpace(expr)
	if strings.HasPrefix(expr, "[") && strings.HasSuffix(expr, "]") {
		expr = expr[1:len(expr)-1]
	}
	
	// Handle empty arrays
	if strings.TrimSpace(expr) == "" {
		return "[]"
	}
	
	// Split by comma (handling nested structures carefully)
	parts := splitArrayElements(expr)
	var transformed []string
	
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		
		// Check if it's already an object (has braces)
		if strings.Contains(part, "{") && strings.Contains(part, "}") {
			// Already transformed, keep as is
			transformed = append(transformed, part)
			continue
		}
		
		// Transform to object with wrapper field
		transformed = append(transformed, fmt.Sprintf(`{ %s = %s }`, wrapperField, part))
	}
	
	return "[" + strings.Join(transformed, ", ") + "]"
}

// splitArrayElements splits array elements by commas, handling nested structures
func splitArrayElements(expr string) []string {
	var parts []string
	var current strings.Builder
	depth := 0
	inString := false
	escape := false
	
	for i, ch := range expr {
		if escape {
			current.WriteRune(ch)
			escape = false
			continue
		}
		
		switch ch {
		case '\\':
			escape = true
			current.WriteRune(ch)
		case '"':
			if !escape {
				inString = !inString
			}
			current.WriteRune(ch)
		case '[', '{', '(':
			if !inString {
				depth++
			}
			current.WriteRune(ch)
		case ']', '}', ')':
			if !inString {
				depth--
			}
			current.WriteRune(ch)
		case ',':
			if depth == 0 && !inString {
				parts = append(parts, current.String())
				current.Reset()
			} else {
				current.WriteRune(ch)
			}
		default:
			current.WriteRune(ch)
		}
		
		// Handle last character
		if i == len(expr)-1 && current.Len() > 0 {
			parts = append(parts, current.String())
		}
	}
	
	return parts
}