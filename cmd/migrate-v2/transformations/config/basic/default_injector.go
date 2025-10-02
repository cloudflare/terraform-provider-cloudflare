package basic

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// DefaultValueSetter sets default values for missing attributes
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