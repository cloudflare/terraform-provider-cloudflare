// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &TunnelDataSource{}
var _ datasource.DataSourceWithValidateConfig = &TunnelDataSource{}

func (r TunnelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID",
				Optional:    true,
			},
			"tunnel_id": schema.StringAttribute{
				Description: "UUID of the tunnel.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "UUID of the tunnel.",
				Optional:    true,
			},
			"connections": schema.ListNestedAttribute{
				Description: "The tunnel connections between your origin and Cloudflare's edge.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"colo_name": schema.StringAttribute{
							Description: "The Cloudflare data center used for this connection.",
							Computed:    true,
							Optional:    true,
						},
						"is_pending_reconnect": schema.BoolAttribute{
							Description: "Cloudflare continues to track connections for several minutes after they disconnect. This is an optimization to improve latency and reliability of reconnecting.  If `true`, the connection has disconnected but is still being tracked. If `false`, the connection is actively serving traffic.",
							Computed:    true,
							Optional:    true,
						},
						"uuid": schema.StringAttribute{
							Description: "UUID of the Cloudflare Tunnel connection.",
							Computed:    true,
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was created.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "A user-friendly name for a tunnel.",
				Optional:    true,
			},
			"deleted_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
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
					},
					"include_prefix": schema.StringAttribute{
						Optional: true,
					},
					"is_deleted": schema.BoolAttribute{
						Description: "If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If empty, all tunnels will be included.",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "A user-friendly name for the tunnel.",
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
					"status": schema.StringAttribute{
						Description: "The status of the tunnel. Valid values are `inactive` (tunnel has never been run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy state), `healthy` (tunnel is active and able to serve traffic), or `down` (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("inactive", "degraded", "healthy", "down"),
						},
					},
					"tun_types": schema.StringAttribute{
						Description: "The types of tunnels to filter separated by a comma.",
						Optional:    true,
					},
					"uuid": schema.StringAttribute{
						Description: "UUID of the tunnel.",
						Optional:    true,
					},
					"was_active_at": schema.StringAttribute{
						Optional: true,
					},
					"was_inactive_at": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

func (r *TunnelDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *TunnelDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
