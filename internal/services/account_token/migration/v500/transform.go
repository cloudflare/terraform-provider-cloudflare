package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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
		// Convert permission groups (preserving meta and name)
		upgradedPermGroups := make([]TargetPermissionGroupV500, 0, len(oldPolicy.PermissionGroups))
		for _, oldPG := range oldPolicy.PermissionGroups {
			pg := TargetPermissionGroupV500{
				ID:   oldPG.ID,
				Name: oldPG.Name,
			}
			if oldPG.Meta != nil {
				pg.Meta = &TargetMetaV500{
					Key:   oldPG.Meta.Key,
					Value: oldPG.Meta.Value,
				}
			}
			upgradedPermGroups = append(upgradedPermGroups, pg)
		}

		upgradedPolicies = append(upgradedPolicies, TargetPolicyV500{
			ID:               oldPolicy.ID,
			Effect:           oldPolicy.Effect,
			PermissionGroups: upgradedPermGroups,
			Resources:        oldPolicy.Resources,
		})
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
