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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*AISearchInstanceResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"source": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description: `Available values: "r2", "web-crawler".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("r2", "web-crawler"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"ai_gateway_id": schema.StringAttribute{
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"aisearch_model": schema.StringAttribute{
				Description:   `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/zai-org/glm-4.7-flash", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "@cf/google/gemma-3-12b-it", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"embedding_model": schema.StringAttribute{
				Description:   `Available values: "@cf/qwen/qwen3-embedding-0.6b", "@cf/baai/bge-m3", "@cf/baai/bge-large-en-v1.5", "@cf/google/embeddinggemma-300m", "google-ai-studio/gemini-embedding-001", "google-ai-studio/gemini-embedding-2-preview", "openai/text-embedding-3-small", "openai/text-embedding-3-large", "".`,
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"reranking_model": schema.StringAttribute{
				Description:   `Available values: "@cf/baai/bge-reranker-base", "".`,
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("@cf/baai/bge-reranker-base", ""),
				},
			},
			"rewrite_model": schema.StringAttribute{
				Description:   `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/zai-org/glm-4.7-flash", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "@cf/google/gemma-3-12b-it", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"summarization_model": schema.StringAttribute{
				Description: `Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/zai-org/glm-4.7-flash", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "@cf/google/gemma-3-12b-it", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".`,
				Optional:    true,
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
							Description: `Available values: "text", "number", "boolean", "datetime".`,
							Required:    true,
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
				Description:   `Available values: "super_strict_match", "close_enough", "flexible_friend", "anything_goes".`,
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{useStateForUnknownIncludingNullString()},
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
			"fusion_method": schema.StringAttribute{
				Description: `Available values: "max", "rrf".`,
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("max", "rrf"),
				},
				Default: stringdefault.StaticString("rrf"),
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
			"index_method": schema.SingleNestedAttribute{
				Description: "Controls which storage backends are used during indexing. Defaults to vector-only.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[AISearchInstanceIndexMethodModel](ctx),
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{
					"keyword": schema.BoolAttribute{
						Description: "Enable keyword (BM25) storage backend.",
						Required:    true,
					},
					"vector": schema.BoolAttribute{
						Description: "Enable vector (embedding) storage backend.",
						Required:    true,
					},
				},
			},
			"indexing_options": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[AISearchInstanceIndexingOptionsModel](ctx),
				PlanModifiers: []planmodifier.Object{
					useStateForUnknownIncludingNullObject(),
				},
				Attributes: map[string]schema.Attribute{
					"keyword_tokenizer": schema.StringAttribute{
						Description: "Tokenizer used for keyword search indexing. porter provides word-level tokenization with Porter stemming (good for natural language queries). trigram enables character-level substring matching (good for partial matches, code, identifiers). Changing this triggers a full re-index. Defaults to porter.\nAvailable values: \"porter\", \"trigram\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("porter", "trigram"),
						},
						Default: stringdefault.StaticString("porter"),
					},
				},
			},
			"public_endpoint_params": schema.SingleNestedAttribute{
				Computed:      true,
				Optional:      true,
				CustomType:    customfield.NewNestedObjectType[AISearchInstancePublicEndpointParamsModel](ctx),
				PlanModifiers: []planmodifier.Object{useStateForUnknownIncludingNullObject()},
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
							"description": schema.StringAttribute{
								Computed: true,
								Optional: true,
								Default:  stringdefault.StaticString("Finds exactly what you're looking for"),
							},
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
			"retrieval_options": schema.SingleNestedAttribute{
				Computed:      true,
				Optional:      true,
				CustomType:    customfield.NewNestedObjectType[AISearchInstanceRetrievalOptionsModel](ctx),
				PlanModifiers: []planmodifier.Object{useStateForUnknownIncludingNullObject()},
				Attributes: map[string]schema.Attribute{
					"boost_by": schema.ListNestedAttribute{
						Description: "Metadata fields to boost search results by. Each entry specifies a metadata field and an optional direction. Direction defaults to 'asc' for numeric fields and 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined custom_metadata field.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"field": schema.StringAttribute{
									Description: "Metadata field name to boost by. Use 'timestamp' for document freshness, or any custom_metadata field. Numeric and datetime fields support asc/desc directions; text/boolean fields support exists/not_exists.",
									Required:    true,
								},
								"direction": schema.StringAttribute{
									Description: "Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps). 'asc' = lower values rank higher. 'exists' = boost chunks that have the field. 'not_exists' = boost chunks that lack the field. Optional - defaults to 'asc' for numeric/datetime fields, 'exists' for text/boolean fields.\nAvailable values: \"asc\", \"desc\", \"exists\", \"not_exists\".",
									Optional:    true,
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
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("and", "or"),
						},
						Default: stringdefault.StaticString("and"),
					},
				},
			},
			"source_params": schema.SingleNestedAttribute{
				Computed:      true,
				Optional:      true,
				CustomType:    customfield.NewNestedObjectType[AISearchInstanceSourceParamsModel](ctx),
				PlanModifiers: []planmodifier.Object{useStateForUnknownIncludingNullObject()},
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
							"crawl_options": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"depth": schema.Float64Attribute{
										Optional: true,
										Validators: []validator.Float64{
											float64validator.Between(1, 100000),
										},
									},
									"include_external_links": schema.BoolAttribute{
										Computed: true,
										Optional: true,
										Default:  booldefault.StaticBool(false),
									},
									"include_subdomains": schema.BoolAttribute{
										Computed: true,
										Optional: true,
										Default:  booldefault.StaticBool(false),
									},
									"max_age": schema.Float64Attribute{
										Optional: true,
										Validators: []validator.Float64{
											float64validator.Between(0, 604800),
										},
									},
									"source": schema.StringAttribute{
										Description: `Available values: "all", "sitemaps", "links".`,
										Computed:    true,
										Optional:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"all",
												"sitemaps",
												"links",
											),
										},
										Default: stringdefault.StaticString("all"),
									},
								},
							},
							"parse_options": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"content_selector": schema.ListNestedAttribute{
										Description: "List of path-to-selector mappings for extracting specific content from crawled pages. Each entry pairs a URL glob pattern with a CSS selector. The first matching path wins. Only the matched HTML fragment is stored and indexed.",
										Optional:    true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"path": schema.StringAttribute{
													Description: "Glob pattern to match against the page URL path. Uses standard glob syntax: * matches within a segment, ** crosses directories.",
													Required:    true,
												},
												"selector": schema.StringAttribute{
													Description: "CSS selector to extract content from pages matching the path pattern. Supports standard CSS selectors including class, ID, element, and attribute selectors.",
													Required:    true,
												},
											},
										},
									},
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
								Description: `Available values: "sitemap", "feed-rss", "crawl".`,
								Computed:    true,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"sitemap",
										"feed-rss",
										"crawl",
									),
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
			"created_at": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_by": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enable": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"engine_version": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"hybrid_search_enabled": schema.BoolAttribute{
				Description:        "Deprecated — use index_method instead.",
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
				Default:            booldefault.StaticBool(false),
			},
			"last_activity": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{useStateForUnknownIncludingNullString()},
			},
			"modified_at": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"modified_by": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"namespace": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"public_endpoint_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{useStateForUnknownIncludingNullString()},
			},
			"status": schema.StringAttribute{
				Computed: true,
				Default:  stringdefault.StaticString("waiting"),
			},
			"vectorize_name": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{useStateForUnknownIncludingNullString()},
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
