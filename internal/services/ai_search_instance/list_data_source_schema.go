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
				Optional: true,
			},
			"search": schema.StringAttribute{
				Description: "Search by id",
				Optional:    true,
			},
			"order_by": schema.StringAttribute{
				Description: "Order By Column Name\nAvailable values: \"created_at\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("created_at"),
				},
			},
			"order_by_direction": schema.StringAttribute{
				Description: "Order By Direction\nAvailable values: \"asc\", \"desc\".",
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
							Description: "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"modified_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"ai_gateway_id": schema.StringAttribute{
							Computed: true,
						},
						"aisearch_model": schema.StringAttribute{
							Description: `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/zai-org/glm-4.7-flash", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "@cf/google/gemma-3-12b-it", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"@cf/meta/llama-3.3-70b-instruct-fp8-fast",
									"@cf/zai-org/glm-4.7-flash",
									"@cf/meta/llama-3.1-8b-instruct-fast",
									"@cf/meta/llama-3.1-8b-instruct-fp8",
									"@cf/meta/llama-4-scout-17b-16e-instruct",
									"@cf/qwen/qwen3-30b-a3b-fp8",
									"@cf/deepseek-ai/deepseek-r1-distill-qwen-32b",
									"@cf/moonshotai/kimi-k2-instruct",
									"@cf/google/gemma-3-12b-it",
									"anthropic/claude-3-7-sonnet",
									"anthropic/claude-sonnet-4",
									"anthropic/claude-opus-4",
									"anthropic/claude-3-5-haiku",
									"cerebras/qwen-3-235b-a22b-instruct",
									"cerebras/qwen-3-235b-a22b-thinking",
									"cerebras/llama-3.3-70b",
									"cerebras/llama-4-maverick-17b-128e-instruct",
									"cerebras/llama-4-scout-17b-16e-instruct",
									"cerebras/gpt-oss-120b",
									"google-ai-studio/gemini-2.5-flash",
									"google-ai-studio/gemini-2.5-pro",
									"grok/grok-4",
									"groq/llama-3.3-70b-versatile",
									"groq/llama-3.1-8b-instant",
									"openai/gpt-5",
									"openai/gpt-5-mini",
									"openai/gpt-5-nano",
									"",
								),
							},
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
						"chunk_overlap": schema.Int64Attribute{
							Computed: true,
							Validators: []validator.Int64{
								int64validator.Between(0, 30),
							},
						},
						"chunk_size": schema.Int64Attribute{
							Computed: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(64),
							},
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
							Description: `Available values: "@cf/qwen/qwen3-embedding-0.6b", "@cf/baai/bge-m3", "@cf/baai/bge-large-en-v1.5", "@cf/google/embeddinggemma-300m", "google-ai-studio/gemini-embedding-001", "google-ai-studio/gemini-embedding-2-preview", "openai/text-embedding-3-small", "openai/text-embedding-3-large", "".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"@cf/qwen/qwen3-embedding-0.6b",
									"@cf/baai/bge-m3",
									"@cf/baai/bge-large-en-v1.5",
									"@cf/google/embeddinggemma-300m",
									"google-ai-studio/gemini-embedding-001",
									"google-ai-studio/gemini-embedding-2-preview",
									"openai/text-embedding-3-small",
									"openai/text-embedding-3-large",
									"",
								),
							},
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
							Description:        "Deprecated — use index_method instead.",
							Computed:           true,
							DeprecationMessage: "This attribute is deprecated.",
						},
						"index_method": schema.SingleNestedAttribute{
							Description: "Controls which storage backends are used during indexing. Defaults to vector-only.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[AISearchInstancesIndexMethodDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"keyword": schema.BoolAttribute{
									Description: "Enable keyword (BM25) storage backend.",
									Computed:    true,
								},
								"vector": schema.BoolAttribute{
									Description: "Enable vector (embedding) storage backend.",
									Computed:    true,
								},
							},
						},
						"indexing_options": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AISearchInstancesIndexingOptionsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"keyword_tokenizer": schema.StringAttribute{
									Description: "Tokenizer used for keyword search indexing. porter provides word-level tokenization with Porter stemming (good for natural language queries). trigram enables character-level substring matching (good for partial matches, code, identifiers). Changing this triggers a full re-index. Defaults to porter.\nAvailable values: \"porter\", \"trigram\".",
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
						"max_num_results": schema.Int64Attribute{
							Computed: true,
							Validators: []validator.Int64{
								int64validator.Between(1, 50),
							},
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
											Description: "Disable chat completions endpoint for this public endpoint",
											Computed:    true,
										},
									},
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
											Description: "Disable MCP endpoint for this public endpoint",
											Computed:    true,
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
											Description: "Disable search endpoint for this public endpoint",
											Computed:    true,
										},
									},
								},
							},
						},
						"reranking": schema.BoolAttribute{
							Computed: true,
						},
						"reranking_model": schema.StringAttribute{
							Description: `Available values: "@cf/baai/bge-reranker-base", "".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("@cf/baai/bge-reranker-base", ""),
							},
						},
						"retrieval_options": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AISearchInstancesRetrievalOptionsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"boost_by": schema.ListNestedAttribute{
									Description: "Metadata fields to boost search results by. Each entry specifies a metadata field and an optional direction. Direction defaults to 'asc' for numeric fields and 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined custom_metadata field.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[AISearchInstancesRetrievalOptionsBoostByDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"field": schema.StringAttribute{
												Description: "Metadata field name to boost by. Use 'timestamp' for document freshness, or any custom_metadata field. Numeric and datetime fields support asc/desc directions; text/boolean fields support exists/not_exists.",
												Computed:    true,
											},
											"direction": schema.StringAttribute{
												Description: "Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps). 'asc' = lower values rank higher. 'exists' = boost chunks that have the field. 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc' for numeric/datetime fields, 'exists' for text/boolean fields.\nAvailable values: \"asc\", \"desc\", \"exists\", \"not_exists\".",
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
									Description: "Controls which documents are candidates for BM25 scoring. 'and' restricts candidates to documents containing all query terms; 'or' includes any document containing at least one term, ranked by BM25 relevance. Defaults to 'and'. Legacy values 'exact_match' and 'fuzzy_match' are accepted and map to 'and' and 'or' respectively.\nAvailable values: \"and\", \"or\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("and", "or"),
									},
								},
							},
						},
						"rewrite_model": schema.StringAttribute{
							Description: `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/zai-org/glm-4.7-flash", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "@cf/google/gemma-3-12b-it", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"@cf/meta/llama-3.3-70b-instruct-fp8-fast",
									"@cf/zai-org/glm-4.7-flash",
									"@cf/meta/llama-3.1-8b-instruct-fast",
									"@cf/meta/llama-3.1-8b-instruct-fp8",
									"@cf/meta/llama-4-scout-17b-16e-instruct",
									"@cf/qwen/qwen3-30b-a3b-fp8",
									"@cf/deepseek-ai/deepseek-r1-distill-qwen-32b",
									"@cf/moonshotai/kimi-k2-instruct",
									"@cf/google/gemma-3-12b-it",
									"anthropic/claude-3-7-sonnet",
									"anthropic/claude-sonnet-4",
									"anthropic/claude-opus-4",
									"anthropic/claude-3-5-haiku",
									"cerebras/qwen-3-235b-a22b-instruct",
									"cerebras/qwen-3-235b-a22b-thinking",
									"cerebras/llama-3.3-70b",
									"cerebras/llama-4-maverick-17b-128e-instruct",
									"cerebras/llama-4-scout-17b-16e-instruct",
									"cerebras/gpt-oss-120b",
									"google-ai-studio/gemini-2.5-flash",
									"google-ai-studio/gemini-2.5-pro",
									"grok/grok-4",
									"groq/llama-3.3-70b-versatile",
									"groq/llama-3.1-8b-instant",
									"openai/gpt-5",
									"openai/gpt-5-mini",
									"openai/gpt-5-nano",
									"",
								),
							},
						},
						"rewrite_query": schema.BoolAttribute{
							Computed: true,
						},
						"score_threshold": schema.Float64Attribute{
							Computed: true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
						"source": schema.StringAttribute{
							Computed: true,
						},
						"source_params": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AISearchInstancesSourceParamsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"exclude_items": schema.ListAttribute{
									Description: "List of path patterns to exclude. Uses micromatch glob syntax: * matches within a path segment, ** matches across path segments (e.g., /admin/** matches /admin/users and /admin/settings/advanced)",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"include_items": schema.ListAttribute{
									Description: "List of path patterns to include. Uses micromatch glob syntax: * matches within a path segment, ** matches across path segments (e.g., /blog/** matches /blog/post and /blog/2024/post)",
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
										"crawl_options": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[AISearchInstancesSourceParamsWebCrawlerCrawlOptionsDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"depth": schema.Float64Attribute{
													Computed: true,
													Validators: []validator.Float64{
														float64validator.Between(1, 100000),
													},
												},
												"include_external_links": schema.BoolAttribute{
													Computed: true,
												},
												"include_subdomains": schema.BoolAttribute{
													Computed: true,
												},
												"max_age": schema.Float64Attribute{
													Computed: true,
													Validators: []validator.Float64{
														float64validator.Between(0, 604800),
													},
												},
												"source": schema.StringAttribute{
													Description: `Available values: "all", "sitemaps", "links".`,
													Computed:    true,
													Validators: []validator.String{
														stringvalidator.OneOfCaseInsensitive(
															"all",
															"sitemaps",
															"links",
														),
													},
												},
											},
										},
										"parse_options": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[AISearchInstancesSourceParamsWebCrawlerParseOptionsDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"content_selector": schema.ListNestedAttribute{
													Description: "List of path-to-selector mappings for extracting specific content from crawled pages. Each entry pairs a URL glob pattern with a CSS selector. The first matching path wins. Only the matched HTML fragment is stored and indexed.",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectListType[AISearchInstancesSourceParamsWebCrawlerParseOptionsContentSelectorDataSourceModel](ctx),
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"path": schema.StringAttribute{
																Description: "Glob pattern to match against the page URL path. Uses standard glob syntax: * matches within a segment, ** crosses directories.",
																Computed:    true,
															},
															"selector": schema.StringAttribute{
																Description: "CSS selector to extract content from pages matching the path pattern. Supports standard CSS selectors including class, ID, element, and attribute selectors.",
																Computed:    true,
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
													Description: "List of specific sitemap URLs to use for crawling. Only valid when parse_type is 'sitemap'.",
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
											Description: `Available values: "sitemap", "feed-rss", "crawl".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"sitemap",
													"feed-rss",
													"crawl",
												),
											},
										},
										"store_options": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[AISearchInstancesSourceParamsWebCrawlerStoreOptionsDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"storage_id": schema.StringAttribute{
													Computed: true,
												},
												"r2_jurisdiction": schema.StringAttribute{
													Computed: true,
												},
												"storage_type": schema.StringAttribute{
													Description: `Available values: "r2".`,
													Computed:    true,
													Validators: []validator.String{
														stringvalidator.OneOfCaseInsensitive("r2"),
													},
												},
											},
										},
									},
								},
							},
						},
						"status": schema.StringAttribute{
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
