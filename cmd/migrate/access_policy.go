package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isAccessPolicyResource checks if a block is a cloudflare_zero_trust_access_policy resource
// (grit has already renamed from cloudflare_access_policy)
func isAccessPolicyResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_zero_trust_access_policy"
}

// transformAccessPolicyBlock transforms include/exclude/require attributes
// Currently only handles boolean attributes like everyone, certificate, any_valid_service_token
// After grit: include = [{ everyone = true }]
// After this: include = [{ everyone = {} }]
func transformAccessPolicyBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Process include, exclude, and require attributes (grit has already converted them to lists)
	conditionAttributes := []string{"include", "exclude", "require"}

	for _, attrName := range conditionAttributes {
		attr := block.Body().GetAttribute(attrName)
		if attr == nil {
			continue // Attribute doesn't exist
		}

		transformedTokens := transformBooleanAttributesInList(*attr, diags)
		if transformedTokens != nil {
			block.Body().SetAttributeRaw(attrName, transformedTokens)
		}
	}
}

// transformBooleanAttributesInList transforms boolean attributes within a condition list
// Only handles: everyone, certificate, any_valid_service_token
func transformBooleanAttributesInList(attr hclwrite.Attribute, diags ast.Diagnostics) hclwrite.Tokens {
	expr := ast.WriteExpr2Expr(*attr.Expr(), diags)

	tup, ok := expr.(*hclsyntax.TupleConsExpr)
	if !ok {
		diags.ComplicatedHCL.Append((expr))
		return nil
	}
	for _, expr := range tup.Exprs {
		obj, ok := expr.(*hclsyntax.ObjectConsExpr)
		if !ok {
			return nil
		}

		transforms := map[string]ast.ExprTransformer{
			"everyone":                boolToEmptyObject,
			"certificate":             boolToEmptyObject,
			"any_valid_service_token": boolToEmptyObject,
		}

		ast.ApplyTransformToAttributes(obj, transforms, diags)
	}

	write := ast.Expr2WriteExpr(tup, diags)
	return write.BuildTokens(nil)
}

func boolToEmptyObject(attrVal *hclsyntax.Expression, diags ast.Diagnostics) {
	if *attrVal == nil {
		// don't add this attribute if it doesn't already exit
		return
	}
	val := ast.Expr2S(*attrVal, diags)
	if val == "false" {
		*attrVal = nil
	}

	*attrVal = &hclsyntax.ObjectConsExpr{}
}
