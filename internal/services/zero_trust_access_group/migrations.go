// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessGroupResource)(nil)

// zeroTrustAccessGroupResourceSchemaV0 defines the v0 schema (v4 provider format)
var zeroTrustAccessGroupResourceSchemaV0 = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"account_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"zone_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required: true,
		},
	},
	Blocks: map[string]schema.Block{
		"include": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"email": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"email_domain": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"ip": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"everyone": schema.BoolAttribute{
						Optional: true,
					},
					"certificate": schema.BoolAttribute{
						Optional: true,
					},
					"any_valid_service_token": schema.BoolAttribute{
						Optional: true,
					},
				},
				Blocks: map[string]schema.Block{
					"azure": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"github": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Optional: true,
								},
								"teams": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"gsuite": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"email": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"okta": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		"exclude": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"email": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"email_domain": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"ip": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"everyone": schema.BoolAttribute{
						Optional: true,
					},
					"certificate": schema.BoolAttribute{
						Optional: true,
					},
					"any_valid_service_token": schema.BoolAttribute{
						Optional: true,
					},
				},
				Blocks: map[string]schema.Block{
					"azure": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"github": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Optional: true,
								},
								"teams": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"gsuite": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"email": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"okta": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		"require": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"email": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"email_domain": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"ip": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"everyone": schema.BoolAttribute{
						Optional: true,
					},
					"certificate": schema.BoolAttribute{
						Optional: true,
					},
					"any_valid_service_token": schema.BoolAttribute{
						Optional: true,
					},
				},
				Blocks: map[string]schema.Block{
					"azure": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"github": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Optional: true,
								},
								"teams": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"gsuite": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"email": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"okta": schema.ListNestedBlock{
						NestedObject: schema.NestedBlockObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"identity_provider_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
	},
}

func (r *ZeroTrustAccessGroupResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// State upgrade from v0 (v4 provider) to v1 (v5 provider)
		0: {
			PriorSchema:   &zeroTrustAccessGroupResourceSchemaV0,
			StateUpgrader: upgradeZeroTrustAccessGroupStateV0toV1,
		},
	}
}

func upgradeZeroTrustAccessGroupStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Parse the old state using the raw state data
	var oldState map[string]interface{}
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return
	}

	// Extract raw state attributes from JSON
	err := json.Unmarshal(req.RawState.JSON, &oldState)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to parse raw state",
			fmt.Sprintf("Could not parse raw state during migration: %s", err),
		)
		return
	}

	tflog.Info(ctx, "Starting zero trust access group state migration from v0 to v1")

	// Create new state structure
	var newState ZeroTrustAccessGroupModel

	// Migrate basic attributes
	if id, ok := oldState["id"].(string); ok && id != "" {
		newState.ID = types.StringValue(id)
	}
	if accountID, ok := oldState["account_id"].(string); ok && accountID != "" {
		newState.AccountID = types.StringValue(accountID)
	}
	if zoneID, ok := oldState["zone_id"].(string); ok && zoneID != "" {
		newState.ZoneID = types.StringValue(zoneID)
	}
	if name, ok := oldState["name"].(string); ok && name != "" {
		newState.Name = types.StringValue(name)
	}

	// Migrate include rules
	if includeData, ok := oldState["include"].([]interface{}); ok && len(includeData) > 0 {
		var includeRules []ZeroTrustAccessGroupIncludeModel
		for _, includeItem := range includeData {
			if includeMap, ok := includeItem.(map[string]interface{}); ok {
				migratedRules := migrateV4IncludeRuleBlock(includeMap)
				includeRules = append(includeRules, migratedRules...)
			}
		}
		if len(includeRules) > 0 {
			var includeRulesPtrs []*ZeroTrustAccessGroupIncludeModel
			for i := range includeRules {
				includeRulesPtrs = append(includeRulesPtrs, &includeRules[i])
			}
			newState.Include = &includeRulesPtrs
		}
	}

	// Migrate exclude rules
	if excludeData, ok := oldState["exclude"].([]interface{}); ok && len(excludeData) > 0 {
		var excludeRules []*ZeroTrustAccessGroupExcludeModel
		for _, excludeItem := range excludeData {
			if excludeMap, ok := excludeItem.(map[string]interface{}); ok {
				migratedRules := migrateV4ExcludeRuleBlock(excludeMap)
				excludeRules = append(excludeRules, migratedRules...)
			}
		}
		if len(excludeRules) > 0 {
			newState.Exclude = &excludeRules
		}
	}

	// Migrate require rules
	if requireData, ok := oldState["require"].([]interface{}); ok && len(requireData) > 0 {
		var requireRules []*ZeroTrustAccessGroupRequireModel
		for _, requireItem := range requireData {
			if requireMap, ok := requireItem.(map[string]interface{}); ok {
				migratedRules := migrateV4RequireRuleBlock(requireMap)
				requireRules = append(requireRules, migratedRules...)
			}
		}
		if len(requireRules) > 0 {
			newState.Require = &requireRules
		}
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Failed to set new state: %v", resp.Diagnostics.Errors()))
		return
	}

	tflog.Info(ctx, "Successfully migrated zero trust access group state from v0 to v1")
}

