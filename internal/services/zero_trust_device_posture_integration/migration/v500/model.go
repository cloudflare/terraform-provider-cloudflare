package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareDevicePostureIntegrationModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (SDKv2 implicit)
// Resource type: cloudflare_device_posture_integration
type SourceCloudflareDevicePostureIntegrationModel struct {
	ID         types.String        `tfsdk:"id"`
	AccountID  types.String        `tfsdk:"account_id"`
	Name       types.String        `tfsdk:"name"`
	Type       types.String        `tfsdk:"type"`
	Identifier types.String        `tfsdk:"identifier"` // REMOVED in v5
	Interval   types.String        `tfsdk:"interval"`   // Optional in v4, Required in v5
	Config     []SourceConfigModel `tfsdk:"config"`     // TypeList MaxItems:1 in v4
}

// SourceConfigModel represents the nested config structure from v4.x provider.
// Stored as TypeList MaxItems:1 in v4, so state has this as an array.
type SourceConfigModel struct {
	APIURL             types.String `tfsdk:"api_url"`
	AuthURL            types.String `tfsdk:"auth_url"`
	ClientID           types.String `tfsdk:"client_id"`
	ClientSecret       types.String `tfsdk:"client_secret"`
	CustomerID         types.String `tfsdk:"customer_id"`
	ClientKey          types.String `tfsdk:"client_key"`
	AccessClientID     types.String `tfsdk:"access_client_id"`
	AccessClientSecret types.String `tfsdk:"access_client_secret"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetZeroTrustDevicePostureIntegrationModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_zero_trust_device_posture_integration
//
// Note: This matches the model in the parent package's model.go file.
// We duplicate it here to keep the migration package self-contained.
type TargetZeroTrustDevicePostureIntegrationModel struct {
	ID        types.String       `tfsdk:"id"`
	AccountID types.String       `tfsdk:"account_id"`
	Name      types.String       `tfsdk:"name"`
	Type      types.String       `tfsdk:"type"`
	Interval  types.String       `tfsdk:"interval"` // Required in v5
	Config    *TargetConfigModel `tfsdk:"config"`   // SingleNested (pointer) in v5, Required
}

// TargetConfigModel represents the nested config structure from v5.x+ provider.
// Stored as SingleNestedAttribute in v5 (pointer to object).
type TargetConfigModel struct {
	APIURL             types.String `tfsdk:"api_url"`
	AuthURL            types.String `tfsdk:"auth_url"`
	ClientID           types.String `tfsdk:"client_id"`
	ClientSecret       types.String `tfsdk:"client_secret"`
	CustomerID         types.String `tfsdk:"customer_id"`
	ClientKey          types.String `tfsdk:"client_key"`
	AccessClientID     types.String `tfsdk:"access_client_id"`
	AccessClientSecret types.String `tfsdk:"access_client_secret"`
}
