// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_wan

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitSiteWANsDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[MagicTransitSiteWANsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"health_check_rate": schema.StringAttribute{
							Description: "Magic WAN health check rate for tunnels created on this link. The default value is `mid`.\navailable values: \"low\", \"mid\", \"high\"",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"low",
									"mid",
									"high",
								),
							},
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"physport": schema.Int64Attribute{
							Computed: true,
						},
						"priority": schema.Int64Attribute{
							Description: "Priority of WAN for traffic loadbalancing.",
							Computed:    true,
						},
						"site_id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"static_addressing": schema.SingleNestedAttribute{
							Description: "(optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[MagicTransitSiteWANsStaticAddressingDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Description: "A valid CIDR notation representing an IP range.",
									Computed:    true,
								},
								"gateway_address": schema.StringAttribute{
									Description: "A valid IPv4 address.",
									Computed:    true,
								},
								"secondary_address": schema.StringAttribute{
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

func (d *MagicTransitSiteWANsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MagicTransitSiteWANsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
