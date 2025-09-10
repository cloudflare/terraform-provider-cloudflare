// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_token

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*AccountTokenResource)(nil)

func (r *AccountTokenResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var rawState map[string]interface{}
				if err := json.Unmarshal(req.RawState.JSON, &rawState); err != nil {
					resp.Diagnostics.AddError(
						"Unable to Unmarshal Prior State",
						err.Error(),
					)
					return
				}
				
				// Convert policies from JSON string array to proper format
				if policies, exists := rawState["policies"]; exists {
					if policiesStr, ok := policies.(string); ok {
						var policiesArray []interface{}
						if err := json.Unmarshal([]byte(policiesStr), &policiesArray); err == nil {
							rawState["policies"] = policiesArray
						}
					}
				}

				// Ensure account_id is preserved during migration
				if _, exists := rawState["account_id"]; !exists {
					resp.Diagnostics.AddError(
						"Missing account_id in prior state",
						"The account_id field is required but missing from the previous state",
					)
					return
				}
				
				// Build the upgraded model manually to handle framework types properly
				upgradedModel := AccountTokenModel{}
				
				// Convert string fields to types.String
				if id, ok := rawState["id"].(string); ok {
					upgradedModel.ID = types.StringValue(id)
				}
				if accountID, ok := rawState["account_id"].(string); ok {
					upgradedModel.AccountID = types.StringValue(accountID)
				}
				if name, ok := rawState["name"].(string); ok {
					upgradedModel.Name = types.StringValue(name)
				}
				if status, ok := rawState["status"].(string); ok {
					upgradedModel.Status = types.StringValue(status)
				}
				if value, ok := rawState["value"].(string); ok {
					upgradedModel.Value = types.StringValue(value)
				}
				
				// Handle the policies field as Dynamic - let the framework handle the conversion
				if policies, exists := rawState["policies"]; exists {
					// Marshal and use the framework's JSON unmarshaling for Dynamic types
					policiesJSON, _ := json.Marshal(policies)
					var tempModel struct {
						Policies types.Dynamic `json:"policies"`
					}
					if err := json.Unmarshal(policiesJSON, &tempModel); err == nil {
						upgradedModel.Policies = tempModel.Policies
					}
				}
				
				resp.Diagnostics.Append(resp.State.Set(ctx, upgradedModel)...)
			},
		},
	}
}
