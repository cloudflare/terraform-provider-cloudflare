package basic

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// AttributeRemover creates a transformer that removes attributes
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
func ConditionalRemover(attribute string, condition func(*hclwrite.Block) bool) TransformerFunc {
	return func(block *hclwrite.Block, ctx *TransformContext) error {
		if condition(block) {
			block.Body().RemoveAttribute(attribute)
		}
		return nil
	}
}