// migrateV4IncludeRuleBlock converts a v4 include rule block to multiple v5 include rule objects
// Processes attributes in a fixed order for deterministic results
func migrateV4IncludeRuleBlock(ruleMap map[string]interface{}) []ZeroTrustAccessGroupIncludeModel {
	var rules []ZeroTrustAccessGroupIncludeModel

	// Process in fixed order: email, email_domain, ip
	// This matches the config migration order
	
	// Handle email arrays -> multiple email objects
	if emails, ok := ruleMap["email"].([]interface{}); ok {
		for _, email := range emails {
			if emailStr, ok := email.(string); ok && emailStr != "" && emailStr != "<nil>" && emailStr != "null" {
				// Create completely clean object with ONLY email field initialized
				rule := ZeroTrustAccessGroupIncludeModel{}
				rule.Email = &ZeroTrustAccessGroupIncludeEmailModel{
					Email: types.StringValue(emailStr),
				}
				rules = append(rules, rule)
			}
		}
	}

	// Handle email_domain arrays -> multiple email_domain objects
	if emailDomains, ok := ruleMap["email_domain"].([]interface{}); ok {
		for _, domain := range emailDomains {
			if domain == nil {
				continue // Skip nil values explicitly
			}
			if domainStr, ok := domain.(string); ok && domainStr != "" && domainStr != "<nil>" && domainStr != "null" {
				// Create completely clean object with ONLY email_domain field initialized
				rule := ZeroTrustAccessGroupIncludeModel{}
				rule.EmailDomain = &ZeroTrustAccessGroupIncludeEmailDomainModel{
					Domain: types.StringValue(domainStr),
				}
				rules = append(rules, rule)
			}
		}
	}

	// Handle ip arrays -> multiple ip objects
	if ips, ok := ruleMap["ip"].([]interface{}); ok {
		for _, ip := range ips {
			if ipStr, ok := ip.(string); ok && ipStr != "" && ipStr != "<nil>" && ipStr != "null" {
				// Create completely clean object with ONLY ip field initialized
				rule := ZeroTrustAccessGroupIncludeModel{}
				rule.IP = &ZeroTrustAccessGroupIncludeIPModel{
					IP: types.StringValue(ipStr),
				}
				rules = append(rules, rule)
			}
		}
	}

	// Handle azure blocks -> multiple azure_ad objects
	// V4 structure: azure = [{ id = ["group1", "group2"], identity_provider_id = "provider" }]
	// V5 structure: azure_ad = { id = "group1", identity_provider_id = "provider" }
	if azureBlocks, ok := ruleMap["azure"].([]interface{}); ok {
		for _, azureBlock := range azureBlocks {
			if azureMap, ok := azureBlock.(map[string]interface{}); ok {
				identityProviderID := ""
				if providerID, ok := azureMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				// Handle ID arrays
				if ids, ok := azureMap["id"].([]interface{}); ok {
					for _, id := range ids {
						if idStr, ok := id.(string); ok {
							// Create completely clean object with ONLY azure_ad field initialized
							rule := ZeroTrustAccessGroupIncludeModel{}
							rule.AzureAD = &ZeroTrustAccessGroupIncludeAzureADModel{
								ID: types.StringValue(idStr),
								IdentityProviderID: types.StringValue(identityProviderID),
							}
							rules = append(rules, rule)
						}
					}
				}
			}
		}
	}

	// Handle github blocks -> multiple github_organization objects
	// V4 structure: github = [{ name = "org", teams = ["team1", "team2"], identity_provider_id = "provider" }]
	// V5 structure: github_organization = { name = "org", team = "team1", identity_provider_id = "provider" }
	if githubBlocks, ok := ruleMap["github"].([]interface{}); ok {
		for _, githubBlock := range githubBlocks {
			if githubMap, ok := githubBlock.(map[string]interface{}); ok {
				name := ""
				identityProviderID := ""
				if nameStr, ok := githubMap["name"].(string); ok {
					name = nameStr
				}
				if providerID, ok := githubMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				// Handle teams arrays
				if teams, ok := githubMap["teams"].([]interface{}); ok {
					for _, team := range teams {
						if teamStr, ok := team.(string); ok {
							// Create completely clean object with ONLY github_organization field initialized
							rule := ZeroTrustAccessGroupIncludeModel{}
							rule.GitHubOrganization = &ZeroTrustAccessGroupIncludeGitHubOrganizationModel{
								Name: types.StringValue(name),
								Team: types.StringValue(teamStr),
								IdentityProviderID: types.StringValue(identityProviderID),
							}
							rules = append(rules, rule)
						}
					}
				}
			}
		}
	}

	// Handle gsuite blocks -> expand email arrays
	// V4 structure: gsuite = [{ email = ["user1@example.com", "user2@example.com"], identity_provider_id = "provider" }]
	// V5 structure: gsuite = { email = "user1@example.com", identity_provider_id = "provider" }
	if gsuiteBlocks, ok := ruleMap["gsuite"].([]interface{}); ok {
		for _, gsuiteBlock := range gsuiteBlocks {
			if gsuiteMap, ok := gsuiteBlock.(map[string]interface{}); ok {
				identityProviderID := ""
				if providerID, ok := gsuiteMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				// Handle email arrays
				if emails, ok := gsuiteMap["email"].([]interface{}); ok {
					for _, email := range emails {
						if emailStr, ok := email.(string); ok {
							// Create completely clean object with ONLY gsuite field initialized
							rule := ZeroTrustAccessGroupIncludeModel{}
							rule.GSuite = &ZeroTrustAccessGroupIncludeGSuiteModel{
								Email: types.StringValue(emailStr),
								IdentityProviderID: types.StringValue(identityProviderID),
							}
							rules = append(rules, rule)
						}
					}
				}
			}
		}
	}

	// Handle okta blocks -> expand name arrays
	// V4 structure: okta = [{ name = ["group1", "group2"], identity_provider_id = "provider" }]
	// V5 structure: okta = { name = "group1", identity_provider_id = "provider" }
	if oktaBlocks, ok := ruleMap["okta"].([]interface{}); ok {
		for _, oktaBlock := range oktaBlocks {
			if oktaMap, ok := oktaBlock.(map[string]interface{}); ok {
				identityProviderID := ""
				if providerID, ok := oktaMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				// Handle name arrays
				if names, ok := oktaMap["name"].([]interface{}); ok {
					for _, name := range names {
						if nameStr, ok := name.(string); ok {
							// Create completely clean object with ONLY okta field initialized
							rule := ZeroTrustAccessGroupIncludeModel{}
							rule.Okta = &ZeroTrustAccessGroupIncludeOktaModel{
								Name: types.StringValue(nameStr),
								IdentityProviderID: types.StringValue(identityProviderID),
							}
							rules = append(rules, rule)
						}
					}
				}
			}
		}
	}

	// Handle boolean attributes -> empty objects
	// V4 structure: everyone = true, certificate = true, any_valid_service_token = true
	// V5 structure: everyone = {}, certificate = {}, any_valid_service_token = {}
	if everyone, ok := ruleMap["everyone"].(bool); ok && everyone {
		// Create completely clean object with ONLY everyone field initialized
		rule := ZeroTrustAccessGroupIncludeModel{}
		rule.Everyone = &ZeroTrustAccessGroupIncludeEveryoneModel{}
		rules = append(rules, rule)
	}
	
	if certificate, ok := ruleMap["certificate"].(bool); ok && certificate {
		// Create completely clean object with ONLY certificate field initialized
		rule := ZeroTrustAccessGroupIncludeModel{}
		rule.Certificate = &ZeroTrustAccessGroupIncludeCertificateModel{}
		rules = append(rules, rule)
	}
	
	if anyValidServiceToken, ok := ruleMap["any_valid_service_token"].(bool); ok && anyValidServiceToken {
		// Create completely clean object with ONLY any_valid_service_token field initialized
		rule := ZeroTrustAccessGroupIncludeModel{}
		rule.AnyValidServiceToken = &ZeroTrustAccessGroupIncludeAnyValidServiceTokenModel{}
		rules = append(rules, rule)
	}

	return rules
}

