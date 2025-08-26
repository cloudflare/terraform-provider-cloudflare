package account_member

import (
	"context"
	"encoding/json"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m AccountMemberModel) marshalCustom() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccountMemberModel) marshalCustomForUpdate(state AccountMemberModel) (data []byte, err error) {
	data, err = apijson.MarshalForUpdate(m, state)
	if err != nil {
		return
	}

	if m.Roles == nil || len(*m.Roles) == 0 {
		return
	}

	var payload map[string]interface{}
	if err = json.Unmarshal(data, &payload); err != nil {
		return
	}

	// Transform roles: ["role_id"] -> [{"id": "role_id"}]
	roleObjects := make([]map[string]interface{}, 0, len(*m.Roles))
	for _, role := range *m.Roles {
		if !role.IsNull() && !role.IsUnknown() {
			roleObjects = append(roleObjects, map[string]interface{}{
				"id": role.ValueString(),
			})
		}
	}

	if len(roleObjects) > 0 {
		payload["roles"] = roleObjects
	}

	// Ensure account_id is included for the API even though it's a path param
	if !m.AccountID.IsNull() && !m.AccountID.IsUnknown() {
		payload["account_id"] = m.AccountID.ValueString()
	}

	return json.Marshal(payload)
}

func unmarshalCustom(data []byte, configuredModel *AccountMemberModel) (*AccountMemberModel, error) {
	var env AccountMemberResultEnvelope
	if err := apijson.UnmarshalComputed(data, &env); err != nil {
		return nil, err
	}

	result := &env.Result

	// Preserve required fields from configured model to avoid replacement
	result.AccountID = configuredModel.AccountID
	result.Email = configuredModel.Email

	// Only preserve status if it was explicitly configured
	if !configuredModel.Status.IsNull() && !configuredModel.Status.IsUnknown() {
		result.Status = configuredModel.Status
	}

	// Determine comparison strategy based on what's configured in Terraform
	// as user can use roles or policies to configure the account member
	configUsesRoles := configuredModel.Roles != nil && len(*configuredModel.Roles) > 0
	configUsesPolicies := !configuredModel.Policies.IsNull() && !configuredModel.Policies.IsUnknown()

	if configUsesRoles && !configUsesPolicies {
		// User configured roles - preserve computed policies from API to prevent diffs
		// Don't modify result.Policies - let it keep the computed values from API response

		// Extract role IDs from GET response (roles come as objects with id field)
		var fullResponse struct {
			Result struct {
				Roles []struct {
					ID string `json:"id"`
				} `json:"roles"`
			} `json:"result"`
		}

		if err := json.Unmarshal(data, &fullResponse); err == nil && len(fullResponse.Result.Roles) > 0 {
			roleIDs := make([]types.String, 0, len(fullResponse.Result.Roles))
			for _, role := range fullResponse.Result.Roles {
				roleIDs = append(roleIDs, types.StringValue(role.ID))
			}
			result.Roles = &roleIDs
		} else {
			result.Roles = configuredModel.Roles
		}
	} else if configUsesPolicies && !configUsesRoles {
		// User configured policies - ignore roles, set roles to nil to prevent diffs
		result.Roles = nil

		// Extract policies from the API response and ensure they match the configured structure
		var fullResponse struct {
			Result struct {
				Policies []AccountMemberPoliciesModel `json:"policies"`
			} `json:"result"`
		}

		if err := json.Unmarshal(data, &fullResponse); err == nil && len(fullResponse.Result.Policies) > 0 {
			policies := append([]AccountMemberPoliciesModel(nil), fullResponse.Result.Policies...)
			policiesList, diags := customfield.NewObjectList(context.Background(), policies)
			if !diags.HasError() {
				result.Policies = policiesList
			} else {
				result.Policies = configuredModel.Policies
			}
		} else {
			result.Policies = configuredModel.Policies
		}
	} else {
		// Fallback: preserve configured values
		result.Roles = configuredModel.Roles
		result.Policies = configuredModel.Policies
	}

	return result, nil
}
