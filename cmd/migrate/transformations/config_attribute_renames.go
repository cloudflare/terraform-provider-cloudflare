package transformations

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// ApplyAttributeRenames applies simple attribute renames to a resource block
func ApplyAttributeRenames(config *TransformationConfig, block *hclwrite.Block, resourceType string) error {
	renames, exists := config.AttributeRenames[resourceType]
	if !exists {
		return nil // No renames for this resource type
	}

	body := block.Body()

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

	return nil
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