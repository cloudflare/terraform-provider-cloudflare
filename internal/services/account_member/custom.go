package account_member

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConfiguredPermissionType int

const (
	Unknown ConfiguredPermissionType = iota
	Policies
	Roles
)

// Given a config value, determine which permission type is configured. This
// should only be called with the config value, never the plan or state value.
func checkConfiguredPermissionType(configuredModel *AccountMemberModel) ConfiguredPermissionType {
	permissionType := Unknown

	if !configuredModel.Roles.IsNull() && len(configuredModel.Roles.Elements()) > 0 {
		permissionType = Roles
	}
	if !configuredModel.Policies.IsNull() && len(configuredModel.Policies.Elements()) > 0 {
		if permissionType == Roles {
			// if both are found, return Unknown
			return Unknown
		}
		return Policies
	}
	return permissionType
}

func (m AccountMemberModel) marshalCustom() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccountMemberModel) marshalCustomForUpdate(state AccountMemberModel, configuredPermissionType ConfiguredPermissionType) (data []byte, err error) {
	data, err = apijson.MarshalForUpdate(m, state)
	if err != nil {
		return
	}

	var payload map[string]interface{}
	if err = json.Unmarshal(data, &payload); err != nil {
		return
	}

	if configuredPermissionType == Roles {
		delete(payload, "policies")

		roles := m.Roles.Elements()

		// Transform roles: ["role_id"] -> [{"id": "role_id"}]
		roleObjects := make([]map[string]interface{}, 0, len(roles))
		for _, role := range roles {
			if !role.IsNull() && !role.IsUnknown() {
				roleString, ok := role.(types.String)
				if !ok {
					return nil, fmt.Errorf("unexpected role type: %T", role)
				}
				roleObjects = append(roleObjects, map[string]interface{}{
					"id": roleString.ValueString(),
				})
			}
		}

		if len(roleObjects) > 0 {
			payload["roles"] = roleObjects
		}
	}

	if configuredPermissionType == Policies {
		delete(payload, "roles")
	}

	// Ensure account_id is included for the API even though it's a path param
	if !m.AccountID.IsNull() && !m.AccountID.IsUnknown() {
		payload["account_id"] = m.AccountID.ValueString()
	}

	return json.Marshal(payload)
}

func unmarshalCustom(data []byte, configuredModel *AccountMemberModel) (*AccountMemberModel, error) {
	ctx := context.Background()
	var env AccountMemberResultEnvelope
	if err := apijson.Unmarshal(data, &env); err != nil {
		return nil, err
	}
	result := &env.Result

	result.AccountID = configuredModel.AccountID

	return parsePoliciesAndRoles(ctx, data, result)
}

func unmarshalComputedCustom(data []byte, configuredModel *AccountMemberModel) (*AccountMemberModel, error) {
	ctx := context.Background()
	var env AccountMemberResultEnvelope
	if err := apijson.UnmarshalComputed(data, &env); err != nil {
		return nil, err
	}
	result := &env.Result

	result.AccountID = configuredModel.AccountID
	if result.Email.IsNull() && !configuredModel.Email.IsNull() {
		// Preserve required fields from configured model to avoid replacement
		result.Email = configuredModel.Email
	}

	// Allow the status from the API to override the state,
    // even if not explicitly set in the config.
    if result.Status.IsNull() && !configuredModel.Status.IsNull() {
        result.Status = configuredModel.Status
    }
    // If result.Status has a value from the API (e.g., "accepted"),
    // we let it stay so it updates the Terraform state.

	return parsePoliciesAndRoles(ctx, data, result)
}

func parsePoliciesAndRoles(ctx context.Context, data []byte, result *AccountMemberModel) (*AccountMemberModel, error) {
	// Extract role IDs from GET response (roles come as objects with id field)
	var fullResponse struct {
		Result struct {
			Policies []AccountMemberPoliciesModel `json:"policies"`
			Roles    []struct {
				ID types.String `json:"id"`
			} `json:"roles"`
		} `json:"result"`
	}

	err := apijson.Unmarshal(data, &fullResponse)
	if err != nil {
		return nil, err
	}

	roleIDs := make([]types.String, 0, len(fullResponse.Result.Roles))
	for _, role := range fullResponse.Result.Roles {
		roleIDs = append(roleIDs, role.ID)
	}

	roleSet, diags := customfield.NewSet[types.String](ctx, roleIDs)
	if !diags.HasError() {
		result.Roles = roleSet
	} else {
		return result, fmt.Errorf("failed to parse roles")
	}

	policies := append([]AccountMemberPoliciesModel(nil), fullResponse.Result.Policies...)
	policiesSet, diags := customfield.NewObjectSet(ctx, policies)
	if !diags.HasError() {
		result.Policies = policiesSet
	} else {
		return result, fmt.Errorf("failed to parse policies")
	}

	return result, nil
}
