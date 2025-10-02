package config

import (
	"strings"
	
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/common"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// ConditionalAttributeRemover removes an attribute based on another attribute's value
// For example: Remove "skip_app_launcher_login_page" if "type" != "app_launcher"
func ConditionalAttributeRemover(targetAttr string, conditionAttr string, allowedValues []string, addDiagnostic bool) common.TransformerFunc {
	return func(block *hclwrite.Block, ctx *common.TransformContext) error {
		body := block.Body()
		
		// Check if target attribute exists
		target := body.GetAttribute(targetAttr)
		if target == nil {
			return nil // Nothing to remove
		}
		
		// Check condition attribute
		condition := body.GetAttribute(conditionAttr)
		
		// Determine if we should keep or remove
		shouldRemove := true
		
		if condition != nil {
			// Get the condition attribute value
			tokens := condition.Expr().BuildTokens(nil)
			condValue := strings.Trim(strings.TrimSpace(string(tokens.Bytes())), `"`)
			
			// Check if value is in allowed list
			for _, allowed := range allowedValues {
				if condValue == allowed {
					shouldRemove = false
					break
				}
			}
		}
		// If condition attribute doesn't exist, we remove by default
		
		if shouldRemove {
			body.RemoveAttribute(targetAttr)
			
			if addDiagnostic && ctx != nil {
				reason := "Removed " + targetAttr + ": Attribute is only valid for specific values"
				if len(allowedValues) > 0 {
					reason = "Removed " + targetAttr + ": Attribute is only valid when " + conditionAttr + " is one of: " + strings.Join(allowedValues, ", ")
				}
				
				ctx.Diagnostics = append(ctx.Diagnostics, reason)
			}
		}
		
		return nil
	}
}

// ConditionalAttributeRemoverWithDefault is similar but handles default values
// If the condition attribute doesn't exist, it uses a default value for comparison
func ConditionalAttributeRemoverWithDefault(targetAttr string, conditionAttr string, allowedValues []string, defaultValue string, addDiagnostic bool) common.TransformerFunc {
	return func(block *hclwrite.Block, ctx *common.TransformContext) error {
		body := block.Body()
		
		// Check if target attribute exists
		target := body.GetAttribute(targetAttr)
		if target == nil {
			return nil // Nothing to remove
		}
		
		// Get condition attribute or use default
		condValue := defaultValue
		condition := body.GetAttribute(conditionAttr)
		
		if condition != nil {
			// Get the actual value
			tokens := condition.Expr().BuildTokens(nil)
			condValue = strings.Trim(strings.TrimSpace(string(tokens.Bytes())), `"`)
		}
		
		// Check if value is in allowed list
		shouldKeep := false
		for _, allowed := range allowedValues {
			if condValue == allowed {
				shouldKeep = true
				break
			}
		}
		
		if !shouldKeep {
			body.RemoveAttribute(targetAttr)
			
			if addDiagnostic && ctx != nil {
				reason := "Removed " + targetAttr + ": Attribute is only valid when " + conditionAttr + " is one of: " + strings.Join(allowedValues, ", ")
				if condition == nil {
					reason += " (defaulted to " + defaultValue + ")"
				}
				
				ctx.Diagnostics = append(ctx.Diagnostics, reason)
			}
		}
		
		return nil
	}
}