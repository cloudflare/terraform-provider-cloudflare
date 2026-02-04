// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_instance

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*AISearchInstanceResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Use your AI Search ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"source": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description: `Available values: "r2", "web-crawler".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("r2", "web-crawler"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"ai_gateway_id": schema.StringAttribute{
				Optional: true,
			},
			"aisearch_model": schema.StringAttribute{
				Description: `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
				Optional:    true,
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
			"embedding_model": schema.StringAttribute{
				Description: `Available values: "@cf/qwen/qwen3-embedding-0.6b", "@cf/baai/bge-m3", "@cf/baai/bge-large-en-v1.5", "@cf/google/embeddinggemma-300m", "google-ai-studio/gemini-embedding-001", "openai/text-embedding-3-small", "openai/text-embedding-3-large", "".`,
				Optional:    true,
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
			"reranking_model": schema.StringAttribute{
				Description: `Available values: "@cf/baai/bge-reranker-base", "".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("@cf/baai/bge-reranker-base", ""),
				},
			},
			"rewrite_model": schema.StringAttribute{
				Description: `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
				Optional:    true,
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
			"summarization_model": schema.StringAttribute{
				Description: `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
				Optional:    true,
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
				Optional: true,
			},
			"system_prompt_index_summarization": schema.StringAttribute{
				Optional: true,
			},
			"system_prompt_rewrite_query": schema.StringAttribute{
				Optional: true,
			},
			"token_id": schema.StringAttribute{
				Optional: true,
			},
			"custom_metadata": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"data_type": schema.StringAttribute{
							Description: `Available values: "text", "number", "boolean".`,
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"text",
									"number",
									"boolean",
								),
							},
						},
						"field_name": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"metadata": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"created_from_aisearch_wizard": schema.BoolAttribute{
						Optional: true,
					},
					"worker_domain": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			"cache": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(true),
			},
			"cache_threshold": schema.StringAttribute{
				Description: `Available values: "super_strict_match", "close_enough", "flexible_friend", "anything_goes".`,
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"super_strict_match",
						"close_enough",
						"flexible_friend",
						"anything_goes",
					),
				},
				Default: stringdefault.StaticString("close_enough"),
			},
			"chunk": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(true),
			},
			"chunk_overlap": schema.Int64Attribute{
				Computed: true,
				Optional: true,
				Validators: []validator.Int64{
					int64validator.Between(0, 30),
				},
				Default: int64default.StaticInt64(10),
			},
			"chunk_size": schema.Int64Attribute{
				Computed: true,
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(64),
				},
				Default: int64default.StaticInt64(256),
			},
			"hybrid_search_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"max_num_results": schema.Int64Attribute{
				Computed: true,
				Optional: true,
				Validators: []validator.Int64{
					int64validator.Between(1, 50),
				},
				Default: int64default.StaticInt64(10),
			},
			"paused": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"reranking": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"rewrite_query": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"score_threshold": schema.Float64Attribute{
				Computed: true,
				Optional: true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
				Default: float64default.StaticFloat64(0.4),
			},
			"summarization": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"public_endpoint_params": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[AISearchInstancePublicEndpointParamsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"authorized_hosts": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"chat_completions_endpoint": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[AISearchInstancePublicEndpointParamsChatCompletionsEndpointModel](ctx),
						Attributes: map[string]schema.Attribute{
							"disabled": schema.BoolAttribute{
								Description: "Disable chat completions endpoint for this public endpoint",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
						},
					},
					"enabled": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						Default:  booldefault.StaticBool(false),
					},
					"mcp": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[AISearchInstancePublicEndpointParamsMcpModel](ctx),
						Attributes: map[string]schema.Attribute{
							"disabled": schema.BoolAttribute{
								Description: "Disable MCP endpoint for this public endpoint",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
						},
					},
					"rate_limit": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"period_ms": schema.Int64Attribute{
								Optional: true,
								Validators: []validator.Int64{
									int64validator.Between(60000, 3600000),
								},
							},
							"requests": schema.Int64Attribute{
								Optional: true,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
							"technique": schema.StringAttribute{
								Description: `Available values: "fixed", "sliding".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("fixed", "sliding"),
								},
							},
						},
					},
					"search_endpoint": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[AISearchInstancePublicEndpointParamsSearchEndpointModel](ctx),
						Attributes: map[string]schema.Attribute{
							"disabled": schema.BoolAttribute{
								Description: "Disable search endpoint for this public endpoint",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
						},
					},
				},
			},
			"source_params": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[AISearchInstanceSourceParamsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"exclude_items": schema.ListAttribute{
						Description: "List of path patterns to exclude. Uses micromatch glob syntax: * matches within a path segment, ** matches across path segments (e.g., /admin/** matches /admin/users and /admin/settings/advanced)",
						Optional:    true,
						ElementType: types.StringType,
					},
					"include_items": schema.ListAttribute{
						Description: "List of path patterns to include. Uses micromatch glob syntax: * matches within a path segment, ** matches across path segments (e.g., /blog/** matches /blog/post and /blog/2024/post)",
						Optional:    true,
						ElementType: types.StringType,
					},
					"prefix": schema.StringAttribute{
						Optional: true,
					},
					"r2_jurisdiction": schema.StringAttribute{
						Computed: true,
						Optional: true,
						Default:  stringdefault.StaticString("default"),
					},
					"web_crawler": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[AISearchInstanceSourceParamsWebCrawlerModel](ctx),
						Attributes: map[string]schema.Attribute{
							"parse_options": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"include_headers": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"include_images": schema.BoolAttribute{
										Computed: true,
										Optional: true,
										Default:  booldefault.StaticBool(false),
									},
									"specific_sitemaps": schema.ListAttribute{
										Description: "List of specific sitemap URLs to use for crawling. Only valid when parse_type is 'sitemap'.",
										Optional:    true,
										ElementType: types.StringType,
									},
									"use_browser_rendering": schema.BoolAttribute{
										Computed: true,
										Optional: true,
										Default:  booldefault.StaticBool(false),
									},
								},
							},
							"parse_type": schema.StringAttribute{
								Description: `Available values: "sitemap", "feed-rss".`,
								Computed:    true,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("sitemap", "feed-rss"),
								},
								Default: stringdefault.StaticString("sitemap"),
							},
							"store_options": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"storage_id": schema.StringAttribute{
										Required: true,
									},
									"r2_jurisdiction": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Default:  stringdefault.StaticString("default"),
									},
									"storage_type": schema.StringAttribute{
										Description: `Available values: "r2".`,
										Optional:    true,
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
			"account_tag": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"created_by": schema.StringAttribute{
				Computed: true,
			},
			"enable": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"engine_version": schema.Float64Attribute{
				Computed: true,
				Default:  float64default.StaticFloat64(1),
			},
			"internal_id": schema.StringAttribute{
				Computed: true,
			},
			"last_activity": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_by": schema.StringAttribute{
				Computed: true,
			},
			"public_endpoint_id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
				Default:  stringdefault.StaticString("waiting"),
			},
			"vectorize_active_namespace": schema.StringAttribute{
				Computed: true,
			},
			"vectorize_name": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *AISearchInstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AISearchInstanceResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
