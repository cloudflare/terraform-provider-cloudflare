package basic

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// AttributeRemover creates a transformer that removes attributes
//
// Example YAML configuration:
//   field_removals:
//     - deprecated_field
//     - obsolete_setting
//     - legacy_flag
//
// Transforms:
//   resource "example" "test" {
//     name = "test"
//     deprecated_field = "old_value"
//     obsolete_setting = true
//     legacy_flag = 123
//     current_field = "keep"
//   }
//
// Into:
//   resource "example" "test" {
//     name = "test"
//     current_field = "keep"
//   }
func AttributeRemover(attributes ...string) TransformerFunc {
	return func(block *hclwrite.Block, ctx *TransformContext) error {
		body := block.Body()
		for _, attr := range attributes {
			body.RemoveAttribute(attr)
		}
		return nil
	}
}

// ConditionalRemover removes attributes based on conditions
//
// Example usage:
//   ConditionalRemover("premium_feature", func(block *hclwrite.Block) bool {
//     attr := block.Body().GetAttribute("tier")
//     return attr != nil && getAttributeValue(attr) == "basic"
//   })
//
// Transforms (when tier is "basic"):
//   resource "example" "test" {
//     tier = "basic"
//     premium_feature = "advanced_setting"
//     name = "test"
//   }
//
// Into:
//   resource "example" "test" {
//     tier = "basic"
//     name = "test"
//   }
func ConditionalRemover(attribute string, condition func(*hclwrite.Block) bool) TransformerFunc {
	return func(block *hclwrite.Block, ctx *TransformContext) error {
		if condition(block) {
			block.Body().RemoveAttribute(attribute)
		}
		return nil
	}
}