// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &RateLimitsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &RateLimitsDataSource{}

func (r RateLimitsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"page": schema.Float64Attribute{
				Description: "The page number of paginated results.",
				Computed:    true,
				Optional:    true,
			},
			"per_page": schema.Float64Attribute{
				Description: "The maximum number of results per page. You can only set the value to `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.",
				Computed:    true,
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the rate limit.",
							Computed:    true,
						},
						"bypass": schema.ListNestedAttribute{
							Description: "Criteria specifying when the current rate limit should be bypassed. You can specify that the rate limit should not apply to one or more URLs.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("url"),
										},
									},
									"value": schema.StringAttribute{
										Description: "The URL to bypass.",
										Computed:    true,
									},
								},
							},
						},
						"description": schema.StringAttribute{
							Description: "An informative summary of the rate limit. This value is sanitized and any tags will be removed.",
							Computed:    true,
						},
						"disabled": schema.BoolAttribute{
							Description: "When true, indicates that the rate limit is currently disabled.",
							Computed:    true,
						},
						"period": schema.Float64Attribute{
							Description: "The time in seconds (an integer value) to count matching traffic. If the count exceeds the configured threshold within this period, Cloudflare will perform the configured action.",
							Computed:    true,
						},
						"threshold": schema.Float64Attribute{
							Description: "The threshold that will trigger the configured mitigation action. Configure this value along with the `period` property to establish a threshold per period.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *RateLimitsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *RateLimitsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
