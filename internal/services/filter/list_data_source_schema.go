// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*FiltersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "A case-insensitive string to find in the description.",
				Optional:    true,
			},
			"expression": schema.StringAttribute{
				Description: "A case-insensitive string to find in the expression.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique identifier of the filter.",
				Optional:    true,
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the filter is currently paused.",
				Optional:    true,
			},
			"ref": schema.StringAttribute{
				Description: "The filter ref (a short reference tag) to search for. Must be an exact match.",
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
				CustomType:  customfield.NewNestedObjectListType[FiltersResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the filter.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An informative summary of the filter.",
							Computed:    true,
						},
						"expression": schema.StringAttribute{
							Description: "The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).",
							Computed:    true,
						},
						"paused": schema.BoolAttribute{
							Description: "When true, indicates that the filter is currently paused.",
							Computed:    true,
						},
						"ref": schema.StringAttribute{
							Description: "A short reference tag. Allows you to select related filters.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *FiltersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *FiltersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
