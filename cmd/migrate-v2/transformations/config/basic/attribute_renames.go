package basic

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// AttributeRenamer creates a transformer that renames attributes
//
// Example YAML configuration:
//   attribute_mappings:
//     old_field: new_field
//     deprecated_name: current_name
//     legacy_setting: modern_setting
//
// Transforms:
//   resource "example" "test" {
//     old_field = "value1"
//     deprecated_name = "value2"
//     legacy_setting = true
//   }
//
// Into:
//   resource "example" "test" {
//     new_field = "value1"
//     current_name = "value2"
//     modern_setting = true
//   }
func AttributeRenamer(mappings map[string]string) TransformerFunc {
	return func(block *hclwrite.Block, ctx *TransformContext) error {
		body := block.Body()
		for oldName, newName := range mappings {
			if attr := body.GetAttribute(oldName); attr != nil {
				expr := attr.Expr()
				body.RemoveAttribute(oldName)
				body.SetAttributeRaw(newName, expr.BuildTokens(nil))
			}
		}
		return nil
	}
}