// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_lan

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitSiteLANDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"lan_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"site_id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"ha_link": schema.BoolAttribute{
				Description: "mark true to use this LAN for HA probing. only works for site with HA turned on. only one LAN can be set as the ha_link.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"physport": schema.Int64Attribute{
				Computed: true,
			},
			"vlan_tag": schema.Int64Attribute{
				Description: "VLAN port number.",
				Computed:    true,
			},
			"nat": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANNatDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"static_prefix": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Computed:    true,
					},
				},
			},
			"routed_subnets": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[MagicTransitSiteLANRoutedSubnetsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"next_hop": schema.StringAttribute{
							Description: "A valid IPv4 address.",
							Computed:    true,
						},
						"prefix": schema.StringAttribute{
							Description: "A valid CIDR notation representing an IP range.",
							Computed:    true,
						},
						"nat": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANRoutedSubnetsNatDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"static_prefix": schema.StringAttribute{
									Description: "A valid CIDR notation representing an IP range.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
			"static_addressing": schema.SingleNestedAttribute{
				Description: "If the site is not configured in high availability mode, this configuration is optional (if omitted, use DHCP). However, if in high availability mode, static_address is required along with secondary and virtual address.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[MagicTransitSiteLANStaticAddressingDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Computed:    true,
					},
					"dhcp_relay": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANStaticAddressingDHCPRelayDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"server_addresses": schema.ListAttribute{
								Description: "List of DHCP server IPs.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
					"dhcp_server": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANStaticAddressingDHCPServerDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"dhcp_pool_end": schema.StringAttribute{
								Description: "A valid IPv4 address.",
								Computed:    true,
							},
							"dhcp_pool_start": schema.StringAttribute{
								Description: "A valid IPv4 address.",
								Computed:    true,
							},
							"dns_server": schema.StringAttribute{
								Description: "A valid IPv4 address.",
								Computed:    true,
							},
							"dns_servers": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"reservations": schema.MapAttribute{
								Description: "Mapping of MAC addresses to IP addresses",
								Computed:    true,
								CustomType:  customfield.NewMapType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
					"secondary_address": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Computed:    true,
					},
					"virtual_address": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *MagicTransitSiteLANDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicTransitSiteLANDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
