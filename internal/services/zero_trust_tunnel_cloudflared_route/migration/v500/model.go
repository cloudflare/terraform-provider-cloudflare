package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceTunnelRouteModel represents the source cloudflare_tunnel_route state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
// Used by both MoveState (Terraform 1.8+) and UpgradeFromV4 (Terraform < 1.8) to parse legacy state.
// Applies to both: cloudflare_tunnel_route and cloudflare_zero_trust_tunnel_route (v4 aliases).
type SourceTunnelRouteModel struct {
	ID               types.String `tfsdk:"id"`
	AccountID        types.String `tfsdk:"account_id"`
	TunnelID         types.String `tfsdk:"tunnel_id"`
	Network          types.String `tfsdk:"network"`
	Comment          types.String `tfsdk:"comment"`
	VirtualNetworkID types.String `tfsdk:"virtual_network_id"`
}

// TargetTunnelRouteModel represents the target cloudflare_zero_trust_tunnel_cloudflared_route state structure (v500).
type TargetTunnelRouteModel struct {
	ID               types.String      `tfsdk:"id"`
	AccountID        types.String      `tfsdk:"account_id"`
	Network          types.String      `tfsdk:"network"`
	TunnelID         types.String      `tfsdk:"tunnel_id"`
	Comment          types.String      `tfsdk:"comment"`
	VirtualNetworkID types.String      `tfsdk:"virtual_network_id"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at"`
}
