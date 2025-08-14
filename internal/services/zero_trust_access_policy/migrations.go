// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

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

// zeroTrustAccessPolicyResourceSchemaV0 defines the v0 schema (v4 provider format)
var zeroTrustAccessPolicyResourceSchemaV0 = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"account_id": schema.StringAttribute{
			Required: true,
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
		"decision": schema.StringAttribute{
			Required: true,
		},
		"approval_required": schema.BoolAttribute{
			Optional: true,
		},
		"isolation_required": schema.BoolAttribute{
			Optional: true,
		},
		"purpose_justification_required": schema.BoolAttribute{
			Optional: true,
		},
		"purpose_justification_prompt": schema.StringAttribute{
			Optional: true,
		},
		"session_duration": schema.StringAttribute{
			Optional: true,
		},
	},
	Blocks: map[string]schema.Block{
		"approval_group": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"approvals_needed": schema.Int64Attribute{
						Required: true,
					},
					"email_addresses": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"email_list_uuid": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
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
					"geo": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"group": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"service_token": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"email_list": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"ip_list": schema.ListAttribute{
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
					"common_name": schema.StringAttribute{
						Optional: true,
					},
					"auth_method": schema.StringAttribute{
						Optional: true,
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
					"geo": schema.ListAttribute{
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
					"common_name": schema.StringAttribute{
						Optional: true,
					},
					"auth_method": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
		"require": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"certificate": schema.BoolAttribute{
						Optional: true,
					},
				},
			},
		},
	},
}

func (r *ZeroTrustAccessPolicyResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &zeroTrustAccessPolicyResourceSchemaV0,
			StateUpgrader: upgradeZeroTrustAccessPolicyStateV0toV1,
		},
	}
}

