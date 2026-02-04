// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_instance

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*AISearchInstanceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Use your AI Search ID.",
				Computed:    true,
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"account_tag": schema.StringAttribute{
				Computed: true,
			},
			"ai_gateway_id": schema.StringAttribute{
				Computed: true,
			},
			"aisearch_model": schema.StringAttribute{
				Description: `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"@cf/meta/llama-3.3-70b-instruct-fp8-fast",
						"@cf/meta/llama-3.1-8b-instruct-fast",
						"@cf/meta/llama-3.1-8b-instruct-fp8",
						"@cf/meta/llama-4-scout-17b-16e-instruct",
						"@cf/qwen/qwen3-30b-a3b-fp8",
						"@cf/deepseek-ai/deepseek-r1-distill-qwen-32b",
						"@cf/moonshotai/kimi-k2-instruct",
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
			"chunk": schema.BoolAttribute{
				Computed: true,
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
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"created_by": schema.StringAttribute{
				Computed: true,
			},
			"embedding_model": schema.StringAttribute{
				Description: `Available values: "@cf/qwen/qwen3-embedding-0.6b", "@cf/baai/bge-m3", "@cf/baai/bge-large-en-v1.5", "@cf/google/embeddinggemma-300m", "google-ai-studio/gemini-embedding-001", "openai/text-embedding-3-small", "openai/text-embedding-3-large", "".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"@cf/qwen/qwen3-embedding-0.6b",
						"@cf/baai/bge-m3",
						"@cf/baai/bge-large-en-v1.5",
						"@cf/google/embeddinggemma-300m",
						"google-ai-studio/gemini-embedding-001",
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
			"hybrid_search_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"internal_id": schema.StringAttribute{
				Computed: true,
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
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_by": schema.StringAttribute{
				Computed: true,
			},
			"paused": schema.BoolAttribute{
				Computed: true,
			},
			"public_endpoint_id": schema.StringAttribute{
				Computed: true,
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
			"rewrite_model": schema.StringAttribute{
				Description: `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"@cf/meta/llama-3.3-70b-instruct-fp8-fast",
						"@cf/meta/llama-3.1-8b-instruct-fast",
						"@cf/meta/llama-3.1-8b-instruct-fp8",
						"@cf/meta/llama-4-scout-17b-16e-instruct",
						"@cf/qwen/qwen3-30b-a3b-fp8",
						"@cf/deepseek-ai/deepseek-r1-distill-qwen-32b",
						"@cf/moonshotai/kimi-k2-instruct",
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
			"status": schema.StringAttribute{
				Computed: true,
			},
			"summarization": schema.BoolAttribute{
				Computed: true,
			},
			"summarization_model": schema.StringAttribute{
				Description: `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"@cf/meta/llama-3.3-70b-instruct-fp8-fast",
						"@cf/meta/llama-3.1-8b-instruct-fast",
						"@cf/meta/llama-3.1-8b-instruct-fp8",
						"@cf/meta/llama-4-scout-17b-16e-instruct",
						"@cf/qwen/qwen3-30b-a3b-fp8",
						"@cf/deepseek-ai/deepseek-r1-distill-qwen-32b",
						"@cf/moonshotai/kimi-k2-instruct",
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
			"vectorize_active_namespace": schema.StringAttribute{
				Computed: true,
			},
			"vectorize_name": schema.StringAttribute{
				Computed: true,
			},
			"custom_metadata": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[AISearchInstanceCustomMetadataDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"data_type": schema.StringAttribute{
							Description: `Available values: "text", "number", "boolean".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"text",
									"number",
									"boolean",
								),
							},
						},
						"field_name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"metadata": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AISearchInstanceMetadataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"created_from_aisearch_wizard": schema.BoolAttribute{
						Computed: true,
					},
					"worker_domain": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"public_endpoint_params": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AISearchInstancePublicEndpointParamsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"authorized_hosts": schema.ListAttribute{
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"chat_completions_endpoint": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[AISearchInstancePublicEndpointParamsChatCompletionsEndpointDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AISearchInstancePublicEndpointParamsMcpDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"disabled": schema.BoolAttribute{
								Description: "Disable MCP endpoint for this public endpoint",
								Computed:    true,
							},
						},
					},
					"rate_limit": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[AISearchInstancePublicEndpointParamsRateLimitDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AISearchInstancePublicEndpointParamsSearchEndpointDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"disabled": schema.BoolAttribute{
								Description: "Disable search endpoint for this public endpoint",
								Computed:    true,
							},
						},
					},
				},
			},
			"source_params": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AISearchInstanceSourceParamsDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AISearchInstanceSourceParamsWebCrawlerDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"parse_options": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[AISearchInstanceSourceParamsWebCrawlerParseOptionsDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
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
								Description: `Available values: "sitemap", "feed-rss".`,
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("sitemap", "feed-rss"),
								},
							},
							"store_options": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[AISearchInstanceSourceParamsWebCrawlerStoreOptionsDataSourceModel](ctx),
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"search": schema.StringAttribute{
						Description: "Search by id",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *AISearchInstanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AISearchInstanceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("filter")),
	}
}
