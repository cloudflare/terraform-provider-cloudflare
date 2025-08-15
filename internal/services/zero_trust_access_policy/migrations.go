// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessPolicyResource)(nil)

func (r *ZeroTrustAccessPolicyResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// Version 0 to 1: Transform boolean attributes in include/exclude/require
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"decision": schema.StringAttribute{
						Required: true,
					},
					"name": schema.StringAttribute{
						Required: true,
					},
					"approval_required": schema.BoolAttribute{
						Optional: true,
					},
					"isolation_required": schema.BoolAttribute{
						Optional: true,
					},
					"purpose_justification_prompt": schema.StringAttribute{
						Optional: true,
					},
					"purpose_justification_required": schema.BoolAttribute{
						Optional: true,
					},
					"session_duration": schema.StringAttribute{
						Optional: true,
						Computed: true,
					},
					"precedence": schema.Int64Attribute{
						Optional: true,
					},
					"approval_groups": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"approvals_needed": schema.Float64Attribute{
									Required: true,
								},
								"email_addresses": schema.ListAttribute{
									Optional:    true,
									ElementType: types.StringType,
								},
								"email_list_uuid": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				// Get the raw state
				var priorStateData struct {
					ID                           types.String `tfsdk:"id"`
					AccountID                    types.String `tfsdk:"account_id"`
					Decision                     types.String `tfsdk:"decision"`
					Name                         types.String `tfsdk:"name"`
					ApprovalRequired             types.Bool   `tfsdk:"approval_required"`
					IsolationRequired            types.Bool   `tfsdk:"isolation_required"`
					PurposeJustificationPrompt   types.String `tfsdk:"purpose_justification_prompt"`
					PurposeJustificationRequired types.Bool   `tfsdk:"purpose_justification_required"`
					SessionDuration              types.String `tfsdk:"session_duration"`
					Precedence                   types.Int64  `tfsdk:"precedence"`
					ApprovalGroups               types.List   `tfsdk:"approval_groups"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Set session_duration to default if null
				sessionDuration := priorStateData.SessionDuration
				if sessionDuration.IsNull() {
					fmt.Printf("Setting session duration to default!!")
					sessionDuration = types.StringValue("24h")
				} else {
					fmt.Printf("Current sessionDuration is %s, leaving as-is", sessionDuration.String())
				}

				// Create new state with transformed values
				upgradedStateData := ZeroTrustAccessPolicyModel{
					ID:                           priorStateData.ID,
					AccountID:                    priorStateData.AccountID,
					Decision:                     priorStateData.Decision,
					Name:                         priorStateData.Name,
					ApprovalRequired:             priorStateData.ApprovalRequired,
					IsolationRequired:            priorStateData.IsolationRequired,
					PurposeJustificationPrompt:   priorStateData.PurposeJustificationPrompt,
					PurposeJustificationRequired: priorStateData.PurposeJustificationRequired,
					SessionDuration:              sessionDuration,
					//Include:                      transformedInclude,
					//Exclude:                      transformedExclude,
					//Require:                      transformedRequire,
				}

				// Transform include/exclude/require to handle boolean attributes
				rawStateJSON := req.RawState.JSON
				var rawState map[string]interface{}
				err := json.Unmarshal(rawStateJSON, &rawState)
				if err != nil {
					resp.Diagnostics.AddError("failed to unmarshal state!!!", err.Error())
				}
				if includeRaw, exists := rawState["include"]; exists && includeRaw != nil {
					switch v := includeRaw.(type) {
					case []map[string]interface{}:
						upgradedStateData.Include = transformIncludeConditionList(ctx, v)
					}
				}

				//transformedExclude := transformConditionList(ctx, priorStateData.Exclude)
				//transformedRequire := transformConditionList(ctx, priorStateData.Require)

				// Dump the entire new state as JSON for debugging
				newStateJSON, err := json.MarshalIndent(upgradedStateData, "", "  ")
				if err != nil {
					fmt.Printf("DEBUG: Failed to marshal new state to JSON: %v\n", err)
				} else {
					fmt.Printf("DEBUG: New state JSON for access policy:\n%s\n", string(newStateJSON))
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, upgradedStateData)...)
			},
		},
	}
}

// transformConditionList transforms boolean attributes in include/exclude/require lists
// Converts "everyone": true to "everyone": {} and removes "everyone": false
func transformIncludeConditionList(ctx context.Context, conditionList []map[string]interface{}) *[]*ZeroTrustAccessPolicyIncludeModel {
	if conditionList == nil {
		return nil
	}

	var result []*ZeroTrustAccessPolicyIncludeModel

	// For now, just copy boolean attributes that we understand
	// The rest will be handled by normal state migration
	for _, attrs := range conditionList {
		includeModel := &ZeroTrustAccessPolicyIncludeModel{}

		// Handle "everyone" boolean -> empty object transformation
		if everyoneVal, ok := attrs["everyone"]; ok {
			if boolVal, ok := everyoneVal.(types.Bool); ok && !boolVal.IsNull() {
				if boolVal.ValueBool() {
					// Set everyone as an empty object (it has no fields)
					includeModel.Everyone = &ZeroTrustAccessPolicyIncludeEveryoneModel{}
				}
				// If false, we don't set it (effectively removing it)
			}
		}

		// Handle "certificate" boolean -> empty object transformation
		if certVal, ok := attrs["certificate"]; ok {
			if boolVal, ok := certVal.(types.Bool); ok && !boolVal.IsNull() {
				if boolVal.ValueBool() {
					includeModel.Certificate = &ZeroTrustAccessPolicyIncludeCertificateModel{}
				}
			}
		}

		// Handle "any_valid_service_token" boolean -> empty object transformation
		if tokenVal, ok := attrs["any_valid_service_token"]; ok {
			if boolVal, ok := tokenVal.(types.Bool); ok && !boolVal.IsNull() {
				if boolVal.ValueBool() {
					includeModel.AnyValidServiceToken = &ZeroTrustAccessPolicyIncludeAnyValidServiceTokenModel{}
				}
			}
		}

		// TODO: Copy other non-boolean attributes as-is
		// For now we're only handling the boolean transformations

		result = append(result, includeModel)
	}

	return &result
}
