package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareAccountModel represents the source cloudflare_account state structure.
// This corresponds to the legacy (SDKv2 v4) cloudflare provider.
// Schema version: 0 (implicit in v4)
// Resource type: cloudflare_account
//
// In v4, enforce_twofactor was a top-level boolean attribute.
// In v5, it moved into a nested settings object.
type SourceCloudflareAccountModel struct {
	ID               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Type             types.String `tfsdk:"type"`
	EnforceTwofactor types.Bool   `tfsdk:"enforce_twofactor"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetAccountSettingsModel represents the nested settings object in v5.
type TargetAccountSettingsModel struct {
	AbuseContactEmail types.String `tfsdk:"abuse_contact_email"`
	EnforceTwofactor  types.Bool   `tfsdk:"enforce_twofactor"`
}

// TargetAccountUnitModel represents the nested unit object in v5.
type TargetAccountUnitModel struct {
	ID types.String `tfsdk:"id"`
}

// TargetAccountManagedByModel represents the nested managed_by object in v5.
type TargetAccountManagedByModel struct {
	ParentOrgID   types.String `tfsdk:"parent_org_id"`
	ParentOrgName types.String `tfsdk:"parent_org_name"`
}

// TargetAccountModel represents the target cloudflare_account state structure (v500).
// This corresponds to the current provider implementation.
// Schema version: 500 (target)
// Resource type: cloudflare_account (same name, no rename)
//
// Key differences from source (v4):
// - enforce_twofactor: top-level bool → settings.enforce_twofactor (nested)
// - settings: NEW nested object (contains enforce_twofactor + abuse_contact_email)
// - unit: NEW nested object (tenant unit info)
// - managed_by: NEW nested object (parent container details)
// - created_on: NEW computed timestamp
//
// Note: This must match the model in account/model.go structurally.
// We duplicate it here to keep the migration package self-contained,
// but use simpler types to avoid customfield dependency issues during migration.
type TargetAccountModel struct {
	ID        types.String                                          `tfsdk:"id"`
	Unit      customfield.NestedObject[TargetAccountUnitModel]      `tfsdk:"unit"`
	Name      types.String                                          `tfsdk:"name"`
	Type      types.String                                          `tfsdk:"type"`
	ManagedBy customfield.NestedObject[TargetAccountManagedByModel] `tfsdk:"managed_by"`
	Settings  customfield.NestedObject[TargetAccountSettingsModel]  `tfsdk:"settings"`
	CreatedOn timetypes.RFC3339                                     `tfsdk:"created_on"`
}
