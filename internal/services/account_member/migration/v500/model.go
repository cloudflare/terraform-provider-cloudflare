package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceV4AccountMemberModel represents the v4 SDKv2 account_member state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
//
// Key differences from v5:
//   - email_address (v4) → email (v5)
//   - role_ids (v4) → roles (v5)
//   - policies and user don't exist in v4
type SourceV4AccountMemberModel struct {
	ID           types.String `tfsdk:"id"`
	AccountID    types.String `tfsdk:"account_id"`
	EmailAddress types.String `tfsdk:"email_address"` // v4 field name
	RoleIDs      types.Set    `tfsdk:"role_ids"`      // v4 field name, TypeSet of strings
	Status       types.String `tfsdk:"status"`
}

// ============================================================================
// Source Models (v5.13.0 - stepping stone release)
// ============================================================================

// SourceV513AccountMemberModel represents the v5.13.0 account_member state structure.
// This corresponds to schema_version=1 from the stepping stone release.
//
// Key differences from current v5:
//   - policies was ListNestedAttribute (now SetNestedAttribute)
//   - policies had an 'id' field (now removed)
//   - permission_groups was ListNestedAttribute (now SetNestedAttribute)
//   - resource_groups was ListNestedAttribute (now SetNestedAttribute)
//   - roles was ListAttribute (now SetAttribute)
type SourceV513AccountMemberModel struct {
	ID        types.String                 `tfsdk:"id"`
	AccountID types.String                 `tfsdk:"account_id"`
	Email     types.String                 `tfsdk:"email"`
	Status    types.String                 `tfsdk:"status"`
	Roles     types.List                   `tfsdk:"roles"`    // v5.13 used List, now Set
	Policies  types.List                   `tfsdk:"policies"` // v5.13 used List, now Set
	User      *SourceV513AccountMemberUser `tfsdk:"user"`
}

// SourceV513AccountMemberPolicy represents a single policy in v5.13 state.
type SourceV513AccountMemberPolicy struct {
	ID               types.String `tfsdk:"id"` // v5.13 had 'id' field, now removed
	Access           types.String `tfsdk:"access"`
	PermissionGroups types.List   `tfsdk:"permission_groups"` // v5.13 used List, now Set
	ResourceGroups   types.List   `tfsdk:"resource_groups"`   // v5.13 used List, now Set
}

// SourceV513AccountMemberPermissionGroup represents a permission group in v5.13 state.
type SourceV513AccountMemberPermissionGroup struct {
	ID types.String `tfsdk:"id"`
}

// SourceV513AccountMemberResourceGroup represents a resource group in v5.13 state.
type SourceV513AccountMemberResourceGroup struct {
	ID types.String `tfsdk:"id"`
}

// SourceV513AccountMemberUser represents the user object in v5.13 state.
type SourceV513AccountMemberUser struct {
	Email                          types.String `tfsdk:"email"`
	ID                             types.String `tfsdk:"id"`
	FirstName                      types.String `tfsdk:"first_name"`
	LastName                       types.String `tfsdk:"last_name"`
	TwoFactorAuthenticationEnabled types.Bool   `tfsdk:"two_factor_authentication_enabled"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetAccountMemberModel represents the current cloudflare_account_member state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_account_member
//
// Note: Duplicates AccountMemberModel from the parent package to keep the migration
// package self-contained and avoid import cycles.
type TargetAccountMemberModel struct {
	ID        types.String                                             `tfsdk:"id"`
	AccountID types.String                                             `tfsdk:"account_id"`
	Email     types.String                                             `tfsdk:"email"`
	Status    types.String                                             `tfsdk:"status"`
	Roles     customfield.Set[types.String]                            `tfsdk:"roles"`
	Policies  customfield.NestedObjectSet[TargetAccountMemberPolicies] `tfsdk:"policies"`
	User      customfield.NestedObject[TargetAccountMemberUser]        `tfsdk:"user"`
}

// TargetAccountMemberPolicies represents the nested policies object in v5.
type TargetAccountMemberPolicies struct {
	Access           types.String                                                     `tfsdk:"access"`
	PermissionGroups customfield.NestedObjectSet[TargetAccountMemberPermissionGroups] `tfsdk:"permission_groups"`
	ResourceGroups   customfield.NestedObjectSet[TargetAccountMemberResourceGroups]   `tfsdk:"resource_groups"`
}

// TargetAccountMemberPermissionGroups represents the nested permission_groups object in v5.
type TargetAccountMemberPermissionGroups struct {
	ID types.String `tfsdk:"id"`
}

// TargetAccountMemberResourceGroups represents the nested resource_groups object in v5.
type TargetAccountMemberResourceGroups struct {
	ID types.String `tfsdk:"id"`
}

// TargetAccountMemberUser represents the nested user object in v5.
type TargetAccountMemberUser struct {
	Email                          types.String `tfsdk:"email"`
	ID                             types.String `tfsdk:"id"`
	FirstName                      types.String `tfsdk:"first_name"`
	LastName                       types.String `tfsdk:"last_name"`
	TwoFactorAuthenticationEnabled types.Bool   `tfsdk:"two_factor_authentication_enabled"`
}
