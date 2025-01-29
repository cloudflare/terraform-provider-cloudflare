// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustTunnelCloudflaredsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID",
				Required:    true,
			},
			"exclude_prefix": schema.StringAttribute{
				Optional: true,
			},
			"existed_at": schema.StringAttribute{
				Description: "If provided, include only tunnels that were created (and not deleted) before this time.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"include_prefix": schema.StringAttribute{
				Optional: true,
			},
			"is_deleted": schema.BoolAttribute{
				Description: "If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If empty, all tunnels will be included.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "A user-friendly name for a tunnel.",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "The status of the tunnel. Valid values are `inactive` (tunnel has never been run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy state), `healthy` (tunnel is active and able to serve traffic), or `down` (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"inactive",
						"degraded",
						"healthy",
						"down",
					),
				},
			},
			"uuid": schema.StringAttribute{
				Description: "UUID of the tunnel.",
				Optional:    true,
			},
			"was_active_at": schema.StringAttribute{
				Optional:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"was_inactive_at": schema.StringAttribute{
				Optional:   true,
				CustomType: timetypes.RFC3339Type{},
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustTunnelCloudflaredsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID of the tunnel.",
							Computed:    true,
						},
						"account_tag": schema.StringAttribute{
							Description: "Cloudflare account ID",
							Computed:    true,
						},
						"connections": schema.ListNestedAttribute{
							Description: "The Cloudflare Tunnel connections between your origin and Cloudflare's edge.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustTunnelCloudflaredsConnectionsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "UUID of the Cloudflare Tunnel connection.",
										Computed:    true,
									},
									"client_id": schema.StringAttribute{
										Description: "UUID of the Cloudflare Tunnel connector.",
										Computed:    true,
									},
									"client_version": schema.StringAttribute{
										Description: "The cloudflared version used to establish this connection.",
										Computed:    true,
									},
									"colo_name": schema.StringAttribute{
										Description: "The Cloudflare data center used for this connection.",
										Computed:    true,
									},
									"is_pending_reconnect": schema.BoolAttribute{
										Description: "Cloudflare continues to track connections for several minutes after they disconnect. This is an optimization to improve latency and reliability of reconnecting.  If `true`, the connection has disconnected but is still being tracked. If `false`, the connection is actively serving traffic.",
										Computed:    true,
									},
									"opened_at": schema.StringAttribute{
										Description: "Timestamp of when the connection was established.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"origin_ip": schema.StringAttribute{
										Description: "The public IP address of the host running cloudflared.",
										Computed:    true,
									},
									"uuid": schema.StringAttribute{
										Description: "UUID of the Cloudflare Tunnel connection.",
										Computed:    true,
									},
								},
							},
						},
						"conns_active_at": schema.StringAttribute{
							Description: "Timestamp of when the tunnel established at least one connection to Cloudflare's edge. If `null`, the tunnel is inactive.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"conns_inactive_at": schema.StringAttribute{
							Description: "Timestamp of when the tunnel became inactive (no connections to Cloudflare's edge). If `null`, the tunnel is active.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
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
						"metadata": schema.StringAttribute{
							Description: "Metadata associated with the tunnel.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"name": schema.StringAttribute{
							Description: "A user-friendly name for a tunnel.",
							Computed:    true,
						},
						"remote_config": schema.BoolAttribute{
							Description: "If `true`, the tunnel can be configured remotely from the Zero Trust dashboard. If `false`, the tunnel must be configured locally on the origin machine.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "The status of the tunnel. Valid values are `inactive` (tunnel has never been run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy state), `healthy` (tunnel is active and able to serve traffic), or `down` (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"inactive",
									"degraded",
									"healthy",
									"down",
								),
							},
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
					},
				},
			},
		},
	}
}

func (d *ZeroTrustTunnelCloudflaredsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustTunnelCloudflaredsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
