package main

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isAccessGroupResource checks if a block is a cloudflare_zero_trust_access_group resource
// (grit has already renamed from cloudflare_access_group)
func isAccessGroupResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_zero_trust_access_group"
}

// transformAccessGroupBlock transforms include/exclude/require attributes
// Handles:
// 1. Boolean attributes (everyone, certificate, any_valid_service_token) -> empty objects
// 2. Array attributes (email, group, ip, email_domain, etc.) -> split into multiple objects
// 3. Azure blocks -> rename to azure_ad and expand ID arrays
// 4. Github blocks -> rename to github_organization and expand teams
// 5. GSuite blocks -> expand email arrays
// 6. Okta blocks -> expand name arrays
// 7. Common names arrays -> split into multiple common_name objects
//
// Example transformations:
// Before: include = [{ everyone = true }]
// After:  include = [{ everyone = {} }]
//
// Before: include = [{ email = ["user1@example.com", "user2@example.com"] }]
// After:  include = [{ email = { email = "user1@example.com" } }, { email = { email = "user2@example.com" } }]
//
// Before: include = [{ azure = { id = ["group1", "group2"], identity_provider_id = "provider" } }]
// After:  include = [{ azure_ad = { id = "group1", identity_provider_id = "provider" } }, { azure_ad = { id = "group2", identity_provider_id = "provider" } }]
//
// Before: include = [{ github = { name = "org", teams = ["team1", "team2"], identity_provider_id = "provider" } }]
// After:  include = [{ github_organization = { name = "org", team = "team1", identity_provider_id = "provider" } }, { github_organization = { name = "org", team = "team2", identity_provider_id = "provider" } }]
func transformAccessGroupBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	transforms := map[string]ast.ExprTransformer{
		"include": transformGroupRuleListItem,
		"exclude": transformGroupRuleListItem,
		"require": transformGroupRuleListItem,
	}
	ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)
}

// transformGroupRuleListItem transforms condition lists for access groups
func transformGroupRuleListItem(expr *hclsyntax.Expression, diags ast.Diagnostics) {
	if *expr == nil {
		// Ignore missing attributes
		return
	}

	tup, ok := (*expr).(*hclsyntax.TupleConsExpr)
	if !ok {
		*expr = *ast.WarnManualMigration4Expr("resource/zero_trust_access_group", expr, diags)
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

		// Expand this single rule object into multiple rule objects
		expandedObjs := expandRuleObjectInAccessGroup(obj, diags)
		newExprs = append(newExprs, expandedObjs...)
	}

	// Replace original tuple with new expanded tuple
	*expr = &hclsyntax.TupleConsExpr{
		Exprs: newExprs,
	}
}

