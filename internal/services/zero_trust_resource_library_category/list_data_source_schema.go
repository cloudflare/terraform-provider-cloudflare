// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_category

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustResourceLibraryCategoriesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"limit": schema.Int64Attribute{
				Description: "Number of results to return. Defaults to 25; values are clamped to 1–250.",
				Computed:    true,
				Optional:    true,
			},
			"offset": schema.Int64Attribute{
				Description: "Number of results to skip. Defaults to 0; negative values are clamped to 0.",
				Computed:    true,
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustResourceLibraryCategoriesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description: "Numeric identifier for an application category.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"created_at": schema.StringAttribute{
							Description: "Returns the category creation time.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Returns the category description.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Returns the category name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustResourceLibraryCategoriesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustResourceLibraryCategoriesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