// migrateV4ExcludeRuleBlock converts a v4 exclude rule block to multiple v5 exclude rule objects
func migrateV4ExcludeRuleBlock(ruleMap map[string]interface{}) []*ZeroTrustAccessGroupExcludeModel {
	var rules []*ZeroTrustAccessGroupExcludeModel

	// Process attributes in same order as include rules for consistency
	
	// Handle email arrays -> multiple email objects
	if emails, ok := ruleMap["email"].([]interface{}); ok {
		for _, email := range emails {
			if emailStr, ok := email.(string); ok {
				// Create completely clean object with ONLY email field initialized
				rule := &ZeroTrustAccessGroupExcludeModel{}
				rule.Email = &ZeroTrustAccessGroupExcludeEmailModel{
					Email: types.StringValue(emailStr),
				}
				rules = append(rules, rule)
			}
		}
	}

	// Handle email_domain arrays -> multiple email_domain objects
	if emailDomains, ok := ruleMap["email_domain"].([]interface{}); ok {
		for _, domain := range emailDomains {
			if domainStr, ok := domain.(string); ok {
				// Create completely clean object with ONLY email_domain field initialized
				rule := &ZeroTrustAccessGroupExcludeModel{}
				rule.EmailDomain = &ZeroTrustAccessGroupExcludeEmailDomainModel{
					Domain: types.StringValue(domainStr),
				}
				rules = append(rules, rule)
			}
		}
	}

	// Handle ip arrays -> multiple ip objects
	if ips, ok := ruleMap["ip"].([]interface{}); ok {
		for _, ip := range ips {
			if ipStr, ok := ip.(string); ok {
				// Create completely clean object with ONLY ip field initialized
				rule := &ZeroTrustAccessGroupExcludeModel{}
				rule.IP = &ZeroTrustAccessGroupExcludeIPModel{
					IP: types.StringValue(ipStr),
				}
				rules = append(rules, rule)
			}
		}
	}

	// Handle azure blocks -> multiple azure_ad objects
	if azureBlocks, ok := ruleMap["azure"].([]interface{}); ok {
		for _, azureBlock := range azureBlocks {
			if azureMap, ok := azureBlock.(map[string]interface{}); ok {
				identityProviderID := ""
				if providerID, ok := azureMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				if ids, ok := azureMap["id"].([]interface{}); ok {
					for _, id := range ids {
						if idStr, ok := id.(string); ok {
							rules = append(rules, &ZeroTrustAccessGroupExcludeModel{
								AzureAD: &ZeroTrustAccessGroupExcludeAzureADModel{
									ID: types.StringValue(idStr),
									IdentityProviderID: types.StringValue(identityProviderID),
								},
							})
						}
					}
				}
			}
		}
	}

	// Handle github blocks -> multiple github_organization objects
	if githubBlocks, ok := ruleMap["github"].([]interface{}); ok {
		for _, githubBlock := range githubBlocks {
			if githubMap, ok := githubBlock.(map[string]interface{}); ok {
				name := ""
				identityProviderID := ""
				if nameStr, ok := githubMap["name"].(string); ok {
					name = nameStr
				}
				if providerID, ok := githubMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				if teams, ok := githubMap["teams"].([]interface{}); ok {
					for _, team := range teams {
						if teamStr, ok := team.(string); ok {
							rules = append(rules, &ZeroTrustAccessGroupExcludeModel{
								GitHubOrganization: &ZeroTrustAccessGroupExcludeGitHubOrganizationModel{
									Name: types.StringValue(name),
									Team: types.StringValue(teamStr),
									IdentityProviderID: types.StringValue(identityProviderID),
								},
							})
						}
					}
				}
			}
		}
	}

	// Handle gsuite blocks -> expand email arrays
	if gsuiteBlocks, ok := ruleMap["gsuite"].([]interface{}); ok {
		for _, gsuiteBlock := range gsuiteBlocks {
			if gsuiteMap, ok := gsuiteBlock.(map[string]interface{}); ok {
				identityProviderID := ""
				if providerID, ok := gsuiteMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				if emails, ok := gsuiteMap["email"].([]interface{}); ok {
					for _, email := range emails {
						if emailStr, ok := email.(string); ok {
							rules = append(rules, &ZeroTrustAccessGroupExcludeModel{
								GSuite: &ZeroTrustAccessGroupExcludeGSuiteModel{
									Email: types.StringValue(emailStr),
									IdentityProviderID: types.StringValue(identityProviderID),
								},
							})
						}
					}
				}
			}
		}
	}

	// Handle okta blocks -> expand name arrays
	if oktaBlocks, ok := ruleMap["okta"].([]interface{}); ok {
		for _, oktaBlock := range oktaBlocks {
			if oktaMap, ok := oktaBlock.(map[string]interface{}); ok {
				identityProviderID := ""
				if providerID, ok := oktaMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				if names, ok := oktaMap["name"].([]interface{}); ok {
					for _, name := range names {
						if nameStr, ok := name.(string); ok {
							rules = append(rules, &ZeroTrustAccessGroupExcludeModel{
								Okta: &ZeroTrustAccessGroupExcludeOktaModel{
									Name: types.StringValue(nameStr),
									IdentityProviderID: types.StringValue(identityProviderID),
								},
							})
						}
					}
				}
			}
		}
	}

	// Handle boolean attributes -> empty objects
	if everyone, ok := ruleMap["everyone"].(bool); ok && everyone {
		rules = append(rules, &ZeroTrustAccessGroupExcludeModel{
			Everyone: &ZeroTrustAccessGroupExcludeEveryoneModel{},
		})
	}
	
	if certificate, ok := ruleMap["certificate"].(bool); ok && certificate {
		rules = append(rules, &ZeroTrustAccessGroupExcludeModel{
			Certificate: &ZeroTrustAccessGroupExcludeCertificateModel{},
		})
	}
	
	if anyValidServiceToken, ok := ruleMap["any_valid_service_token"].(bool); ok && anyValidServiceToken {
		rules = append(rules, &ZeroTrustAccessGroupExcludeModel{
			AnyValidServiceToken: &ZeroTrustAccessGroupExcludeAnyValidServiceTokenModel{},
		})
	}

	return rules
}

