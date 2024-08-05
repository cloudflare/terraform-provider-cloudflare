// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &TunnelRouteDataSource{}

func (d *TunnelRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"comment": schema.StringAttribute{
				Description: "Optional remark describing the route.",
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
				Description: "UUID of the route.",
				Optional:    true,
			},
			"network": schema.StringAttribute{
				Description: "The private IPv4 or IPv6 range connected by the route, in CIDR notation.",
				Optional:    true,
			},
			"tun_type": schema.StringAttribute{
				Description: "The type of tunnel.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("cfd_tunnel", "warp_connector", "ip_sec", "gre", "cni"),
				},
			},
			"tunnel_id": schema.StringAttribute{
				Description: "UUID of the tunnel.",
				Optional:    true,
			},
			"tunnel_name": schema.StringAttribute{
				Description: "A user-friendly name for a tunnel.",
				Optional:    true,
			},
			"virtual_network_id": schema.StringAttribute{
				Description: "UUID of the virtual network.",
				Optional:    true,
			},
			"virtual_network_name": schema.StringAttribute{
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
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.AtLeast(1),
						},
					},
					"per_page": schema.Float64Attribute{
						Description: "Number of results to display.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(1, 1000),
						},
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
				},
			},
		},
	}
}

func (d *TunnelRouteDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
