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
