package main

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
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
// Handles:
// 1. Boolean attributes (everyone, certificate, any_valid_service_token) -> empty objects
// 2. Array attributes (email, group, ip, email_domain) -> split into multiple objects
// 3. Github blocks -> rename to github_organization and expand teams
//
// Example transformations:
// Before: include = [{ everyone = true }]
// After:  include = [{ everyone = {} }]
//
// Before: include = [{ email = ["user1@example.com", "user2@example.com"] }]
// After:  include = [{ email = { email = "user1@example.com" } }, { email = { email = "user2@example.com" } }]
func transformAccessPolicyBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Process include, exclude, and require attributes (grit has already converted them to lists)

	transforms := map[string]ast.ExprTransformer{
		"include": transformPolicyRuleListItem,
		"exclude": transformPolicyRuleListItem,
		"require": transformPolicyRuleListItem,
	}
	ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)
}

// transformPolicyRuleListItem transforms condition lists:
// 1. Boolean attributes (everyone, certificate, any_valid_service_token) -> empty objects
// 2. Array attributes (email, group, ip) -> split into multiple objects
// 3. Github blocks -> rename to github_organization and expand teams
//
// Example transformation:
//
//	Before: include = [{
//	  everyone = true
//	  email = ["alice@example.com", "bob@example.com"]
//	  group = ["admins"]
//	}]
//
//	After: include = [{
//	  everyone = {}
//	}, {
//
//	  email = { email = "alice@example.com" }
//	}, {
//
//	  email = { email = "bob@example.com" }
//	}, {
//
//	  group = { id = "admins" }
//	}]
func transformPolicyRuleListItem(expr *hclsyntax.Expression, diags ast.Diagnostics) {

	if *expr == nil {
		// Ignore missing attributes
		return
	}

	tup, ok := (*expr).(*hclsyntax.TupleConsExpr)
	if !ok {
		*expr = *ast.WarnManualMigration4Expr("resource/zero_trust_access_policy", expr, diags)
		return
	}

	// Build new list of objects after expansion
	var newExprs []hclsyntax.Expression

	for i := range tup.Exprs {
		obj, ok := tup.Exprs[i].(*hclsyntax.ObjectConsExpr)
		if !ok {
			// Keep non-object expressions as-is
			newExprs = append(newExprs, tup.Exprs[i])
			diags.HclDiagnostics.Append(&hcl.Diagnostic{
				Severity: hcl.DiagWarning,
				Summary:  "Unexpected non-object in condition list",
				Detail:   fmt.Sprintf("Expected ObjectConsExpr in condition list but got %T", tup.Exprs[i]),
			})
			continue
		}

		// First, handle boolean attributes
		transforms := map[string]ast.ExprTransformer{
			"everyone":                boolToEmptyObject,
			"certificate":             boolToEmptyObject,
			"any_valid_service_token": boolToEmptyObject,
		}
		// These are updated in place
		ast.ApplyTransformToAttributes(ast.NewObject(obj, diags), transforms, diags)

		// Then check if we need to expand array attributes or handle github
		expanded := expandAttributes(obj, diags)
		if len(expanded) > 0 {
			// Object was expanded into multiple objects
			newExprs = append(newExprs, expanded...)
		} else {
			// No expansion needed, keep original object
			newExprs = append(newExprs, tup.Exprs[i])
		}
	}

	// Replace the tuple's expressions with the new expanded list
	tup.Exprs = newExprs
}

