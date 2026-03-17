package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TransformV4toV500 transforms v4 SDKv2 state to v5 Plugin Framework state.
//
// Field transformations:
//   - email_address (v4) → email (v5): rename
//   - role_ids (v4 TypeSet) → roles (v5 customfield.Set): rename + type conversion
//   - policies: not in v4, set to null
//   - user: not in v4, set to null
func TransformV4toV500(ctx context.Context, source SourceV4AccountMemberModel) (*TargetAccountMemberModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Convert role_ids (types.Set) to roles (customfield.Set[types.String])
	var rolesSet customfield.Set[types.String]
	if !source.RoleIDs.IsNull() && !source.RoleIDs.IsUnknown() {
		var roleStrings []types.String
		d := source.RoleIDs.ElementsAs(ctx, &roleStrings, false)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}

		var err diag.Diagnostics
		rolesSet, err = customfield.NewSet[types.String](ctx, roleStrings)
		diags.Append(err...)
		if diags.HasError() {
			return nil, diags
		}
	} else {
		rolesSet = customfield.NullSet[types.String](ctx)
	}

	// Initialize null policies and user (v4 didn't have these fields)
	policiesNull := customfield.NullObjectSet[TargetAccountMemberPolicies](ctx)
	userNull := customfield.NullObject[TargetAccountMemberUser](ctx)

	target := &TargetAccountMemberModel{
		ID:        source.ID,
		AccountID: source.AccountID,
		Email:     source.EmailAddress, // email_address → email
		Roles:     rolesSet,            // role_ids → roles
		Status:    source.Status,
		Policies:  policiesNull,
		User:      userNull,
	}

	return target, diags
}

// TransformV513toV500 transforms v5.13 state to v500 state.
//
// Field transformations:
//   - roles: List → Set
//   - policies: List → Set, remove 'id' field from each policy
//   - permission_groups: List → Set
//   - resource_groups: List → Set
func TransformV513toV500(ctx context.Context, source SourceV513AccountMemberModel) (*TargetAccountMemberModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Convert roles (types.List) to roles (customfield.Set[types.String])
	var rolesSet customfield.Set[types.String]
	if !source.Roles.IsNull() && !source.Roles.IsUnknown() {
		var roleStrings []types.String
		d := source.Roles.ElementsAs(ctx, &roleStrings, false)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}

		var err diag.Diagnostics
		rolesSet, err = customfield.NewSet[types.String](ctx, roleStrings)
		diags.Append(err...)
		if diags.HasError() {
			return nil, diags
		}
	} else {
		rolesSet = customfield.NullSet[types.String](ctx)
	}

	// Convert policies (types.List) to policies (customfield.NestedObjectSet)
	var policiesSet customfield.NestedObjectSet[TargetAccountMemberPolicies]
	if !source.Policies.IsNull() && !source.Policies.IsUnknown() {
		var sourcePolicies []SourceV513AccountMemberPolicy
		d := source.Policies.ElementsAs(ctx, &sourcePolicies, false)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}

		var targetPolicies []TargetAccountMemberPolicies
		for _, sp := range sourcePolicies {
			// Convert permission_groups (List → Set)
			var permGroups customfield.NestedObjectSet[TargetAccountMemberPermissionGroups]
			if !sp.PermissionGroups.IsNull() && !sp.PermissionGroups.IsUnknown() {
				var sourcePermGroups []SourceV513AccountMemberPermissionGroup
				d := sp.PermissionGroups.ElementsAs(ctx, &sourcePermGroups, false)
				diags.Append(d...)
				if diags.HasError() {
					return nil, diags
				}

				var targetPermGroups []TargetAccountMemberPermissionGroups
				for _, pg := range sourcePermGroups {
					targetPermGroups = append(targetPermGroups, TargetAccountMemberPermissionGroups{
						ID: pg.ID,
					})
				}
				var err diag.Diagnostics
				permGroups, err = customfield.NewObjectSet[TargetAccountMemberPermissionGroups](ctx, targetPermGroups)
				diags.Append(err...)
				if diags.HasError() {
					return nil, diags
				}
			} else {
				permGroups = customfield.NullObjectSet[TargetAccountMemberPermissionGroups](ctx)
			}

			// Convert resource_groups (List → Set)
			var resGroups customfield.NestedObjectSet[TargetAccountMemberResourceGroups]
			if !sp.ResourceGroups.IsNull() && !sp.ResourceGroups.IsUnknown() {
				var sourceResGroups []SourceV513AccountMemberResourceGroup
				d := sp.ResourceGroups.ElementsAs(ctx, &sourceResGroups, false)
				diags.Append(d...)
				if diags.HasError() {
					return nil, diags
				}

				var targetResGroups []TargetAccountMemberResourceGroups
				for _, rg := range sourceResGroups {
					targetResGroups = append(targetResGroups, TargetAccountMemberResourceGroups{
						ID: rg.ID,
					})
				}
				var err diag.Diagnostics
				resGroups, err = customfield.NewObjectSet[TargetAccountMemberResourceGroups](ctx, targetResGroups)
				diags.Append(err...)
				if diags.HasError() {
					return nil, diags
				}
			} else {
				resGroups = customfield.NullObjectSet[TargetAccountMemberResourceGroups](ctx)
			}

			// Note: we drop the 'id' field from sp - it's not in the target schema
			targetPolicies = append(targetPolicies, TargetAccountMemberPolicies{
				Access:           sp.Access,
				PermissionGroups: permGroups,
				ResourceGroups:   resGroups,
			})
		}

		var err diag.Diagnostics
		policiesSet, err = customfield.NewObjectSet[TargetAccountMemberPolicies](ctx, targetPolicies)
		diags.Append(err...)
		if diags.HasError() {
			return nil, diags
		}
	} else {
		policiesSet = customfield.NullObjectSet[TargetAccountMemberPolicies](ctx)
	}

	// Convert user (if present)
	var userNested customfield.NestedObject[TargetAccountMemberUser]
	if source.User != nil {
		var err diag.Diagnostics
		userNested, err = customfield.NewObject[TargetAccountMemberUser](ctx, &TargetAccountMemberUser{
			Email:                          source.User.Email,
			ID:                             source.User.ID,
			FirstName:                      source.User.FirstName,
			LastName:                       source.User.LastName,
			TwoFactorAuthenticationEnabled: source.User.TwoFactorAuthenticationEnabled,
		})
		diags.Append(err...)
		if diags.HasError() {
			return nil, diags
		}
	} else {
		userNested = customfield.NullObject[TargetAccountMemberUser](ctx)
	}

	target := &TargetAccountMemberModel{
		ID:        source.ID,
		AccountID: source.AccountID,
		Email:     source.Email,
		Status:    source.Status,
		Roles:     rolesSet,
		Policies:  policiesSet,
		User:      userNested,
	}

	return target, diags
}
