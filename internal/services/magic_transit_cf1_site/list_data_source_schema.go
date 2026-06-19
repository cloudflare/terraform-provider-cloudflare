// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_cf1_site

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitCf1SitesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
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
			"account_id": schema.StringAttribute{
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
				CustomType:  customfield.NewNestedObjectListType[MagicTransitCf1SitesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "A human-provided name describing the CF1 Site that should be unique within the account.",
							Computed:    true,
						},
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "A human-provided description of the CF1 Site.",
							Computed:    true,
						},
						"location": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[MagicTransitCf1SitesLocationDataSourceModel](ctx),
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
						"modified_on": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *MagicTransitCf1SitesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MagicTransitCf1SitesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