// upgradeZeroTrustAccessPolicyStateV0toV1 migrates from v4 provider state format to v5
func upgradeZeroTrustAccessPolicyStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Error(ctx, "STATE UPGRADER CALLED - upgrading zero_trust_access_policy state from v0 to v1")

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

	// Create new state structure - initialize with zero values to ensure old attributes are cleared
	var newState ZeroTrustAccessPolicyModel

	// Migrate basic attributes
	if accountID, ok := oldState["account_id"].(string); ok {
		newState.AccountID = types.StringValue(accountID)
	}
	if name, ok := oldState["name"].(string); ok {
		newState.Name = types.StringValue(name)
	}
	if decision, ok := oldState["decision"].(string); ok {
		newState.Decision = types.StringValue(decision)
	}
	if id, ok := oldState["id"].(string); ok {
		newState.ID = types.StringValue(id)
	}

	// Migrate approval_group to approval_groups
	if approvalGroupData, ok := oldState["approval_group"].([]interface{}); ok && len(approvalGroupData) > 0 {
		var approvalGroups []ZeroTrustAccessPolicyApprovalGroupsModel
		for _, groupItem := range approvalGroupData {
			if groupMap, ok := groupItem.(map[string]interface{}); ok {
				var group ZeroTrustAccessPolicyApprovalGroupsModel

				if approvalsNeeded, ok := groupMap["approvals_needed"].(float64); ok {
					group.ApprovalsNeeded = types.Float64Value(approvalsNeeded)
				} else if approvalsNeededInt, ok := groupMap["approvals_needed"].(int64); ok {
					group.ApprovalsNeeded = types.Float64Value(float64(approvalsNeededInt))
				}

				if emailAddresses, ok := groupMap["email_addresses"].([]interface{}); ok {
					var emails []types.String
					for _, email := range emailAddresses {
						if emailStr, ok := email.(string); ok {
							emails = append(emails, types.StringValue(emailStr))
						}
					}
					group.EmailAddresses = &emails
				}

				if emailListUuid, ok := groupMap["email_list_uuid"].(string); ok {
					group.EmailListUUID = types.StringPointerValue(&emailListUuid)
				}

				approvalGroups = append(approvalGroups, group)
			}
		}
		var approvalGroupsPtrs []*ZeroTrustAccessPolicyApprovalGroupsModel
		for i := range approvalGroups {
			approvalGroupsPtrs = append(approvalGroupsPtrs, &approvalGroups[i])
		}
		newState.ApprovalGroups = &approvalGroupsPtrs
	}

	// Migrate boolean flags
	if approvalRequired, ok := oldState["approval_required"].(bool); ok {
		newState.ApprovalRequired = types.BoolValue(approvalRequired)
	}
	if isolationRequired, ok := oldState["isolation_required"].(bool); ok {
		newState.IsolationRequired = types.BoolValue(isolationRequired)
	}
	if purposeJustificationRequired, ok := oldState["purpose_justification_required"].(bool); ok {
		newState.PurposeJustificationRequired = types.BoolValue(purposeJustificationRequired)
	}

	// Migrate string attributes
	if purposeJustificationPrompt, ok := oldState["purpose_justification_prompt"].(string); ok {
		newState.PurposeJustificationPrompt = types.StringPointerValue(&purposeJustificationPrompt)
	}
	if sessionDuration, ok := oldState["session_duration"].(string); ok && sessionDuration != "" {
		newState.SessionDuration = types.StringPointerValue(&sessionDuration)
	}

	// Migrate include rules
	if includeData, ok := oldState["include"].([]interface{}); ok {
		var includeRules []ZeroTrustAccessPolicyIncludeModel
		for _, conditionItem := range includeData {
			if conditionMap, ok := conditionItem.(map[string]interface{}); ok {
				rules := migrateConditionV4ToV5Include(conditionMap)
				includeRules = append(includeRules, rules...)
			}
		}
		if len(includeRules) > 0 {
			var includeRulesPtrs []*ZeroTrustAccessPolicyIncludeModel
			for i := range includeRules {
				includeRulesPtrs = append(includeRulesPtrs, &includeRules[i])
			}
			newState.Include = &includeRulesPtrs
		}
	}

	// Migrate exclude rules
	if excludeData, ok := oldState["exclude"].([]interface{}); ok {
		var excludeRules []ZeroTrustAccessPolicyExcludeModel
		for _, conditionItem := range excludeData {
			if conditionMap, ok := conditionItem.(map[string]interface{}); ok {
				rules := migrateConditionV4ToV5Exclude(conditionMap)
				excludeRules = append(excludeRules, rules...)
			}
		}
		if len(excludeRules) > 0 {
			var excludeRulesPtrs []*ZeroTrustAccessPolicyExcludeModel
			for i := range excludeRules {
				excludeRulesPtrs = append(excludeRulesPtrs, &excludeRules[i])
			}
			newState.Exclude = &excludeRulesPtrs
		}
	}

	// Migrate require rules
	if requireData, ok := oldState["require"].([]interface{}); ok {
		var requireRules []ZeroTrustAccessPolicyRequireModel
		for _, conditionItem := range requireData {
			if conditionMap, ok := conditionItem.(map[string]interface{}); ok {
				rules := migrateConditionV4ToV5Require(conditionMap)
				requireRules = append(requireRules, rules...)
			}
		}
		if len(requireRules) > 0 {
			var requireRulesPtrs []*ZeroTrustAccessPolicyRequireModel
			for i := range requireRules {
				requireRulesPtrs = append(requireRulesPtrs, &requireRules[i])
			}
			newState.Require = &requireRulesPtrs
		}
	}

	// CRITICAL: Explicitly remove the old approval_group from raw state to prevent conflicts
	if oldState["approval_group"] != nil {
		delete(oldState, "approval_group") // Remove from the map we parsed
	}

	// Set the upgraded state - this should completely replace the old state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Failed to set new state: %v", resp.Diagnostics.Errors()))
		return
	}
}

