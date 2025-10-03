package basic

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// DefaultValueSetter sets default values for missing attributes
//
// Example YAML configuration:
//   defaults:
//     enabled: true
//     timeout: 30
//     environment: "production"
//     retries: 3
//
// Transforms:
//   resource "example" "test" {
//     name = "test"
//     timeout = 60  # Existing value preserved
//   }
//
// Into:
//   resource "example" "test" {
//     name = "test"
//     timeout = 60         # Existing value preserved
//     enabled = true       # Default added
//     environment = "production"  # Default added
//     retries = 3          # Default added
//   }
func DefaultValueSetter(defaults map[string]interface{}) TransformerFunc {
	return func(block *hclwrite.Block, ctx *TransformContext) error {
		body := block.Body()
		
		for attrName, defaultValue := range defaults {
			if body.GetAttribute(attrName) == nil {
				switch v := defaultValue.(type) {
				case string:
					body.SetAttributeValue(attrName, cty.StringVal(v))
				case int:
					body.SetAttributeValue(attrName, cty.NumberIntVal(int64(v)))
				case bool:
					body.SetAttributeValue(attrName, cty.BoolVal(v))
				}
			}
		}
		
		return nil
	}
}