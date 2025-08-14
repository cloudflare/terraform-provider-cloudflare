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

	transforms := map[string]ast.ExprTransformer{
		"include": transformPolicyRuleListItem,
		"exclude": transformPolicyRuleListItem,
		"require": transformPolicyRuleListItem,
	}
	ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)
	/*
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
	*/
}

// transformPolicyRuleListItem transforms condition lists:
// 1. Boolean attributes (everyone, certificate, any_valid_service_token) -> empty objects
// 2. Array attributes (email, group, ip) -> split into multiple objects
func transformPolicyRuleListItem(expr *hclsyntax.Expression, diags ast.Diagnostics) {

	tup, ok := (*expr).(*hclsyntax.TupleConsExpr)
	if !ok {
		diags.ComplicatedHCL.Append(*expr)
		return
	}

	// Build new list of objects after expansion
	var newExprs []hclsyntax.Expression

	for i := range tup.Exprs {
		obj, ok := tup.Exprs[i].(*hclsyntax.ObjectConsExpr)
		if !ok {
			// Keep non-object expressions as-is
			newExprs = append(newExprs, tup.Exprs[i])
			continue
		}

		// First, handle boolean attributes
		transforms := map[string]ast.ExprTransformer{
			"everyone":                boolToEmptyObject,
			"certificate":             boolToEmptyObject,
			"any_valid_service_token": boolToEmptyObject,
		}
		ast.ApplyTransformToAttributes(ast.NewObject(obj, diags), transforms, diags)

		// Then check if we need to expand array attributes
		expanded := expandArrayAttributes(obj, diags)
		if expanded != nil {
			// Object was expanded into multiple objects
			for _, expObj := range expanded {
				newExprs = append(newExprs, expObj)
			}
		} else {
			// No expansion needed, keep original object
			newExprs = append(newExprs, tup.Exprs[i])
		}
	}

	// Replace the tuple's expressions with the new expanded list
	tup.Exprs = newExprs
}

// expandArrayAttributes checks if an object has array attributes (email, group, ip)
// and expands them into multiple objects
func expandArrayAttributes(obj *hclsyntax.ObjectConsExpr, diags ast.Diagnostics) []hclsyntax.Expression {
	// Map of attribute names to their inner field names
	arrayAttrs := map[string]string{
		"email": "email",
		"group": "id",
		"ip":    "ip",
	}

	for _, item := range obj.Items {
		key := ast.Expr2S(item.KeyExpr, diags)
		innerFieldName, isArrayAttr := arrayAttrs[key]
		if !isArrayAttr {
			continue
		}

		// Check if the value is a tuple/array
		tup, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
		if !ok {
			continue
		}

		// Create a new object for each item in the array
		var result []hclsyntax.Expression
		for _, elem := range tup.Exprs {
			newObj := &hclsyntax.ObjectConsExpr{
				Items: []hclsyntax.ObjectConsItem{
					{
						KeyExpr: item.KeyExpr,
						ValueExpr: &hclsyntax.ObjectConsExpr{
							Items: []hclsyntax.ObjectConsItem{
								{
									KeyExpr:   ast.NewKeyExpr(innerFieldName),
									ValueExpr: elem,
								},
							},
						},
					},
				},
			}
			result = append(result, newObj)
		}
		return result
	}

	// No array attributes found
	return nil
}

func boolToEmptyObject(attrVal *hclsyntax.Expression, diags ast.Diagnostics) {
	if *attrVal == nil {
		// don't add this attribute if it doesn't already exit
		return
	}
	val := ast.Expr2S(*attrVal, diags)
	if val == "false" {
		*attrVal = nil
		return
	}

	*attrVal = &hclsyntax.ObjectConsExpr{}
}
