package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceVirtualNetworkSchema returns the minimal v4 schema used to parse
// cloudflare_tunnel_virtual_network / cloudflare_zero_trust_tunnel_virtual_network state
// during MoveState and UpgradeState.
// Version is not set (defaults to 0, matching the v4 provider behaviour).
func SourceVirtualNetworkSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"is_default": schema.BoolAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"comment": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"is_default_network": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"deleted_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}
