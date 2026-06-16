// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_cf1_site

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitCf1SiteDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Magic Transit Read",
				"Magic Transit Write",
				"Magic WAN Read",
				"Magic WAN Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"cf1_site_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "A human-provided description of the CF1 Site.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "A human-provided name describing the CF1 Site that should be unique within the account.",
				Computed:    true,
			},
			"location": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicTransitCf1SiteLocationDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"lat": schema.Float64Attribute{
						Description: "Latitude of the CF1 Site.",
						Computed:    true,
					},
					"long": schema.Float64Attribute{
						Description: "Longitude of the CF1 Site.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of nearest town, city, or village.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *MagicTransitCf1SiteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicTransitCf1SiteDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