// migrateV4RequireRuleBlock converts a v4 require rule block to multiple v5 require rule objects
func migrateV4RequireRuleBlock(ruleMap map[string]interface{}) []*ZeroTrustAccessGroupRequireModel {
	var rules []*ZeroTrustAccessGroupRequireModel

	// Process attributes in same order as include rules for consistency
	
	// Handle email arrays -> multiple email objects
	if emails, ok := ruleMap["email"].([]interface{}); ok {
		for _, email := range emails {
			if emailStr, ok := email.(string); ok {
				rules = append(rules, &ZeroTrustAccessGroupRequireModel{
					Email: &ZeroTrustAccessGroupRequireEmailModel{
						Email: types.StringValue(emailStr),
					},
				})
			}
		}
	}

	// Handle email_domain arrays -> multiple email_domain objects
	if emailDomains, ok := ruleMap["email_domain"].([]interface{}); ok {
		for _, domain := range emailDomains {
			if domainStr, ok := domain.(string); ok {
				rules = append(rules, &ZeroTrustAccessGroupRequireModel{
					EmailDomain: &ZeroTrustAccessGroupRequireEmailDomainModel{
						Domain: types.StringValue(domainStr),
					},
				})
			}
		}
	}

	// Handle ip arrays -> multiple ip objects
	if ips, ok := ruleMap["ip"].([]interface{}); ok {
		for _, ip := range ips {
			if ipStr, ok := ip.(string); ok {
				rules = append(rules, &ZeroTrustAccessGroupRequireModel{
					IP: &ZeroTrustAccessGroupRequireIPModel{
						IP: types.StringValue(ipStr),
					},
				})
			}
		}
	}

	// Handle azure blocks -> multiple azure_ad objects
	if azureBlocks, ok := ruleMap["azure"].([]interface{}); ok {
		for _, azureBlock := range azureBlocks {
			if azureMap, ok := azureBlock.(map[string]interface{}); ok {
				identityProviderID := ""
				if providerID, ok := azureMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				if ids, ok := azureMap["id"].([]interface{}); ok {
					for _, id := range ids {
						if idStr, ok := id.(string); ok {
							rules = append(rules, &ZeroTrustAccessGroupRequireModel{
								AzureAD: &ZeroTrustAccessGroupRequireAzureADModel{
									ID: types.StringValue(idStr),
									IdentityProviderID: types.StringValue(identityProviderID),
								},
							})
						}
					}
				}
			}
		}
	}

	// Handle github blocks -> multiple github_organization objects
	if githubBlocks, ok := ruleMap["github"].([]interface{}); ok {
		for _, githubBlock := range githubBlocks {
			if githubMap, ok := githubBlock.(map[string]interface{}); ok {
				name := ""
				identityProviderID := ""
				if nameStr, ok := githubMap["name"].(string); ok {
					name = nameStr
				}
				if providerID, ok := githubMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				if teams, ok := githubMap["teams"].([]interface{}); ok {
					for _, team := range teams {
						if teamStr, ok := team.(string); ok {
							rules = append(rules, &ZeroTrustAccessGroupRequireModel{
								GitHubOrganization: &ZeroTrustAccessGroupRequireGitHubOrganizationModel{
									Name: types.StringValue(name),
									Team: types.StringValue(teamStr),
									IdentityProviderID: types.StringValue(identityProviderID),
								},
							})
						}
					}
				}
			}
		}
	}

	// Handle gsuite blocks -> expand email arrays
	if gsuiteBlocks, ok := ruleMap["gsuite"].([]interface{}); ok {
		for _, gsuiteBlock := range gsuiteBlocks {
			if gsuiteMap, ok := gsuiteBlock.(map[string]interface{}); ok {
				identityProviderID := ""
				if providerID, ok := gsuiteMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				if emails, ok := gsuiteMap["email"].([]interface{}); ok {
					for _, email := range emails {
						if emailStr, ok := email.(string); ok {
							rules = append(rules, &ZeroTrustAccessGroupRequireModel{
								GSuite: &ZeroTrustAccessGroupRequireGSuiteModel{
									Email: types.StringValue(emailStr),
									IdentityProviderID: types.StringValue(identityProviderID),
								},
							})
						}
					}
				}
			}
		}
	}

	// Handle okta blocks -> expand name arrays
	if oktaBlocks, ok := ruleMap["okta"].([]interface{}); ok {
		for _, oktaBlock := range oktaBlocks {
			if oktaMap, ok := oktaBlock.(map[string]interface{}); ok {
				identityProviderID := ""
				if providerID, ok := oktaMap["identity_provider_id"].(string); ok {
					identityProviderID = providerID
				}
				
				if names, ok := oktaMap["name"].([]interface{}); ok {
					for _, name := range names {
						if nameStr, ok := name.(string); ok {
							rules = append(rules, &ZeroTrustAccessGroupRequireModel{
								Okta: &ZeroTrustAccessGroupRequireOktaModel{
									Name: types.StringValue(nameStr),
									IdentityProviderID: types.StringValue(identityProviderID),
								},
							})
						}
					}
				}
			}
		}
	}

	// Handle boolean attributes -> empty objects
	if everyone, ok := ruleMap["everyone"].(bool); ok && everyone {
		rules = append(rules, &ZeroTrustAccessGroupRequireModel{
			Everyone: &ZeroTrustAccessGroupRequireEveryoneModel{},
		})
	}
	
	if certificate, ok := ruleMap["certificate"].(bool); ok && certificate {
		rules = append(rules, &ZeroTrustAccessGroupRequireModel{
			Certificate: &ZeroTrustAccessGroupRequireCertificateModel{},
		})
	}
	
	if anyValidServiceToken, ok := ruleMap["any_valid_service_token"].(bool); ok && anyValidServiceToken {
		rules = append(rules, &ZeroTrustAccessGroupRequireModel{
			AnyValidServiceToken: &ZeroTrustAccessGroupRequireAnyValidServiceTokenModel{},
		})
	}

	return rules
}
