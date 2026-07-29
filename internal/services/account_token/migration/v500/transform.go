package v500

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts v0 source state to v500 target state.
//
// Key transformations:
// 1. policies[].resources: map[string]string → JSON-encoded string
// 2. policies[].id: removed (computed field, no longer in schema)
// 3. policies[].permission_groups[].meta: removed (computed field)
// 4. policies[].permission_groups[].name: removed (computed field)
// 5. All other fields: direct pass-through
func Transform(ctx context.Context, source SourceAccountTokenModelV0) (*TargetAccountTokenModelV500, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Convert policies
	upgradedPolicies := make([]TargetPolicyV500, 0, len(source.Policies))
	for _, oldPolicy := range source.Policies {
		// Convert resources map to JSON string
		resourcesMap := make(map[string]string)
		for k, v := range oldPolicy.Resources {
			if !v.IsNull() && !v.IsUnknown() {
				resourcesMap[k] = v.ValueString()
			}
		}

		resourcesJSON, err := json.Marshal(resourcesMap)
		if err != nil {
			diags.AddError(
				"State Upgrade Error",
				"Failed to convert resources to JSON: "+err.Error(),
			)
			return nil, diags
		}

		// Convert permission groups (removing computed fields: meta, name)
		upgradedPermGroups := make([]TargetPermissionGroupV500, 0, len(oldPolicy.PermissionGroups))
		for _, oldPG := range oldPolicy.PermissionGroups {
			upgradedPermGroups = append(upgradedPermGroups, TargetPermissionGroupV500{
				ID: oldPG.ID,
			})
		}

		upgradedPolicies = append(upgradedPolicies, TargetPolicyV500{
			Effect:           oldPolicy.Effect,
			PermissionGroups: upgradedPermGroups,
			Resources:        types.StringValue(string(resourcesJSON)),
		})
		// Note: policy.ID is intentionally dropped — no longer in v500 schema
	}

	target := &TargetAccountTokenModelV500{
		AccountID:  source.AccountID,
		ID:         source.ID,
		IssuedOn:   source.IssuedOn,
		ModifiedOn: source.ModifiedOn,
		Name:       source.Name,
		Policies:   upgradedPolicies,
		Status:     source.Status,
		Value:      source.Value,
		NotBefore:  source.NotBefore,
		ExpiresOn:  source.ExpiresOn,
		Condition:  source.Condition,
		LastUsedOn: source.LastUsedOn,
	}

	return target, diags
}
