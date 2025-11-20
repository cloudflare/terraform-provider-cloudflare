// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_network_hostname_route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustNetworkHostnameRouteDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The hostname route ID.",
				Computed:    true,
			},
			"hostname_route_id": schema.StringAttribute{
				Description: "The hostname route ID.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID",
				Required:    true,
			},
			"comment": schema.StringAttribute{
				Description: "An optional description of the hostname route.",
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
			"hostname": schema.StringAttribute{
				Description: "The hostname of the route.",
				Computed:    true,
			},
			"tunnel_id": schema.StringAttribute{
				Description: "UUID of the tunnel.",
				Computed:    true,
			},
			"tunnel_name": schema.StringAttribute{
				Description: "A user-friendly name for a tunnel.",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The hostname route ID.",
						Optional:    true,
					},
					"comment": schema.StringAttribute{
						Description: "If set, only list hostname routes with the given comment.",
						Optional:    true,
					},
					"existed_at": schema.StringAttribute{
						Description: "If provided, include only resources that were created (and not deleted) before this time. URL encoded.",
						Optional:    true,
					},
					"hostname": schema.StringAttribute{
						Description: "If set, only list hostname routes that contain a substring of the given value, the filter is case-insensitive.",
						Optional:    true,
					},
					"is_deleted": schema.BoolAttribute{
						Description: "If `true`, only return deleted hostname routes. If `false`, exclude deleted hostname routes.",
						Computed:    true,
						Optional:    true,
					},
					"tunnel_id": schema.StringAttribute{
						Description: "If set, only list hostname routes that point to a specific tunnel.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustNetworkHostnameRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustNetworkHostnameRouteDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("hostname_route_id"), path.MatchRoot("filter")),
	}
}
