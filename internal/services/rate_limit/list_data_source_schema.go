// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*RateLimitsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		DeprecationMessage: "Rate limiting API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version for full details.",
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Defines an identifier.",
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
				CustomType:  customfield.NewNestedObjectListType[RateLimitsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the rate limit.",
							Computed:    true,
						},
						"action": schema.SingleNestedAttribute{
							Description: "The action to perform when the threshold of matched traffic within the configured period is exceeded.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RateLimitsActionDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"mode": schema.StringAttribute{
									Description: "The action to perform.\nAvailable values: \"simulate\", \"ban\", \"challenge\", \"js_challenge\", \"managed_challenge\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"simulate",
											"ban",
											"challenge",
											"js_challenge",
											"managed_challenge",
										),
									},
								},
								"response": schema.SingleNestedAttribute{
									Description: "A custom content type and reponse to return when the threshold is exceeded. The custom response configured in this object will override the custom error for the zone. This object is optional.\nNotes: If you omit this object, Cloudflare will use the default HTML error page. If \"mode\" is \"challenge\", \"managed_challenge\", or \"js_challenge\", Cloudflare will use the zone challenge pages and you should not provide the \"response\" object.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[RateLimitsActionResponseDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"body": schema.StringAttribute{
											Description: "The response body to return. The value must conform to the configured content type.",
											Computed:    true,
										},
										"content_type": schema.StringAttribute{
											Description: "The content type of the body. Must be one of the following: `text/plain`, `text/xml`, or `application/json`.",
											Computed:    true,
										},
									},
								},
								"timeout": schema.Float64Attribute{
									Description: "The time in seconds during which Cloudflare will perform the mitigation action. Must be an integer value greater than or equal to the period.\nNotes: If \"mode\" is \"challenge\", \"managed_challenge\", or \"js_challenge\", Cloudflare will use the zone's Challenge Passage time and you should not provide this value.",
									Computed:    true,
									Validators: []validator.Float64{
										float64validator.Between(1, 86400),
									},
								},
							},
						},
						"bypass": schema.ListNestedAttribute{
							Description: "Criteria specifying when the current rate limit should be bypassed. You can specify that the rate limit should not apply to one or more URLs.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[RateLimitsBypassDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description: `Available values: "url".`,
										Computed:    true,
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
						"match": schema.SingleNestedAttribute{
							Description: "Determines which traffic the rate limit counts towards the threshold.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RateLimitsMatchDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"headers": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[RateLimitsMatchHeadersDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the response header to match.",
												Computed:    true,
											},
											"op": schema.StringAttribute{
												Description: "The operator used when matching: `eq` means \"equal\" and `ne` means \"not equal\".\nAvailable values: \"eq\", \"ne\".",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("eq", "ne"),
												},
											},
											"value": schema.StringAttribute{
												Description: "The value of the response header, which must match exactly.",
												Computed:    true,
											},
										},
									},
								},
								"request": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[RateLimitsMatchRequestDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"methods": schema.ListAttribute{
											Description: "The HTTP methods to match. You can specify a subset (for example, `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when creating a rate limit.",
											Computed:    true,
											Validators: []validator.List{
												listvalidator.ValueStringsAre(
													stringvalidator.OneOfCaseInsensitive(
														"GET",
														"POST",
														"PUT",
														"DELETE",
														"PATCH",
														"HEAD",
														"_ALL_",
													),
												),
											},
											CustomType:  customfield.NewListType[types.String](ctx),
											ElementType: types.StringType,
										},
										"schemes": schema.ListAttribute{
											Description: "The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is optional.",
											Computed:    true,
											CustomType:  customfield.NewListType[types.String](ctx),
											ElementType: types.StringType,
										},
										"url": schema.StringAttribute{
											Description: "The URL pattern to match, composed of a host and a path such as `example.org/path*`. Normalization is applied before the pattern is matched. `*` wildcards are expanded to match applicable traffic. Query strings are not matched. Set the value to `*` to match all traffic to your zone.",
											Computed:    true,
										},
									},
								},
								"response": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[RateLimitsMatchResponseDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"origin_traffic": schema.BoolAttribute{
											Description: "When true, only the uncached traffic served from your origin servers will count towards rate limiting. In this case, any cached traffic served by Cloudflare will not count towards rate limiting. This field is optional.\nNotes: This field is deprecated. Instead, use response headers and set \"origin_traffic\" to \"false\" to avoid legacy behaviour interacting with the \"response_headers\" property.",
											Computed:    true,
										},
									},
								},
							},
						},
						"period": schema.Float64Attribute{
							Description: "The time in seconds (an integer value) to count matching traffic. If the count exceeds the configured threshold within this period, Cloudflare will perform the configured action.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(10, 86400),
							},
						},
						"threshold": schema.Float64Attribute{
							Description: "The threshold that will trigger the configured mitigation action. Configure this value along with the `period` property to establish a threshold per period.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.AtLeast(1),
							},
						},
					},
				},
			},
		},
	}
}

func (d *RateLimitsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *RateLimitsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
