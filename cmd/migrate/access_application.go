package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isAccessApplicationResource checks if a block is a cloudflare_zero_trust_access_application resource
// (grit has already renamed from cloudflare_access_application)
func isAccessApplicationResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_zero_trust_access_application"
}

// transformAccessApplicationBlock transforms policies attribute from list of strings to list of objects
//
// Example transformations:
// Before: policies = ["id1", "id2"]
// After:  policies = [{id = "id1"}, {id = "id2"}]
//
// Before: policies = [cloudflare_zero_trust_access_policy.example.id]
// After:  policies = [{id = cloudflare_zero_trust_access_policy.example.id}]
func transformAccessApplicationBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	transforms := map[string]ast.ExprTransformer{
		"policies": transformPoliciesAttribute,
	}
	ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)
}

// transformPoliciesAttribute transforms a policies list from strings/references to objects
//
// Example transformations:
// Before: ["policy-id-1", "policy-id-2"]
// After:  [{id = "policy-id-1"}, {id = "policy-id-2"}]
//
// Before: [cloudflare_zero_trust_access_policy.allow.id, cloudflare_zero_trust_access_policy.deny.id]
// After:  [{id = cloudflare_zero_trust_access_policy.allow.id}, {id = cloudflare_zero_trust_access_policy.deny.id}]
func transformPoliciesAttribute(expr *hclsyntax.Expression, diags ast.Diagnostics) {
	if *expr == nil {
		// No policies attribute
		return
	}

	// Check if it's a tuple (list)
	tup, ok := (*expr).(*hclsyntax.TupleConsExpr)
	if !ok {
		// Not a list, can't transform
		diags.ComplicatedHCL.Append(*expr)
		return
	}

	// Transform each element to {id = element}
	var newExprs []hclsyntax.Expression
	for _, elem := range tup.Exprs {
		// Create an object with a single "id" attribute
		newObj := &hclsyntax.ObjectConsExpr{
			Items: []hclsyntax.ObjectConsItem{
				{
					KeyExpr:   ast.NewKeyExpr("id"),
					ValueExpr: elem,
				},
			},
		}
		newExprs = append(newExprs, newObj)
	}

	// Replace the tuple's expressions with the new objects
	tup.Exprs = newExprs
}