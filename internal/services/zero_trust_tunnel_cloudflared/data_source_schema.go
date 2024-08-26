// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustTunnelCloudflaredDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID",
				Optional:    true,
			},
			"tunnel_id": schema.StringAttribute{
				Description: "UUID of the tunnel.",
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
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
				},
			},
		},
	}
}

func (d *ZeroTrustTunnelCloudflaredDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustTunnelCloudflaredDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("tunnel_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("tunnel_id")),
	}
}
