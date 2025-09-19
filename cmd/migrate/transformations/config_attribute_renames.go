package transformations

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// ApplyAttributeRenames applies simple attribute renames to a resource block
func ApplyAttributeRenames(config *TransformationConfig, block *hclwrite.Block, resourceType string) error {
	renames, exists := config.AttributeRenames[resourceType]
	if !exists && resourceType != "cloudflare_page_rule" {
		return nil // No renames for this resource type
	}

	body := block.Body()

	// Apply renames if they exist
	if exists {
		for oldName, newName := range renames {
			attr := body.GetAttribute(oldName)
			if attr != nil {
				// Get the expression tokens
				expr := attr.Expr()
				tokens := expr.BuildTokens(nil)

				// Remove old attribute
				body.RemoveAttribute(oldName)

				// Set new attribute with same value
				body.SetAttributeRaw(newName, tokens)
			}
		}
	}

	// Special handling for cloudflare_page_rule to add status attribute if not present
	if resourceType == "cloudflare_page_rule" {
		applyPageRuleStatusDefault(body)
	}

	return nil
}

// applyPageRuleStatusDefault adds status = "active" if the attribute is not set
func applyPageRuleStatusDefault(body *hclwrite.Body) {
	// Check if status attribute exists
	if body.GetAttribute("status") == nil {
		// Add status = "active" as default for v5
		body.SetAttributeValue("status", cty.StringVal("active"))
	}
}

// ApplyAttributeRemovals removes specified attributes from a resource block
func ApplyAttributeRemovals(config *TransformationConfig, block *hclwrite.Block, resourceType string) error {
	removals, exists := config.AttributeRemovals[resourceType]
	if !exists {
		return nil // No removals for this resource type
	}

	body := block.Body()

	for _, attrName := range removals {
		body.RemoveAttribute(attrName)
	}

	return nil
}

// HasAttributeRename checks if an attribute should be renamed
func HasAttributeRename(config *TransformationConfig, resourceType, attrName string) (string, bool) {
	renames, exists := config.AttributeRenames[resourceType]
	if !exists {
		return "", false
	}

	newName, hasRename := renames[attrName]
	return newName, hasRename
}

// ShouldRemoveAttribute checks if an attribute should be removed
func ShouldRemoveAttribute(config *TransformationConfig, resourceType, attrName string) bool {
	removals, exists := config.AttributeRemovals[resourceType]
	if !exists {
		return false
	}

	for _, removal := range removals {
		if removal == attrName {
			return true
		}
	}

	return false
}