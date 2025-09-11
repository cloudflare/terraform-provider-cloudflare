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
//
// Phase 3 - Remove unsupported attributes and convert blocks to attributes:
// Before: domain_type = "public" (removed - not supported in v5)
// Before: destinations { ... } (converted from blocks to list attribute)
func transformAccessApplicationBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// STEP 1: Auto-populate policies from collected access_policy mappings
	injectCollectedPolicies(block, diags)
	
	// STEP 2: Remove unsupported attributes
	removeUnsupportedAttributes(block)
	
	// STEP 3: Convert destinations blocks to list attribute
	convertDestinationsBlocksToAttribute(block, diags)
	
	// STEP 4: Transform existing attribute formats
	transforms := map[string]ast.ExprTransformer{
		"policies":       transformPoliciesAttribute,
		"allowed_idps":   transformSetToListAttribute,
		"custom_pages":   transformSetToListAttribute,
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

// transformSetToListAttribute converts set syntax to list syntax
// This is needed for attributes like allowed_idps and custom_pages that changed from set to list in v5
//
// The transformation preserves the values but changes the syntax:
// Before: toset(["value1", "value2"])  (v4 set)
// After:  ["value1", "value2"]         (v5 list)
func transformSetToListAttribute(expr *hclsyntax.Expression, diags ast.Diagnostics) {
	if *expr == nil {
		return
	}

	// Check if it's a function call (e.g., toset(...))
	if funcCall, ok := (*expr).(*hclsyntax.FunctionCallExpr); ok {
		if funcCall.Name == "toset" && len(funcCall.Args) == 1 {
			// Replace with the inner list expression
			*expr = funcCall.Args[0]
			return
		}
	}

	// If it's already a tuple/list, leave it as-is
	if _, ok := (*expr).(*hclsyntax.TupleConsExpr); ok {
		return
	}

	// If we can't identify the pattern, leave it unchanged
	// (grit should handle most cases)
}

// removeUnsupportedAttributes removes attributes that don't exist in v5 or are conditionally invalid
// The domain_type attribute was removed in v5 as it's no longer supported
// The skip_app_launcher_login_page can only be set when type = "app_launcher"
func removeUnsupportedAttributes(block *hclwrite.Block) {
	body := block.Body()
	
	// Remove domain_type attribute - not supported in v5
	body.RemoveAttribute("domain_type")
	
	// Remove skip_app_launcher_login_page if type is not "app_launcher"
	typeAttr := body.GetAttribute("type")
	if typeAttr != nil {
		// Get the type value to check if it's "app_launcher"
		tokens := typeAttr.Expr().BuildTokens(nil)
		var typeValue string
		for _, token := range tokens {
			typeValue += string(token.Bytes)
		}
		
		// Remove quotes if present and check the value
		typeValue = strings.Trim(typeValue, `"`)
		if typeValue != "app_launcher" {
			// Remove skip_app_launcher_login_page attribute when type is not app_launcher
			body.RemoveAttribute("skip_app_launcher_login_page")
		}
	} else {
		// If no type attribute, remove skip_app_launcher_login_page as it would be invalid
		body.RemoveAttribute("skip_app_launcher_login_page")
	}
}

// convertDestinationsBlocksToAttribute converts destinations blocks to a list attribute
// This handles the change from:
//   destinations { ... }
//   destinations { ... }
// To:
//   destinations = [{ ... }, { ... }]
func convertDestinationsBlocksToAttribute(block *hclwrite.Block, diags ast.Diagnostics) {
	body := block.Body()
	
	// Find all destinations blocks
	var destBlocks []*hclwrite.Block
	for _, childBlock := range body.Blocks() {
		if childBlock.Type() == "destinations" {
			destBlocks = append(destBlocks, childBlock)
		}
	}
	
	if len(destBlocks) == 0 {
		return // No destinations blocks to convert
	}
	
	// Build HCL string for the destinations list attribute
	var destObjects []string
	for _, destBlock := range destBlocks {
		// Convert each block to an object representation
		objectStr := convertDestinationBlockToObject(destBlock)
		if objectStr != "" {
			destObjects = append(destObjects, objectStr)
		}
	}
	
	if len(destObjects) > 0 {
		// Create the destinations attribute HCL
		destAttrHCL := fmt.Sprintf(`destinations = [
  %s
]`, strings.Join(destObjects, ",\n  "))
		
		// Parse the HCL and extract the attribute
		file, parseDiags := hclwrite.ParseConfig([]byte(destAttrHCL), "", hcl.InitialPos)
		if parseDiags.HasErrors() {
			// If we can't parse, add a warning comment instead
			body.AppendNewline()
			commentTokens := []*hclwrite.Token{
				{Type: hclsyntax.TokenComment, Bytes: []byte("# TODO: Manual migration required for destinations blocks\n")},
			}
			body.AppendUnstructuredTokens(commentTokens)
			return
		}
		
		// Find the destinations attribute in the parsed file and copy it
		for name, attr := range file.Body().Attributes() {
			if name == "destinations" {
				body.SetAttributeRaw("destinations", attr.Expr().BuildTokens(nil))
				break
			}
		}
	}
	
	// Remove all destinations blocks after converting to attribute
	for _, destBlock := range destBlocks {
		body.RemoveBlock(destBlock)
	}
}

// convertDestinationBlockToObject converts a destinations block to its object representation
// Converts from:
//   destinations {
//     uri = "https://example.com"
//   }
// To: 
//   { uri = "https://example.com" }
func convertDestinationBlockToObject(block *hclwrite.Block) string {
	if block == nil {
		return ""
	}
	
	// Get all attributes from the block
	attrs := block.Body().Attributes()
	if len(attrs) == 0 {
		return "{}" // Empty object
	}
	
	// Build attribute strings
	var attrStrings []string
	for name, attr := range attrs {
		// Get the raw tokens for the expression to preserve references and formatting
		tokens := attr.Expr().BuildTokens(nil)
		var exprStr string
		for _, token := range tokens {
			exprStr += string(token.Bytes)
		}
		attrStrings = append(attrStrings, fmt.Sprintf("    %s = %s", name, exprStr))
	}
	
	// Sort for consistent output
	sort.Strings(attrStrings)
	
	return fmt.Sprintf("{\n%s\n  }", strings.Join(attrStrings, "\n"))
}