// expandRuleObjectInAccessGroup expands a single rule object into multiple rule objects  
// Each attribute that contains arrays or needs transformation becomes a separate rule object
func expandRuleObjectInAccessGroup(obj *hclsyntax.ObjectConsExpr, diags ast.Diagnostics) []hclsyntax.Expression {
	var newObjs []hclsyntax.Expression

	// Helper to create new object with single attribute
	createSingleAttrObject := func(attrName string, value hclsyntax.Expression, subAttrName string) hclsyntax.Expression {
		return &hclsyntax.ObjectConsExpr{
			Items: []hclsyntax.ObjectConsItem{
				{
					KeyExpr: ast.NewKeyExpr(attrName),
					ValueExpr: &hclsyntax.ObjectConsExpr{
						Items: []hclsyntax.ObjectConsItem{
							{
								KeyExpr:   ast.NewKeyExpr(subAttrName),
								ValueExpr: value,
							},
						},
					},
				},
			},
		}
	}

	// Helper to create new object with multiple sub-attributes
	createMultiAttrObject := func(attrName string, subAttrs map[string]hclsyntax.Expression) hclsyntax.Expression {
		var items []hclsyntax.ObjectConsItem
		for key, value := range subAttrs {
			items = append(items, hclsyntax.ObjectConsItem{
				KeyExpr:   ast.NewKeyExpr(key),
				ValueExpr: value,
			})
		}
		return &hclsyntax.ObjectConsExpr{
			Items: []hclsyntax.ObjectConsItem{
				{
					KeyExpr: ast.NewKeyExpr(attrName),
					ValueExpr: &hclsyntax.ObjectConsExpr{
						Items: items,
					},
				},
			},
		}
	}

	// Build a map of attributes for consistent ordering
	attrMap := make(map[string]hclsyntax.ObjectConsItem)
	for _, item := range obj.Items {
		keyStr := ast.Expr2S(item.KeyExpr, diags)
		attrMap[keyStr] = item
	}

	// Process attributes in the same fixed order as state migration to ensure consistency
	attributeOrder := []string{
		"email", "email_domain", "ip", "geo", "group", "service_token", "email_list", "ip_list", "login_method", "device_posture",
		"common_names", "azure", "github", "gsuite", "okta",
		"everyone", "certificate", "any_valid_service_token",
	}

	for _, keyStr := range attributeOrder {
		item, exists := attrMap[keyStr]
		if !exists {
			continue
		}

		switch keyStr {
		// Boolean attributes - create empty object rules
		case "everyone", "certificate", "any_valid_service_token":
			newObj := &hclsyntax.ObjectConsExpr{
				Items: []hclsyntax.ObjectConsItem{
					{
						KeyExpr:   ast.NewKeyExpr(keyStr),
						ValueExpr: &hclsyntax.ObjectConsExpr{Items: []hclsyntax.ObjectConsItem{}},
					},
				},
			}
			newObjs = append(newObjs, newObj)

		// Array attributes - expand into multiple objects with nested structure
		case "email", "email_domain", "ip", "geo", "group", "service_token", "email_list", "ip_list", "login_method", "device_posture":
			tup, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
			if !ok {
				// Not a tuple, add warning
				diags.HclDiagnostics.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  fmt.Sprintf("Expected list for %s attribute", keyStr),
					Detail:   fmt.Sprintf("Expected TupleConsExpr but got %T", item.ValueExpr),
				})
				continue
			}

			// Determine the sub-attribute name based on the main attribute
			var subAttrName string
			switch keyStr {
			case "email":
				subAttrName = "email"
			case "email_domain":
				subAttrName = "domain"
			case "ip":
				subAttrName = "ip"
			case "geo":
				subAttrName = "country_code"
			case "group":
				subAttrName = "id"
			case "service_token":
				subAttrName = "token_id"
			case "email_list":
				subAttrName = "id"
			case "ip_list":
				subAttrName = "id"
			case "login_method":
				subAttrName = "id"
			case "device_posture":
				subAttrName = "integration_uid"
			}

			// Create multiple objects, one for each array element
			for _, valueExpr := range tup.Exprs {
				newObj := createSingleAttrObject(keyStr, valueExpr, subAttrName)
				newObjs = append(newObjs, newObj)
			}

		case "common_names":
			// common_names array -> multiple common_name objects
			tup, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
			if !ok {
				diags.HclDiagnostics.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "Expected list for common_names attribute",
					Detail:   fmt.Sprintf("Expected TupleConsExpr but got %T", item.ValueExpr),
				})
				continue
			}

			for _, valueExpr := range tup.Exprs {
				newObj := createSingleAttrObject("common_name", valueExpr, "common_name")
				newObjs = append(newObjs, newObj)
			}

		case "azure":
			// Handle azure blocks -> rename to azure_ad and expand ID arrays
			azureBlocks, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
			if !ok {
				diags.HclDiagnostics.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "Expected list for azure blocks",
					Detail:   fmt.Sprintf("Expected TupleConsExpr but got %T", item.ValueExpr),
				})
				continue
			}

			for _, azureBlockExpr := range azureBlocks.Exprs {
				azureObj, ok := azureBlockExpr.(*hclsyntax.ObjectConsExpr)
				if !ok {
					diags.HclDiagnostics.Append(&hcl.Diagnostic{
						Severity: hcl.DiagWarning,
						Summary:  "Expected object for azure block",
						Detail:   fmt.Sprintf("Expected ObjectConsExpr but got %T", azureBlockExpr),
					})
					continue
				}

				var identityProviderID hclsyntax.Expression
				var idArray *hclsyntax.TupleConsExpr

				// Extract identity_provider_id and id array from azure block
				for _, azureItem := range azureObj.Items {
					azureKey := ast.Expr2S(azureItem.KeyExpr, diags)
					if azureKey == "identity_provider_id" {
						identityProviderID = azureItem.ValueExpr
					} else if azureKey == "id" {
						if tup, ok := azureItem.ValueExpr.(*hclsyntax.TupleConsExpr); ok {
							idArray = tup
						}
					}
				}

				// Create multiple azure_ad objects, one for each ID
				if idArray != nil {
					for _, idExpr := range idArray.Exprs {
						subAttrs := map[string]hclsyntax.Expression{
							"id": idExpr,
						}
						if identityProviderID != nil {
							subAttrs["identity_provider_id"] = identityProviderID
						}
						newObj := createMultiAttrObject("azure_ad", subAttrs)
						newObjs = append(newObjs, newObj)
					}
				}
			}

		case "github":
			// Handle github blocks -> rename to github_organization and expand teams
			githubBlocks, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
			if !ok {
				diags.HclDiagnostics.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "Expected list for github blocks",
					Detail:   fmt.Sprintf("Expected TupleConsExpr but got %T", item.ValueExpr),
				})
				continue
			}

			for _, githubBlockExpr := range githubBlocks.Exprs {
				githubObj, ok := githubBlockExpr.(*hclsyntax.ObjectConsExpr)
				if !ok {
					diags.HclDiagnostics.Append(&hcl.Diagnostic{
						Severity: hcl.DiagWarning,
						Summary:  "Expected object for github block",
						Detail:   fmt.Sprintf("Expected ObjectConsExpr but got %T", githubBlockExpr),
					})
					continue
				}

				var name, identityProviderID hclsyntax.Expression
				var teamsArray *hclsyntax.TupleConsExpr

				// Extract name, identity_provider_id, and teams array from github block
				for _, githubItem := range githubObj.Items {
					githubKey := ast.Expr2S(githubItem.KeyExpr, diags)
					switch githubKey {
					case "name":
						name = githubItem.ValueExpr
					case "identity_provider_id":
						identityProviderID = githubItem.ValueExpr
					case "teams":
						if tup, ok := githubItem.ValueExpr.(*hclsyntax.TupleConsExpr); ok {
							teamsArray = tup
						}
					}
				}

				// Create multiple github_organization objects, one for each team
				if teamsArray != nil {
					for _, teamExpr := range teamsArray.Exprs {
						subAttrs := map[string]hclsyntax.Expression{
							"team": teamExpr,
						}
						if name != nil {
							subAttrs["name"] = name
						}
						if identityProviderID != nil {
							subAttrs["identity_provider_id"] = identityProviderID
						}
						newObj := createMultiAttrObject("github_organization", subAttrs)
						newObjs = append(newObjs, newObj)
					}
				}
			}

		case "gsuite":
			// Handle gsuite blocks -> expand email arrays
			gsuiteBlocks, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
			if !ok {
				diags.HclDiagnostics.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "Expected list for gsuite blocks",
					Detail:   fmt.Sprintf("Expected TupleConsExpr but got %T", item.ValueExpr),
				})
				continue
			}

			for _, gsuiteBlockExpr := range gsuiteBlocks.Exprs {
				gsuiteObj, ok := gsuiteBlockExpr.(*hclsyntax.ObjectConsExpr)
				if !ok {
					diags.HclDiagnostics.Append(&hcl.Diagnostic{
						Severity: hcl.DiagWarning,
						Summary:  "Expected object for gsuite block",
						Detail:   fmt.Sprintf("Expected ObjectConsExpr but got %T", gsuiteBlockExpr),
					})
					continue
				}

				var identityProviderID hclsyntax.Expression
				var emailArray *hclsyntax.TupleConsExpr

				// Extract identity_provider_id and email array from gsuite block
				for _, gsuiteItem := range gsuiteObj.Items {
					gsuiteKey := ast.Expr2S(gsuiteItem.KeyExpr, diags)
					if gsuiteKey == "identity_provider_id" {
						identityProviderID = gsuiteItem.ValueExpr
					} else if gsuiteKey == "email" {
						if tup, ok := gsuiteItem.ValueExpr.(*hclsyntax.TupleConsExpr); ok {
							emailArray = tup
						}
					}
				}

				// Create multiple gsuite objects, one for each email
				if emailArray != nil {
					for _, emailExpr := range emailArray.Exprs {
						subAttrs := map[string]hclsyntax.Expression{
							"email": emailExpr,
						}
						if identityProviderID != nil {
							subAttrs["identity_provider_id"] = identityProviderID
						}
						newObj := createMultiAttrObject("gsuite", subAttrs)
						newObjs = append(newObjs, newObj)
					}
				}
			}

		case "okta":
			// Handle okta blocks -> expand name arrays
			oktaBlocks, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
			if !ok {
				diags.HclDiagnostics.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "Expected list for okta blocks",
					Detail:   fmt.Sprintf("Expected TupleConsExpr but got %T", item.ValueExpr),
				})
				continue
			}

			for _, oktaBlockExpr := range oktaBlocks.Exprs {
				oktaObj, ok := oktaBlockExpr.(*hclsyntax.ObjectConsExpr)
				if !ok {
					diags.HclDiagnostics.Append(&hcl.Diagnostic{
						Severity: hcl.DiagWarning,
						Summary:  "Expected object for okta block",
						Detail:   fmt.Sprintf("Expected ObjectConsExpr but got %T", oktaBlockExpr),
					})
					continue
				}

				var identityProviderID hclsyntax.Expression
				var nameArray *hclsyntax.TupleConsExpr

				// Extract identity_provider_id and name array from okta block
				for _, oktaItem := range oktaObj.Items {
					oktaKey := ast.Expr2S(oktaItem.KeyExpr, diags)
					if oktaKey == "identity_provider_id" {
						identityProviderID = oktaItem.ValueExpr
					} else if oktaKey == "name" {
						if tup, ok := oktaItem.ValueExpr.(*hclsyntax.TupleConsExpr); ok {
							nameArray = tup
						}
					}
				}

				// Create multiple okta objects, one for each name
				if nameArray != nil {
					for _, nameExpr := range nameArray.Exprs {
						subAttrs := map[string]hclsyntax.Expression{
							"name": nameExpr,
						}
						if identityProviderID != nil {
							subAttrs["identity_provider_id"] = identityProviderID
						}
						newObj := createMultiAttrObject("okta", subAttrs)
						newObjs = append(newObjs, newObj)
					}
				}
			}

		case "saml":
			// Handle saml blocks (no expansion needed, just keep as-is)
			samlBlocks, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
			if !ok {
				diags.HclDiagnostics.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "Expected list for saml blocks",
					Detail:   fmt.Sprintf("Expected TupleConsExpr but got %T", item.ValueExpr),
				})
				continue
			}

			for _, samlBlockExpr := range samlBlocks.Exprs {
				newObj := &hclsyntax.ObjectConsExpr{
					Items: []hclsyntax.ObjectConsItem{
						{
							KeyExpr:   ast.NewKeyExpr("saml"),
							ValueExpr: samlBlockExpr,
						},
					},
				}
				newObjs = append(newObjs, newObj)
			}

		case "external_evaluation":
			// Handle external_evaluation blocks (no expansion needed, just keep as-is)
			externalEvalBlocks, ok := item.ValueExpr.(*hclsyntax.TupleConsExpr)
			if !ok {
				diags.HclDiagnostics.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "Expected list for external_evaluation blocks",
					Detail:   fmt.Sprintf("Expected TupleConsExpr but got %T", item.ValueExpr),
				})
				continue
			}

			for _, externalEvalBlockExpr := range externalEvalBlocks.Exprs {
				newObj := &hclsyntax.ObjectConsExpr{
					Items: []hclsyntax.ObjectConsItem{
						{
							KeyExpr:   ast.NewKeyExpr("external_evaluation"),
							ValueExpr: externalEvalBlockExpr,
						},
					},
				}
				newObjs = append(newObjs, newObj)
			}

		default:
			// Skip unknown attributes - they don't need expansion
		}
	}

	return newObjs
}