// expandAttributes checks if an object has attributes that need expansion
// Returns nil if no expansion needed, or a slice of expanded objects
//
// Example transformation:
// Before: { email = ["a@ex.com", "b@ex.com"], group = ["g1"], login_method = ["okta"] }
// After: [
//
//	{ login_method = ["okta"] },  // Non-array attributes kept in first object
//	{ email = { email = "a@ex.com" } },
//	{ email = { email = "b@ex.com" } },
//	{ group = { id = "g1" } }
//
// ]
func expandAttributes(obj *hclsyntax.ObjectConsExpr, diags ast.Diagnostics) []hclsyntax.Expression {
	var allExpanded []hclsyntax.Expression
	var remainingItems []hclsyntax.ObjectConsItem

	// Note: obj.Items has already been modified by the boolean transformations
	// Items with false values have been removed, true values converted to empty objects
	for _, item := range obj.Items {
		key := ast.Expr2S(item.KeyExpr, diags)

		// Handle github specially
		if key == "github" {
			expanded := expandGithub(item, diags)
			if expanded != nil {
				allExpanded = append(allExpanded, expanded...)
				continue
			}

		}

		// Handle simple array attributes
		if expanded := expandSimpleArrayAttribute(key, item, diags); expanded != nil {
			allExpanded = append(allExpanded, expanded...)
			continue
		}

		// Keep other attributes as-is
		remainingItems = append(remainingItems, item)
	}

	// If we expanded some attributes, we need to handle remaining items
	if len(allExpanded) > 0 {
		// Each remaining item (like boolean attributes) should be its own object
		for _, item := range remainingItems {
			singleItemObj := &hclsyntax.ObjectConsExpr{
				Items: []hclsyntax.ObjectConsItem{item},
			}
			allExpanded = append(allExpanded, singleItemObj)
		}
		return allExpanded
	}

	// No expansion happened
	return nil
}

