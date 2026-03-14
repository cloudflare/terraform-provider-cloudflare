package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareDeviceManagedNetworksModel represents the source cloudflare_device_managed_networks state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
// Used by both MoveState (Terraform 1.8+) and UpgradeFromLegacyV0 (Terraform < 1.8) to parse legacy state.
type SourceCloudflareDeviceManagedNetworksModel struct {
	ID        types.String        `tfsdk:"id"`
	AccountID types.String        `tfsdk:"account_id"`
	Name      types.String        `tfsdk:"name"`
	Type      types.String        `tfsdk:"type"`
	Config    []SourceConfigModel `tfsdk:"config"` // TypeList MaxItems:1 in v4 (stored as array)
}

// SourceConfigModel represents the source config block structure.
// In the legacy provider, this is a list with MaxItems: 1, so it's stored as a single-element array.
type SourceConfigModel struct {
	TLSSockaddr types.String `tfsdk:"tls_sockaddr"`
	Sha256      types.String `tfsdk:"sha256"`
}

// TargetZeroTrustDeviceManagedNetworksModel represents the target cloudflare_zero_trust_device_managed_networks state structure (v500).
// Must match zero_trust_device_managed_networks.ZeroTrustDeviceManagedNetworksModel structure exactly.
type TargetZeroTrustDeviceManagedNetworksModel struct {
	ID        types.String       `tfsdk:"id"`
	NetworkID types.String       `tfsdk:"network_id"` // New computed field in v5
	AccountID types.String       `tfsdk:"account_id"`
	Name      types.String       `tfsdk:"name"`
	Type      types.String       `tfsdk:"type"`
	Config    *TargetConfigModel `tfsdk:"config"` // SingleNestedAttribute in v5 (stored as object pointer)
}

// TargetConfigModel represents the target config nested object (v500).
// Must match zero_trust_device_managed_networks.ZeroTrustDeviceManagedNetworksConfigModel structure exactly.
type TargetConfigModel struct {
	TLSSockaddr types.String `tfsdk:"tls_sockaddr"`
	Sha256      types.String `tfsdk:"sha256"`
}
