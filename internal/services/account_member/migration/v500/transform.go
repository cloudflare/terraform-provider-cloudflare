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
