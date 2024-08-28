// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_ipsec_tunnel

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicWANIPSECTunnelDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"ipsec_tunnel_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"ipsec_tunnel": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"cloudflare_endpoint": schema.StringAttribute{
						Description: "The IP address assigned to the Cloudflare side of the IPsec tunnel.",
						Computed:    true,
					},
					"interface_address": schema.StringAttribute{
						Description: "A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side of the tunnel. Select the subnet from the following private IP space: 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "The name of the IPsec tunnel. The name cannot share a name with other tunnels.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Tunnel identifier tag.",
						Computed:    true,
					},
					"allow_null_cipher": schema.BoolAttribute{
						Description: "When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel (Phase 2).",
						Computed:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "The date and time the tunnel was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"customer_endpoint": schema.StringAttribute{
						Description: "The IP address assigned to the customer side of the IPsec tunnel. Not required, but must be set for proactive traceroutes to work.",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "An optional description forthe IPsec tunnel.",
						Computed:    true,
					},
					"modified_on": schema.StringAttribute{
						Description: "The date and time the tunnel was last modified.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"psk_metadata": schema.SingleNestedAttribute{
						Description: "The PSK metadata that includes when the PSK was generated.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[MagicWANIPSECTunnelIPSECTunnelPSKMetadataDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"last_generated_on": schema.StringAttribute{
								Description: "The date and time the tunnel was last modified.",
								Computed:    true,
								CustomType:  timetypes.RFC3339Type{},
							},
						},
					},
					"replay_protection": schema.BoolAttribute{
						Description: "If `true`, then IPsec replay protection will be supported in the Cloudflare-to-customer direction.",
						Computed:    true,
					},
					"tunnel_health_check": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[MagicWANIPSECTunnelIPSECTunnelTunnelHealthCheckDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Determines whether to run healthchecks for a tunnel.",
								Computed:    true,
							},
							"rate": schema.StringAttribute{
								Description: "How frequent the health check is run. The default value is `mid`.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"low",
										"mid",
										"high",
									),
								},
							},
							"target": schema.StringAttribute{
								Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`.",
								Computed:    true,
							},
							"type": schema.StringAttribute{
								Description: "The type of healthcheck to run, reply or request. The default value is `reply`.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("reply", "request"),
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *MagicWANIPSECTunnelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicWANIPSECTunnelDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
