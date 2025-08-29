package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// isAccessApplicationResource checks if a block is a cloudflare_zero_trust_access_application resource
// (grit has already renamed from cloudflare_access_application)
func isAccessApplicationResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_zero_trust_access_application"
}

// transformAccessApplicationBlock transforms policies attribute and auto-populates from collected policy mappings
//
// Example transformations:
// Phase 1 - Existing policies attribute:
// Before: policies = ["id1", "id2"]
// After:  policies = [{id = "id1"}, {id = "id2"}]
//
// Phase 2 - Auto-inject collected policies from access_policy blocks:
// Before: (no policies attribute, but collected from access_policy blocks)
// After:  policies = [{id = cloudflare_zero_trust_access_policy.allow.id, precedence = 1}, {id = cloudflare_zero_trust_access_policy.deny.id, precedence = 2}]
func transformAccessApplicationBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// STEP 1: Auto-populate policies from collected access_policy mappings
	injectCollectedPolicies(block, diags)
	
	// STEP 2: Transform existing policies attribute format
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

// injectCollectedPolicies automatically adds policies to applications based on collected mappings from access_policy blocks
func injectCollectedPolicies(block *hclwrite.Block, diags ast.Diagnostics) {
	if len(block.Labels()) < 2 {
		return
	}
	
	// Construct the application reference that policies would have used
	// E.g., "cloudflare_zero_trust_access_application.app.id"
	applicationRef := fmt.Sprintf("%s.%s.id", block.Labels()[0], block.Labels()[1])
	
	// Check if we have collected policies for this application
	policies, exists := applicationPolicyMapping[applicationRef]
	if !exists || len(policies) == 0 {
		return
	}
	
	// Sort policies by precedence to maintain order
	sort.Slice(policies, func(i, j int) bool {
		return policies[i].Precedence < policies[j].Precedence
	})
	
	body := block.Body()
	
	// Check if policies attribute already exists
	existingPolicies := body.GetAttribute("policies")
	if existingPolicies != nil {
		// TODO: Merge with existing policies if needed
		// For now, we'll let the existing policies take precedence
		return
	}
	
	// Create the policies attribute using raw tokens to preserve references
	createPoliciesAttribute(body, policies)
}

// createPoliciesAttribute creates a policies attribute with proper HCL references
// Example output: policies = [{id = cloudflare_zero_trust_access_policy.allow.id, precedence = 1}]
func createPoliciesAttribute(body *hclwrite.Body, policies []PolicyReference) {
	if len(policies) == 0 {
		return
	}
	
	// Build HCL string for the policies attribute
	var policyStrings []string
	for _, policy := range policies {
		policyRef := fmt.Sprintf("%s.id", policy.ResourceName)
		policyString := fmt.Sprintf(`{
    id = %s
    precedence = %d
  }`, policyRef, policy.Precedence)
		policyStrings = append(policyStrings, policyString)
	}
	
	// Create the full policies attribute HCL
	fullAttrHCL := fmt.Sprintf(`policies = [
  %s
]`, strings.Join(policyStrings, ",\n  "))
	
	// Parse the HCL and extract the attribute
	file, diags := hclwrite.ParseConfig([]byte(fullAttrHCL), "", hcl.InitialPos)
	if diags.HasErrors() {
		// Fallback: create a simple string-based attribute
		body.SetAttributeValue("policies", cty.ListValEmpty(cty.String))
		return
	}
	
	// Find the policies attribute in the parsed file
	for name, attr := range file.Body().Attributes() {
		if name == "policies" {
			// Copy the attribute to our target body
			body.SetAttributeRaw("policies", attr.Expr().BuildTokens(nil))
			break
		}
	}
	
	// Add explanatory comment
	body.AppendNewline()
	commentTokens := []*hclwrite.Token{
		{Type: hclsyntax.TokenComment, Bytes: []byte("# Policies auto-migrated from v4 access_policy resources\n")},
	}
	body.AppendUnstructuredTokens(commentTokens)
}