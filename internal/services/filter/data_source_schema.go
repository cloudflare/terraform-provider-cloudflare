// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*FilterDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		DeprecationMessage: "The Filters API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique identifier of the filter.",
				Computed:    true,
			},
			"filter_id": schema.StringAttribute{
				Description: "The unique identifier of the filter.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Defines an identifier.",
				Required:    true,
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
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
					"paused": schema.BoolAttribute{
						Description: "When true, indicates that the filter is currently paused.",
						Optional:    true,
					},
					"ref": schema.StringAttribute{
						Description: "The filter ref (a short reference tag) to search for. Must be an exact match.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *FilterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *FilterDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter_id"), path.MatchRoot("filter")),
	}
}