// migrateConditionV4ToV5Include converts a v4 condition object to v5 include rules
func migrateConditionV4ToV5Include(condition map[string]interface{}) []ZeroTrustAccessPolicyIncludeModel {
	var rules []ZeroTrustAccessPolicyIncludeModel

	// Don't create any rules if this condition object is empty or has no meaningful values
	hasValidAttributes := false
	for key, value := range condition {
		switch key {
		case "everyone", "certificate", "any_valid_service_token":
			if boolVal, ok := value.(bool); ok && boolVal {
				hasValidAttributes = true
			}
		case "email", "email_domain", "ip", "geo", "group", "service_token", "email_list", "ip_list":
			if listVal, ok := value.([]interface{}); ok && len(listVal) > 0 {
				// Check if list has any non-empty, non-null values
				for _, item := range listVal {
					if strVal, ok := item.(string); ok && strVal != "" && strVal != "<nil>" && strVal != "null" {
						hasValidAttributes = true
						break
					}
				}
			}
		case "common_name", "auth_method":
			if strVal, ok := value.(string); ok && strVal != "" && strVal != "<nil>" && strVal != "null" {
				hasValidAttributes = true
			}
		}
		if hasValidAttributes {
			break
		}
	}

	if !hasValidAttributes {
		return rules // Return empty rules if no valid attributes found
	}

	// Boolean attributes that become empty nested objects - only create rule if explicitly true
	if everyone, ok := condition["everyone"].(bool); ok && everyone {
		rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
			Everyone: &ZeroTrustAccessPolicyIncludeEveryoneModel{},
		})
	}
	if certificate, ok := condition["certificate"].(bool); ok && certificate {
		rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
			Certificate: &ZeroTrustAccessPolicyIncludeCertificateModel{},
		})
	}
	if anyValidServiceToken, ok := condition["any_valid_service_token"].(bool); ok && anyValidServiceToken {
		rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
			AnyValidServiceToken: &ZeroTrustAccessPolicyIncludeAnyValidServiceTokenModel{},
		})
	}

	// List attributes that split into separate rules
	if emails, ok := condition["email"].([]interface{}); ok && len(emails) > 0 {
		for _, email := range emails {
			if emailStr, ok := email.(string); ok && emailStr != "" && emailStr != "<nil>" && emailStr != "null" {
				rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
					Email: &ZeroTrustAccessPolicyIncludeEmailModel{
						Email: types.StringValue(emailStr),
					},
				})
			}
		}
	}

	if emailDomains, ok := condition["email_domain"].([]interface{}); ok && len(emailDomains) > 0 {
		for _, domain := range emailDomains {
			if domain == nil {
				continue // Skip nil values explicitly
			}
			if domainStr, ok := domain.(string); ok && domainStr != "" && domainStr != "<nil>" && domainStr != "null" {
				rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
					EmailDomain: &ZeroTrustAccessPolicyIncludeEmailDomainModel{
						Domain: types.StringValue(domainStr),
					},
				})
			}
		}
	}

	if ips, ok := condition["ip"].([]interface{}); ok && len(ips) > 0 {
		for _, ip := range ips {
			if ipStr, ok := ip.(string); ok && ipStr != "" && ipStr != "<nil>" && ipStr != "null" {
				rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
					IP: &ZeroTrustAccessPolicyIncludeIPModel{
						IP: types.StringValue(ipStr),
					},
				})
			}
		}
	}

	if geos, ok := condition["geo"].([]interface{}); ok && len(geos) > 0 {
		for _, geo := range geos {
			if geoStr, ok := geo.(string); ok && geoStr != "" && geoStr != "<nil>" && geoStr != "null" {
				rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
					Geo: &ZeroTrustAccessPolicyIncludeGeoModel{
						CountryCode: types.StringValue(geoStr),
					},
				})
			}
		}
	}

	if groups, ok := condition["group"].([]interface{}); ok && len(groups) > 0 {
		for _, group := range groups {
			if groupStr, ok := group.(string); ok && groupStr != "" {
				rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
					Group: &ZeroTrustAccessPolicyIncludeGroupModel{
						ID: types.StringValue(groupStr),
					},
				})
			}
		}
	}

	if serviceTokens, ok := condition["service_token"].([]interface{}); ok && len(serviceTokens) > 0 {
		for _, token := range serviceTokens {
			if tokenStr, ok := token.(string); ok && tokenStr != "" {
				rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
					ServiceToken: &ZeroTrustAccessPolicyIncludeServiceTokenModel{
						TokenID: types.StringValue(tokenStr),
					},
				})
			}
		}
	}

	if emailLists, ok := condition["email_list"].([]interface{}); ok && len(emailLists) > 0 {
		for _, emailList := range emailLists {
			if emailListStr, ok := emailList.(string); ok && emailListStr != "" {
				rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
					EmailList: &ZeroTrustAccessPolicyIncludeEmailListModel{
						ID: types.StringValue(emailListStr),
					},
				})
			}
		}
	}

	if ipLists, ok := condition["ip_list"].([]interface{}); ok && len(ipLists) > 0 {
		for _, ipList := range ipLists {
			if ipListStr, ok := ipList.(string); ok && ipListStr != "" {
				rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
					IPList: &ZeroTrustAccessPolicyIncludeIPListModel{
						ID: types.StringValue(ipListStr),
					},
				})
			}
		}
	}

	// String attributes
	if commonName, ok := condition["common_name"].(string); ok && commonName != "" {
		rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
			CommonName: &ZeroTrustAccessPolicyIncludeCommonNameModel{
				CommonName: types.StringValue(commonName),
			},
		})
	}

	if authMethod, ok := condition["auth_method"].(string); ok && authMethod != "" {
		rules = append(rules, ZeroTrustAccessPolicyIncludeModel{
			AuthMethod: &ZeroTrustAccessPolicyIncludeAuthMethodModel{
				AuthMethod: types.StringValue(authMethod),
			},
		})
	}

	return rules
}

