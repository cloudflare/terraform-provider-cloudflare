package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceVirtualNetworkModel mirrors the v4 Plugin Framework state for
// cloudflare_tunnel_virtual_network / cloudflare_zero_trust_tunnel_virtual_network.
// Schema version: 0 (no explicit version set in v4 provider).
type SourceVirtualNetworkModel struct {
	ID               types.String      `tfsdk:"id"`
	AccountID        types.String      `tfsdk:"account_id"`
	IsDefault        types.Bool        `tfsdk:"is_default"`
	Name             types.String      `tfsdk:"name"`
	Comment          types.String      `tfsdk:"comment"`
	IsDefaultNetwork types.Bool        `tfsdk:"is_default_network"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at"`
}

// TargetVirtualNetworkModel mirrors the v5 Plugin Framework model for
// cloudflare_zero_trust_tunnel_cloudflared_virtual_network.
// Defined locally to avoid an import cycle with the parent service package.
// Schema version: 500 (with TF_MIG_TEST=1) or 1 (production).
type TargetVirtualNetworkModel struct {
	ID               types.String      `tfsdk:"id"`
	AccountID        types.String      `tfsdk:"account_id"`
	IsDefault        types.Bool        `tfsdk:"is_default"`
	Name             types.String      `tfsdk:"name"`
	Comment          types.String      `tfsdk:"comment"`
	IsDefaultNetwork types.Bool        `tfsdk:"is_default_network"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at"`
	DeletedAt        timetypes.RFC3339 `tfsdk:"deleted_at"`
}
