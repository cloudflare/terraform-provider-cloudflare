package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceTunnelRouteSchema returns the legacy cloudflare_tunnel_route schema (schema_version=0).
// This is used by MoveState and UpgradeFromV4 to parse state from the legacy SDKv2 provider.
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/sdkv2provider/schema_cloudflare_tunnel_route.go
func SourceTunnelRouteSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"tunnel_id": schema.StringAttribute{
				Required: true,
			},
			"network": schema.StringAttribute{
				Required: true,
			},
			"comment": schema.StringAttribute{
				Optional: true,
			},
			"virtual_network_id": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}