// expandSimpleArrayAttribute handles email, group, ip, email_domain arrays
//
// Example transformations:
// Before: email = ["alice@example.com", "bob@example.com"]
// After: [
//
//	{ email = { email = "alice@example.com" } },
//	{ email = { email = "bob@example.com" } }
//
// ]
//
// Before: group = ["group-id-1", "group-id-2"]
// After: [
//
//	{ group = { id = "group-id-1" } },
//	{ group = { id = "group-id-2" } }
//
// ]
func expandSimpleArrayAttribute(key string, item hclsyntax.ObjectConsItem, diags ast.Diagnostics) []hclsyntax.Expression {
	// Map of attribute names to their inner field names
	arrayAttrs := map[string]string{
		"email":        "email",
		"group":        "id",
		"ip":           "ip",
		"email_domain": "domain",
		"geo":          "country_code",
	}

	innerFieldName, isArrayAttr := arrayAttrs[key]
	if !isArrayAttr {
		return nil
	}

	// Check if the value is a tuple/array
	tup, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
	if !ok {
		// Not an array, keep as-is
		return nil
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

// expandGithub handles the special case of github blocks
// V4: github = [{ name = "org", teams = ["team1", "team2"], identity_provider_id = "id" }]
// V5: Multiple github_organization blocks, one per team
//
// Example transformation:
//
//	Before: github = [{
//	  name = "my-org"
//	  teams = ["engineering", "devops"]
//	  identity_provider_id = "provider-123"
//	}]
//
// After: [
//
//	{ github_organization = { name = "my-org", team = "engineering", identity_provider_id = "provider-123" } },
//	{ github_organization = { name = "my-org", team = "devops", identity_provider_id = "provider-123" } }
//
// ]
func expandGithub(item hclsyntax.ObjectConsItem, diags ast.Diagnostics) []hclsyntax.Expression {
	// Check if the value is a tuple/array of github blocks
	tup, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
	if !ok {
		diags.HclDiagnostics.Append(&hcl.Diagnostic{
			Severity: hcl.DiagWarning,
			Summary:  "Failed to cast github value to TupleConsExpr",
			Detail:   fmt.Sprintf("Expected github value to be an array but got %T", item.ValueExpr),
		})
		return nil
	}

	var result []hclsyntax.Expression

	for _, githubExpr := range tup.Exprs {
		githubObj, ok := githubExpr.(*hclsyntax.ObjectConsExpr)
		if !ok {
			diags.HclDiagnostics.Append(&hcl.Diagnostic{
				Severity: hcl.DiagWarning,
				Summary:  "Failed to cast github array element to ObjectConsExpr",
				Detail:   fmt.Sprintf("Expected github array element to be an object but got %T", githubExpr),
			})
			continue
		}

		// Extract the github block fields
		var nameExpr hclsyntax.Expression
		var teamsExpr *hclsyntax.TupleConsExpr
		var identityProviderExpr hclsyntax.Expression
		var otherItems []hclsyntax.ObjectConsItem

		for _, githubItem := range githubObj.Items {
			itemKey := ast.Expr2S(githubItem.KeyExpr, diags)
			switch itemKey {
			case "name":
				nameExpr = githubItem.ValueExpr
			case "teams":
				teamsExpr, ok = githubItem.ValueExpr.(*hclsyntax.TupleConsExpr)
				if !ok && githubItem.ValueExpr != nil {
					diags.HclDiagnostics.Append(&hcl.Diagnostic{
						Severity: hcl.DiagWarning,
						Summary:  "Failed to cast teams value to TupleConsExpr",
						Detail:   fmt.Sprintf("Expected teams to be an array but got %T", githubItem.ValueExpr),
					})
				}
			case "identity_provider_id":
				identityProviderExpr = githubItem.ValueExpr
			default:
				otherItems = append(otherItems, githubItem)
			}
		}

		// If there's a teams array, expand it
		if teamsExpr != nil && len(teamsExpr.Exprs) > 0 {
			for _, teamExpr := range teamsExpr.Exprs {
				// Build the new github_organization object
				var items []hclsyntax.ObjectConsItem

				if nameExpr != nil {
					items = append(items, hclsyntax.ObjectConsItem{
						KeyExpr:   ast.NewKeyExpr("name"),
						ValueExpr: nameExpr,
					})
				}

				// Convert teams array element to single team
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr:   ast.NewKeyExpr("team"),
					ValueExpr: teamExpr,
				})

				if identityProviderExpr != nil {
					items = append(items, hclsyntax.ObjectConsItem{
						KeyExpr:   ast.NewKeyExpr("identity_provider_id"),
						ValueExpr: identityProviderExpr,
					})
				}

				// Add any other fields
				items = append(items, otherItems...)

				newObj := &hclsyntax.ObjectConsExpr{
					Items: []hclsyntax.ObjectConsItem{
						{
							KeyExpr: ast.NewKeyExpr("github_organization"),
							ValueExpr: &hclsyntax.ObjectConsExpr{
								Items: items,
							},
						},
					},
				}
				result = append(result, newObj)
			}
		} else {
			// No teams array or empty teams, create single github_organization
			var items []hclsyntax.ObjectConsItem

			if nameExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr:   ast.NewKeyExpr("name"),
					ValueExpr: nameExpr,
				})
			}

			if identityProviderExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr:   ast.NewKeyExpr("identity_provider_id"),
					ValueExpr: identityProviderExpr,
				})
			}

			// Add any other fields
			items = append(items, otherItems...)

			newObj := &hclsyntax.ObjectConsExpr{
				Items: []hclsyntax.ObjectConsItem{
					{
						KeyExpr: ast.NewKeyExpr("github_organization"),
						ValueExpr: &hclsyntax.ObjectConsExpr{
							Items: items,
						},
					},
				},
			}
			result = append(result, newObj)
		}
	}

	return result
}

// boolToEmptyObject transforms boolean attributes to empty objects or removes them
//
// Example transformations:
// Before: everyone = true
// After:  everyone = {}
//
// Before: certificate = false
// After:  (attribute removed entirely)
func boolToEmptyObject(attrVal *hclsyntax.Expression, diags ast.Diagnostics) {
	if *attrVal == nil {
		// don't add this attribute if it doesn't already exist
		return
	}
	val := ast.Expr2S(*attrVal, diags)
	if val == "false" {
		*attrVal = nil
		return
	}

	*attrVal = &hclsyntax.ObjectConsExpr{}
}
