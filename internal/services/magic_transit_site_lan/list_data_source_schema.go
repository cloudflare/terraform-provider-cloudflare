// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_lan

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitSiteLANsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"site_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
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
				CustomType:  customfield.NewNestedObjectListType[MagicTransitSiteLANsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"ha_link": schema.BoolAttribute{
							Description: "mark true to use this LAN for HA probing. only works for site with HA turned on. only one LAN can be set as the ha_link.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"nat": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANsNatDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"static_prefix": schema.StringAttribute{
									Description: "A valid CIDR notation representing an IP range.",
									Computed:    true,
								},
							},
						},
						"physport": schema.Int64Attribute{
							Computed: true,
						},
						"routed_subnets": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[MagicTransitSiteLANsRoutedSubnetsDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANsRoutedSubnetsNatDataSourceModel](ctx),
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
						"site_id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"static_addressing": schema.SingleNestedAttribute{
							Description: "If the site is not configured in high availability mode, this configuration is optional (if omitted, use DHCP). However, if in high availability mode, static_address is required along with secondary and virtual address.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[MagicTransitSiteLANsStaticAddressingDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Description: "A valid CIDR notation representing an IP range.",
									Computed:    true,
								},
								"dhcp_relay": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANsStaticAddressingDHCPRelayDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANsStaticAddressingDHCPServerDataSourceModel](ctx),
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
						"vlan_tag": schema.Int64Attribute{
							Description: "VLAN port number.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *MagicTransitSiteLANsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MagicTransitSiteLANsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
