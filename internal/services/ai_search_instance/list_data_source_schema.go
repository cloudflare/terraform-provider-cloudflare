// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_instance

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*AISearchInstancesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"namespace": schema.StringAttribute{
				Description: "Filter by namespace.",
				Optional:    true,
			},
			"search": schema.StringAttribute{
				Description: "Filter instances whose id contains this string (case-insensitive).",
				Optional:    true,
			},
			"order_by": schema.StringAttribute{
				Description: "Field to order results by.\nAvailable values: \"created_at\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("created_at"),
				},
			},
			"order_by_direction": schema.StringAttribute{
				Description: "Order direction.\nAvailable values: \"asc\", \"desc\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
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
				CustomType:  customfield.NewNestedObjectListType[AISearchInstancesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"ai_gateway_id": schema.StringAttribute{
							Computed: true,
						},
						"aisearch_model": schema.StringAttribute{
							Computed: true,
						},
						"cache": schema.BoolAttribute{
							Computed: true,
						},
						"cache_threshold": schema.StringAttribute{
							Description: `Available values: "super_strict_match", "close_enough", "flexible_friend", "anything_goes".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"super_strict_match",
									"close_enough",
									"flexible_friend",
									"anything_goes",
								),
							},
						},
						"cache_ttl": schema.Float64Attribute{
							Description: "Available values: 600, 1800, 3600, 7200, 21600, 43200, 86400, 172800, 259200, 518400.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.OneOf(
									600,
									1800,
									3600,
									7200,
									21600,
									43200,
									86400,
									172800,
									259200,
									518400,
								),
							},
						},
						"chunk": schema.BoolAttribute{
							Computed: true,
						},
						"chunk_overlap": schema.Float64Attribute{
							Computed: true,
						},
						"chunk_size": schema.Float64Attribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"created_by": schema.StringAttribute{
							Computed: true,
						},
						"custom_metadata": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[AISearchInstancesCustomMetadataDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"data_type": schema.StringAttribute{
										Description: `Available values: "text", "number", "boolean", "datetime".`,
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"text",
												"number",
												"boolean",
												"datetime",
											),
										},
									},
									"field_name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"embedding_model": schema.StringAttribute{
							Computed: true,
						},
						"enable": schema.BoolAttribute{
							Computed: true,
						},
						"engine_version": schema.Float64Attribute{
							Computed: true,
						},
						"fusion_method": schema.StringAttribute{
							Description: `Available values: "max", "rrf".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("max", "rrf"),
							},
						},
						"hybrid_search_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"index_method": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AISearchInstancesIndexMethodDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"keyword": schema.BoolAttribute{
									Computed: true,
								},
								"vector": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
						"indexing_options": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AISearchInstancesIndexingOptionsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"keyword_tokenizer": schema.StringAttribute{
									Description: `Available values: "porter", "trigram".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("porter", "trigram"),
									},
								},
							},
						},
						"last_activity": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"max_num_results": schema.Float64Attribute{
							Computed: true,
						},
						"metadata": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AISearchInstancesMetadataDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"created_from_aisearch_wizard": schema.BoolAttribute{
									Computed: true,
								},
								"worker_domain": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"modified_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"modified_by": schema.StringAttribute{
							Computed: true,
						},
						"namespace": schema.StringAttribute{
							Computed: true,
						},
						"paused": schema.BoolAttribute{
							Computed: true,
						},
						"public_endpoint_id": schema.StringAttribute{
							Computed: true,
						},
						"public_endpoint_params": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AISearchInstancesPublicEndpointParamsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"authorized_hosts": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"chat_completions_endpoint": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AISearchInstancesPublicEndpointParamsChatCompletionsEndpointDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"disabled": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"custom_domains": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"default_domain_enabled": schema.BoolAttribute{
									Computed: true,
								},
								"enabled": schema.BoolAttribute{
									Computed: true,
								},
								"mcp": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AISearchInstancesPublicEndpointParamsMcpDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Computed: true,
										},
										"disabled": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"rate_limit": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AISearchInstancesPublicEndpointParamsRateLimitDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"period_ms": schema.Int64Attribute{
											Computed: true,
											Validators: []validator.Int64{
												int64validator.Between(60000, 3600000),
											},
										},
										"requests": schema.Int64Attribute{
											Computed: true,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},
										"technique": schema.StringAttribute{
											Description: `Available values: "fixed", "sliding".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("fixed", "sliding"),
											},
										},
									},
								},
								"search_endpoint": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AISearchInstancesPublicEndpointParamsSearchEndpointDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"disabled": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
							},
						},
						"reranking": schema.BoolAttribute{
							Computed: true,
						},
						"reranking_model": schema.StringAttribute{
							Computed: true,
						},
						"retrieval_options": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AISearchInstancesRetrievalOptionsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"boost_by": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[AISearchInstancesRetrievalOptionsBoostByDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"field": schema.StringAttribute{
												Computed: true,
											},
											"data_type": schema.StringAttribute{
												Description: `Available values: "number", "datetime", "text", "boolean".`,
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"number",
														"datetime",
														"text",
														"boolean",
													),
												},
											},
											"direction": schema.StringAttribute{
												Description: `Available values: "asc", "desc", "exists", "not_exists".`,
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"asc",
														"desc",
														"exists",
														"not_exists",
													),
												},
											},
										},
									},
								},
								"keyword_match_mode": schema.StringAttribute{
									Description: `Available values: "and", "or".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("and", "or"),
									},
								},
							},
						},
						"rewrite_model": schema.StringAttribute{
							Computed: true,
						},
						"rewrite_query": schema.BoolAttribute{
							Computed: true,
						},
						"score_threshold": schema.Float64Attribute{
							Computed: true,
						},
						"source": schema.StringAttribute{
							Computed: true,
						},
						"source_params": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AISearchInstancesSourceParamsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"exclude_items": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"include_items": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"prefix": schema.StringAttribute{
									Computed: true,
								},
								"r2_jurisdiction": schema.StringAttribute{
									Computed: true,
								},
								"web_crawler": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AISearchInstancesSourceParamsWebCrawlerDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"parse_options": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[AISearchInstancesSourceParamsWebCrawlerParseOptionsDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"content_selector": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[AISearchInstancesSourceParamsWebCrawlerParseOptionsContentSelectorDataSourceModel](ctx),
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"path": schema.StringAttribute{
																Computed: true,
															},
															"selector": schema.StringAttribute{
																Computed: true,
															},
														},
													},
												},
												"include_headers": schema.MapAttribute{
													Computed:    true,
													CustomType:  customfield.NewMapType[types.String](ctx),
													ElementType: types.StringType,
												},
												"include_images": schema.BoolAttribute{
													Computed: true,
												},
												"specific_sitemaps": schema.ListAttribute{
													Computed:    true,
													CustomType:  customfield.NewListType[types.String](ctx),
													ElementType: types.StringType,
												},
												"use_browser_rendering": schema.BoolAttribute{
													Computed: true,
												},
											},
										},
										"parse_type": schema.StringAttribute{
											Description: `Available values: "sitemap", "crawl".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("sitemap", "crawl"),
											},
										},
									},
								},
							},
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
						"summarization": schema.BoolAttribute{
							Computed: true,
						},
						"summarization_model": schema.StringAttribute{
							Computed: true,
						},
						"sync_interval": schema.Float64Attribute{
							Description: "Available values: 900, 1800, 3600, 7200, 14400, 21600, 43200, 86400.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.OneOf(
									900,
									1800,
									3600,
									7200,
									14400,
									21600,
									43200,
									86400,
								),
							},
						},
						"system_prompt_aisearch": schema.StringAttribute{
							Computed: true,
						},
						"system_prompt_index_summarization": schema.StringAttribute{
							Computed: true,
						},
						"system_prompt_rewrite_query": schema.StringAttribute{
							Computed: true,
						},
						"token_id": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Description: `Available values: "r2", "web-crawler".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("r2", "web-crawler"),
							},
						},
					},
				},
			},
		},
	}
}

func (d *AISearchInstancesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AISearchInstancesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
