package basic

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// AttributeRenamer creates a transformer that renames attributes
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