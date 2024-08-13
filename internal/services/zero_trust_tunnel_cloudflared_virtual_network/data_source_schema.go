// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_virtual_network

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &ZeroTrustTunnelCloudflaredVirtualNetworkDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"comment": schema.StringAttribute{
				Description: "Optional remark describing the virtual network.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was created.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"deleted_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "UUID of the virtual network.",
				Optional:    true,
			},
			"is_default_network": schema.BoolAttribute{
				Description: "If `true`, this virtual network is the default for the account.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "A user-friendly name for the virtual network.",
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Cloudflare account ID",
						Required:    true,
					},
					"id": schema.StringAttribute{
						Description: "UUID of the virtual network.",
						Optional:    true,
					},
					"is_default": schema.BoolAttribute{
						Description: "If `true`, only include the default virtual network. If `false`, exclude the default virtual network. If empty, all virtual networks will be included.",
						Optional:    true,
					},
					"is_deleted": schema.BoolAttribute{
						Description: "If `true`, only include deleted virtual networks. If `false`, exclude deleted virtual networks. If empty, all virtual networks will be included.",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "A user-friendly name for the virtual network.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustTunnelCloudflaredVirtualNetworkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustTunnelCloudflaredVirtualNetworkDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
