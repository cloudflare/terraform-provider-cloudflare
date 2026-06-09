// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_warp_connector_config

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustTunnelWARPConnectorConfigDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Cloudflare One Connector: WARP Read",
				"Cloudflare One Connector: WARP Write",
				"Cloudflare One Connectors Read",
				"Cloudflare One Connectors Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"tunnel_id": schema.StringAttribute{
				Description: "UUID of the tunnel.",
				Required:    true,
			},
			"configuration_version": schema.Int64Attribute{
				Description: "Monotonically increasing configuration version, incremented on each PUT.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"ha_mode": schema.StringAttribute{
				Description: "High-availability mode for the WARP Connector tunnel. `none` means HA is enabled but no provider is configured yet (newly created tunnels default to this). `disabled` means HA is explicitly turned off. `aws` uses AWS ENI move for failover. `local` uses virtual IPs (VIPs) on the local interface.\nAvailable values: \"none\", \"disabled\", \"aws\", \"local\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"none",
						"disabled",
						"aws",
						"local",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "Timestamp of the last update. Null if never updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"config": schema.SingleNestedAttribute{
				Description: "Provider-specific configuration. Present for `aws` and `local` modes.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelWARPConnectorConfigConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"fnr_id": schema.StringAttribute{
						Description: "Floating Network Resource ID — the secondary ENI that is moved between nodes on failover.",
						Computed:    true,
					},
					"vips": schema.ListNestedAttribute{
						Description: "VIPs to assign on the CloudflareWARP interface.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[ZeroTrustTunnelWARPConnectorConfigConfigVipsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Description: "Virtual IP address (IPv4 or IPv6).",
									Computed:    true,
								},
							},
						},
					},
					"vips_previous": schema.ListNestedAttribute{
						Description: "VIPs to clean up on demotion or version drift.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[ZeroTrustTunnelWARPConnectorConfigConfigVipsPreviousDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Description: "Virtual IP address (IPv4 or IPv6).",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustTunnelWARPConnectorConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustTunnelWARPConnectorConfigDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
