// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*APIShieldOperationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "UUID",
				Computed:    true,
			},
			"operation_id": schema.StringAttribute{
				Description: "UUID",
				Computed:    true,
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"feature": schema.ListAttribute{
				Description: "Add feature(s) to the results. The feature name that is given here corresponds to the resulting feature object. Have a look at the top-level object description for more details on the specific meaning.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"thresholds",
							"parameter_schemas",
							"schema_info",
						),
					),
				},
				ElementType: types.StringType,
			},
			"endpoint": schema.StringAttribute{
				Description: "The endpoint which can contain path parameter templates in curly braces, each will be replaced from left to right with {varN}, starting with {var1}, during insertion. This will further be Cloudflare-normalized upon insertion. See: https://developers.cloudflare.com/rules/normalization/how-it-works/.",
				Computed:    true,
			},
			"host": schema.StringAttribute{
				Description: "RFC3986-compliant host.",
				Computed:    true,
			},
			"last_updated": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"method": schema.StringAttribute{
				Description: "The HTTP method used to access the endpoint.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"GET",
						"POST",
						"HEAD",
						"OPTIONS",
						"PUT",
						"DELETE",
						"CONNECT",
						"PATCH",
						"TRACE",
					),
				},
			},
			"features": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"thresholds": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesThresholdsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"auth_id_tokens": schema.Int64Attribute{
								Description: "The total number of auth-ids seen across this calculation.",
								Computed:    true,
							},
							"data_points": schema.Int64Attribute{
								Description: "The number of data points used for the threshold suggestion calculation.",
								Computed:    true,
							},
							"last_updated": schema.StringAttribute{
								Computed:   true,
								CustomType: timetypes.RFC3339Type{},
							},
							"p50": schema.Int64Attribute{
								Description: "The p50 quantile of requests (in period_seconds).",
								Computed:    true,
							},
							"p90": schema.Int64Attribute{
								Description: "The p90 quantile of requests (in period_seconds).",
								Computed:    true,
							},
							"p99": schema.Int64Attribute{
								Description: "The p99 quantile of requests (in period_seconds).",
								Computed:    true,
							},
							"period_seconds": schema.Int64Attribute{
								Description: "The period over which this threshold is suggested.",
								Computed:    true,
							},
							"requests": schema.Int64Attribute{
								Description: "The estimated number of requests covered by these calculations.",
								Computed:    true,
							},
							"suggested_threshold": schema.Int64Attribute{
								Description: "The suggested threshold in requests done by the same auth_id or period_seconds.",
								Computed:    true,
							},
						},
					},
					"parameter_schemas": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesParameterSchemasDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"last_updated": schema.StringAttribute{
								Computed:   true,
								CustomType: timetypes.RFC3339Type{},
							},
							"parameter_schemas": schema.SingleNestedAttribute{
								Description: "An operation schema object containing a response.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[APIShieldOperationFeaturesParameterSchemasParameterSchemasDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"parameters": schema.ListAttribute{
										Description: "An array containing the learned parameter schemas.",
										Computed:    true,
										CustomType:  customfield.NewListType[jsontypes.Normalized](ctx),
										ElementType: jsontypes.NormalizedType{},
									},
									"responses": schema.StringAttribute{
										Description: "An empty response object. This field is required to yield a valid operation schema.",
										Computed:    true,
										CustomType:  jsontypes.NormalizedType{},
									},
								},
							},
						},
					},
					"api_routing": schema.SingleNestedAttribute{
						Description: "API Routing settings on endpoint.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[APIShieldOperationFeaturesAPIRoutingDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"last_updated": schema.StringAttribute{
								Computed:   true,
								CustomType: timetypes.RFC3339Type{},
							},
							"route": schema.StringAttribute{
								Description: "Target route.",
								Computed:    true,
							},
						},
					},
					"confidence_intervals": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"last_updated": schema.StringAttribute{
								Computed:   true,
								CustomType: timetypes.RFC3339Type{},
							},
							"suggested_threshold": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"confidence_intervals": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"p90": schema.SingleNestedAttribute{
												Description: "Upper and lower bound for percentile estimate",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90DataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"lower": schema.Float64Attribute{
														Description: "Lower bound for percentile estimate",
														Computed:    true,
													},
													"upper": schema.Float64Attribute{
														Description: "Upper bound for percentile estimate",
														Computed:    true,
													},
												},
											},
											"p95": schema.SingleNestedAttribute{
												Description: "Upper and lower bound for percentile estimate",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95DataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"lower": schema.Float64Attribute{
														Description: "Lower bound for percentile estimate",
														Computed:    true,
													},
													"upper": schema.Float64Attribute{
														Description: "Upper bound for percentile estimate",
														Computed:    true,
													},
												},
											},
											"p99": schema.SingleNestedAttribute{
												Description: "Upper and lower bound for percentile estimate",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99DataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"lower": schema.Float64Attribute{
														Description: "Lower bound for percentile estimate",
														Computed:    true,
													},
													"upper": schema.Float64Attribute{
														Description: "Upper bound for percentile estimate",
														Computed:    true,
													},
												},
											},
										},
									},
									"mean": schema.Float64Attribute{
										Description: "Suggested threshold.",
										Computed:    true,
									},
								},
							},
						},
					},
					"schema_info": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesSchemaInfoDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"active_schema": schema.SingleNestedAttribute{
								Description: "Schema active on endpoint.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[APIShieldOperationFeaturesSchemaInfoActiveSchemaDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "UUID",
										Computed:    true,
									},
									"created_at": schema.StringAttribute{
										Computed:   true,
										CustomType: timetypes.RFC3339Type{},
									},
									"is_learned": schema.BoolAttribute{
										Description: "True if schema is Cloudflare-provided.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "Schema file name.",
										Computed:    true,
									},
								},
							},
							"learned_available": schema.BoolAttribute{
								Description: "True if a Cloudflare-provided learned schema is available for this endpoint.",
								Computed:    true,
							},
							"mitigation_action": schema.StringAttribute{
								Description: "Action taken on requests failing validation.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"none",
										"log",
										"block",
									),
								},
							},
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Direction to order results.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"endpoint": schema.StringAttribute{
						Description: "Filter results to only include endpoints containing this pattern.",
						Optional:    true,
					},
					"feature": schema.ListAttribute{
						Description: "Add feature(s) to the results. The feature name that is given here corresponds to the resulting feature object. Have a look at the top-level object description for more details on the specific meaning.",
						Optional:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"thresholds",
									"parameter_schemas",
									"schema_info",
								),
							),
						},
						ElementType: types.StringType,
					},
					"host": schema.ListAttribute{
						Description: "Filter results to only include the specified hosts.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"method": schema.ListAttribute{
						Description: "Filter results to only include the specified HTTP methods.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"order": schema.StringAttribute{
						Description: "Field to order by. When requesting a feature, the feature keys are available for ordering as well, e.g., `thresholds.suggested_threshold`.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"method",
								"host",
								"endpoint",
								"thresholds.$key",
							),
						},
					},
				},
			},
		},
	}
}

func (d *APIShieldOperationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *APIShieldOperationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("operation_id"), path.MatchRoot("filter")),
	}
}
