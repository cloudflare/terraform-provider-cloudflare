// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_wan

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitSiteWANDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"wan_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"site_id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
				Optional:    true,
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
			"priority": schema.Int64Attribute{
				Description: "Priority of WAN for traffic loadbalancing.",
				Computed:    true,
			},
			"vlan_tag": schema.Int64Attribute{
				Description: "VLAN port number.",
				Computed:    true,
			},
			"static_addressing": schema.SingleNestedAttribute{
				Description: "(optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[MagicTransitSiteWANStaticAddressingDataSourceModel](ctx),
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"site_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *MagicTransitSiteWANDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicTransitSiteWANDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(
			path.MatchRoot("account_id"),
			path.MatchRoot("site_id"),
			path.MatchRoot("wan_id"),
		),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("site_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("wan_id")),
	}
}
