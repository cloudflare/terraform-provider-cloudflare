// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &FiltersDataSource{}
var _ datasource.DataSourceWithValidateConfig = &FiltersDataSource{}

func (r FiltersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique identifier of the filter.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A case-insensitive string to find in the description.",
				Optional:    true,
			},
			"expression": schema.StringAttribute{
				Description: "A case-insensitive string to find in the expression.",
				Optional:    true,
			},
			"page": schema.Float64Attribute{
				Description: "Page number of paginated results.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the filter is currently paused.",
				Optional:    true,
			},
			"per_page": schema.Float64Attribute{
				Description: "Number of filters per page.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(5, 100),
				},
			},
			"ref": schema.StringAttribute{
				Description: "The filter ref (a short reference tag) to search for. Must be an exact match.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the filter.",
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
						"description": schema.StringAttribute{
							Description: "An informative summary of the filter.",
							Computed:    true,
							Optional:    true,
						},
						"ref": schema.StringAttribute{
							Description: "A short reference tag. Allows you to select related filters.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (r *FiltersDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *FiltersDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
