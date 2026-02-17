package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceAccessServiceTokenModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_access_service_token (deprecated name)
//
// This model includes all v4 fields, including those removed in v5 (min_days_for_renewal).
type SourceAccessServiceTokenModel struct {
	ID                            types.String `tfsdk:"id"`
	AccountID                     types.String `tfsdk:"account_id"`
	ZoneID                        types.String `tfsdk:"zone_id"`
	Name                          types.String `tfsdk:"name"`
	ClientID                      types.String `tfsdk:"client_id"`
	ClientSecret                  types.String `tfsdk:"client_secret"`
	ExpiresAt                     types.String `tfsdk:"expires_at"`
	MinDaysForRenewal             types.Int64  `tfsdk:"min_days_for_renewal"`
	Duration                      types.String `tfsdk:"duration"`
	ClientSecretVersion           types.Int64  `tfsdk:"client_secret_version"`
	PreviousClientSecretExpiresAt types.String `tfsdk:"previous_client_secret_expires_at"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetAccessServiceTokenModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_zero_trust_access_service_token
//
// Note: This matches the model in the parent package's model.go file.
// We duplicate it here to keep the migration package self-contained.
type TargetAccessServiceTokenModel struct {
	ID                            types.String      `tfsdk:"id"`
	AccountID                     types.String      `tfsdk:"account_id"`
	ZoneID                        types.String      `tfsdk:"zone_id"`
	Name                          types.String      `tfsdk:"name"`
	PreviousClientSecretExpiresAt timetypes.RFC3339 `tfsdk:"previous_client_secret_expires_at"`
	ClientSecretVersion           types.Float64     `tfsdk:"client_secret_version"`
	Duration                      types.String      `tfsdk:"duration"`
	ClientID                      types.String      `tfsdk:"client_id"`
	ClientSecret                  types.String      `tfsdk:"client_secret"`
	ExpiresAt                     timetypes.RFC3339 `tfsdk:"expires_at"`
}