// migrateConditionV4ToV5Exclude converts a v4 condition object to v5 exclude rules
// In v4, each condition block had arrays of values per attribute type. In v5, we need to create separate rules
// for each value, but preserve the fact that they were separate conditions.
func migrateConditionV4ToV5Exclude(condition map[string]interface{}) []ZeroTrustAccessPolicyExcludeModel {
	var rules []ZeroTrustAccessPolicyExcludeModel

	// Boolean attributes - each creates one rule
	if everyone, ok := condition["everyone"].(bool); ok && everyone {
		rules = append(rules, ZeroTrustAccessPolicyExcludeModel{
			Everyone: &ZeroTrustAccessPolicyExcludeEveryoneModel{},
		})
	}
	if certificate, ok := condition["certificate"].(bool); ok && certificate {
		rules = append(rules, ZeroTrustAccessPolicyExcludeModel{
			Certificate: &ZeroTrustAccessPolicyExcludeCertificateModel{},
		})
	}
	if anyValidServiceToken, ok := condition["any_valid_service_token"].(bool); ok && anyValidServiceToken {
		rules = append(rules, ZeroTrustAccessPolicyExcludeModel{
			AnyValidServiceToken: &ZeroTrustAccessPolicyExcludeAnyValidServiceTokenModel{},
		})
	}

	// List attributes - each item creates one rule
	if emails, ok := condition["email"].([]interface{}); ok && len(emails) > 0 {
		for _, email := range emails {
			if emailStr, ok := email.(string); ok && emailStr != "" && emailStr != "<nil>" && emailStr != "null" {
				rules = append(rules, ZeroTrustAccessPolicyExcludeModel{
					Email: &ZeroTrustAccessPolicyExcludeEmailModel{
						Email: types.StringValue(emailStr),
					},
				})
			}
		}
	}

	if geos, ok := condition["geo"].([]interface{}); ok && len(geos) > 0 {
		for _, geo := range geos {
			if geo == nil {
				continue // Skip nil values
			}
			if geoStr, ok := geo.(string); ok && geoStr != "" && geoStr != "<nil>" && geoStr != "null" {
				rules = append(rules, ZeroTrustAccessPolicyExcludeModel{
					Geo: &ZeroTrustAccessPolicyExcludeGeoModel{
						CountryCode: types.StringValue(geoStr),
					},
				})
			}
		}
	}

	// String attributes - each creates one rule
	if commonName, ok := condition["common_name"].(string); ok && commonName != "" && commonName != "<nil>" && commonName != "null" {
		rules = append(rules, ZeroTrustAccessPolicyExcludeModel{
			CommonName: &ZeroTrustAccessPolicyExcludeCommonNameModel{
				CommonName: types.StringValue(commonName),
			},
		})
	}

	if authMethod, ok := condition["auth_method"].(string); ok && authMethod != "" && authMethod != "<nil>" && authMethod != "null" {
		rules = append(rules, ZeroTrustAccessPolicyExcludeModel{
			AuthMethod: &ZeroTrustAccessPolicyExcludeAuthMethodModel{
				AuthMethod: types.StringValue(authMethod),
			},
		})
	}

	return rules
}

// migrateConditionV4ToV5Require converts a v4 condition object to v5 require rules
func migrateConditionV4ToV5Require(condition map[string]interface{}) []ZeroTrustAccessPolicyRequireModel {
	var rules []ZeroTrustAccessPolicyRequireModel

	// Don't create any rules if this condition object is empty or has no meaningful values
	hasValidAttributes := false
	for key, value := range condition {
		switch key {
		case "certificate":
			if boolVal, ok := value.(bool); ok && boolVal {
				hasValidAttributes = true
			}
		}
		if hasValidAttributes {
			break
		}
	}

	if !hasValidAttributes {
		return rules // Return empty rules if no valid attributes found
	}

	// Boolean attributes
	if certificate, ok := condition["certificate"].(bool); ok && certificate {
		rules = append(rules, ZeroTrustAccessPolicyRequireModel{
			Certificate: &ZeroTrustAccessPolicyRequireCertificateModel{},
		})
	}

	return rules
}
