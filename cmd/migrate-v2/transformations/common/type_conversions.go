package common

import (
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// SetToListConverter converts set attributes to lists
func SetToListConverter(attributes ...string) TransformerFunc {
	return func(block *hclwrite.Block, ctx *TransformContext) error {
		body := block.Body()
		
		for _, attrName := range attributes {
			attr := body.GetAttribute(attrName)
			if attr == nil {
				continue
			}
			
			tokens := attr.Expr().BuildTokens(nil)
			str := strings.TrimSpace(string(tokens.Bytes()))
			
			// Remove toset() wrapper if present
			if strings.HasPrefix(str, "toset(") && strings.HasSuffix(str, ")") {
				content := str[6:len(str)-1]
				newTokens := hclwrite.Tokens{
					&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(content)},
				}
				body.SetAttributeRaw(attrName, newTokens)
			}
		}
		
		return nil
	}
}