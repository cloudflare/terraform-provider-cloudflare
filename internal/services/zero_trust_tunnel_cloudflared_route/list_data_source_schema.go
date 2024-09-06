// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_route

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustTunnelCloudflaredRoutesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID",
				Required:    true,
			},
			"comment": schema.StringAttribute{
				Description: "Optional remark describing the route.",
				Optional:    true,
			},
			"existed_at": schema.StringAttribute{
				Description: "If provided, include only tunnels that were created (and not deleted) before this time.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"is_deleted": schema.BoolAttribute{
				Description: "If `true`, only include deleted routes. If `false`, exclude deleted routes. If empty, all routes will be included.",
				Optional:    true,
			},
			"network_subset": schema.StringAttribute{
				Description: "If set, only list routes that are contained within this IP range.",
				Optional:    true,
			},
			"network_superset": schema.StringAttribute{
				Description: "If set, only list routes that contain this IP range.",
				Optional:    true,
			},
			"route_id": schema.StringAttribute{
				Description: "UUID of the route.",
				Optional:    true,
			},
			"tun_types": schema.StringAttribute{
				Description: "The types of tunnels to filter separated by a comma.",
				Optional:    true,
			},
			"tunnel_id": schema.StringAttribute{
				Description: "UUID of the tunnel.",
				Optional:    true,
			},
			"virtual_network_id": schema.StringAttribute{
				Description: "UUID of the virtual network.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustTunnelCloudflaredRoutesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID of the route.",
							Computed:    true,
						},
						"comment": schema.StringAttribute{
							Description: "Optional remark describing the route.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "Timestamp of when the resource was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"deleted_at": schema.StringAttribute{
							Description: "Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"network": schema.StringAttribute{
							Description: "The private IPv4 or IPv6 range connected by the route, in CIDR notation.",
							Computed:    true,
						},
						"tun_type": schema.StringAttribute{
							Description: "The type of tunnel.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"cfd_tunnel",
									"warp_connector",
									"ip_sec",
									"gre",
									"cni",
								),
							},
						},
						"tunnel_id": schema.StringAttribute{
							Description: "UUID of the tunnel.",
							Computed:    true,
						},
						"tunnel_name": schema.StringAttribute{
							Description: "A user-friendly name for a tunnel.",
							Computed:    true,
						},
						"virtual_network_id": schema.StringAttribute{
							Description: "UUID of the virtual network.",
							Computed:    true,
						},
						"virtual_network_name": schema.StringAttribute{
							Description: "A user-friendly name for the virtual network.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustTunnelCloudflaredRoutesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustTunnelCloudflaredRoutesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
