// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_instance

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISearchInstanceResultEnvelope struct {
	Result AISearchInstanceModel `json:"result"`
}

type AISearchInstanceModel struct {
	ID                             types.String                                                        `tfsdk:"id" json:"id,required"`
	AccountID                      types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	Source                         types.String                                                        `tfsdk:"source" json:"source,optional"`
	Type                           types.String                                                        `tfsdk:"type" json:"type,optional"`
	AIGatewayID                    types.String                                                        `tfsdk:"ai_gateway_id" json:"ai_gateway_id,computed_optional"`
	AISearchModel                  types.String                                                        `tfsdk:"aisearch_model" json:"ai_search_model,computed_optional"`
	EmbeddingModel                 types.String                                                        `tfsdk:"embedding_model" json:"embedding_model,computed_optional"`
	RerankingModel                 types.String                                                        `tfsdk:"reranking_model" json:"reranking_model,computed_optional"`
	RewriteModel                   types.String                                                        `tfsdk:"rewrite_model" json:"rewrite_model,computed_optional"`
	SummarizationModel             types.String                                                        `tfsdk:"summarization_model" json:"summarization_model,optional,no_refresh"`
	SystemPromptAISearch           types.String                                                        `tfsdk:"system_prompt_aisearch" json:"system_prompt_ai_search,optional,no_refresh"`
	SystemPromptIndexSummarization types.String                                                        `tfsdk:"system_prompt_index_summarization" json:"system_prompt_index_summarization,optional,no_refresh"`
	SystemPromptRewriteQuery       types.String                                                        `tfsdk:"system_prompt_rewrite_query" json:"system_prompt_rewrite_query,optional,no_refresh"`
	TokenID                        types.String                                                        `tfsdk:"token_id" json:"token_id,optional"`
	CustomMetadata                 *[]*AISearchInstanceCustomMetadataModel                             `tfsdk:"custom_metadata" json:"custom_metadata,optional"`
	Metadata                       *AISearchInstanceMetadataModel                                      `tfsdk:"metadata" json:"metadata,optional"`
	Cache                          types.Bool                                                          `tfsdk:"cache" json:"cache,computed_optional"`
	CacheThreshold                 types.String                                                        `tfsdk:"cache_threshold" json:"cache_threshold,computed_optional"`
	Chunk                          types.Bool                                                          `tfsdk:"chunk" json:"chunk,computed_optional,no_refresh"`
	ChunkOverlap                   types.Int64                                                         `tfsdk:"chunk_overlap" json:"chunk_overlap,computed_optional"`
	ChunkSize                      types.Int64                                                         `tfsdk:"chunk_size" json:"chunk_size,computed_optional"`
	FusionMethod                   types.String                                                        `tfsdk:"fusion_method" json:"fusion_method,computed_optional"`
	HybridSearchEnabled            types.Bool                                                          `tfsdk:"hybrid_search_enabled" json:"hybrid_search_enabled,computed"`
	MaxNumResults                  types.Int64                                                         `tfsdk:"max_num_results" json:"max_num_results,computed_optional"`
	Paused                         types.Bool                                                          `tfsdk:"paused" json:"paused,computed_optional"`
	Reranking                      types.Bool                                                          `tfsdk:"reranking" json:"reranking,computed_optional"`
	RewriteQuery                   types.Bool                                                          `tfsdk:"rewrite_query" json:"rewrite_query,computed_optional"`
	ScoreThreshold                 types.Float64                                                       `tfsdk:"score_threshold" json:"score_threshold,computed_optional"`
	Summarization                  types.Bool                                                          `tfsdk:"summarization" json:"summarization,computed_optional,no_refresh"`
	IndexMethod                    customfield.NestedObject[AISearchInstanceIndexMethodModel]          `tfsdk:"index_method" json:"index_method,computed_optional"`
	IndexingOptions                customfield.NestedObject[AISearchInstanceIndexingOptionsModel]      `tfsdk:"indexing_options" json:"indexing_options,computed_optional"`
	PublicEndpointParams           customfield.NestedObject[AISearchInstancePublicEndpointParamsModel] `tfsdk:"public_endpoint_params" json:"public_endpoint_params,computed_optional"`
	RetrievalOptions               customfield.NestedObject[AISearchInstanceRetrievalOptionsModel]     `tfsdk:"retrieval_options" json:"retrieval_options,computed_optional"`
	SourceParams                   customfield.NestedObject[AISearchInstanceSourceParamsModel]         `tfsdk:"source_params" json:"source_params,computed_optional"`
	CreatedAt                      timetypes.RFC3339                                                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy                      types.String                                                        `tfsdk:"created_by" json:"created_by,computed"`
	Enable                         types.Bool                                                          `tfsdk:"enable" json:"enable,computed"`
	EngineVersion                  types.Float64                                                       `tfsdk:"engine_version" json:"engine_version,computed"`
	LastActivity                   timetypes.RFC3339                                                   `tfsdk:"last_activity" json:"last_activity,computed" format:"date-time"`
	ModifiedAt                     timetypes.RFC3339                                                   `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	ModifiedBy                     types.String                                                        `tfsdk:"modified_by" json:"modified_by,computed"`
	Namespace                      types.String                                                        `tfsdk:"namespace" json:"namespace,computed"`
	PublicEndpointID               types.String                                                        `tfsdk:"public_endpoint_id" json:"public_endpoint_id,computed"`
	Status                         types.String                                                        `tfsdk:"status" json:"status,computed"`
	VectorizeName                  types.String                                                        `tfsdk:"vectorize_name" json:"vectorize_name,computed"`
}

func (m AISearchInstanceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AISearchInstanceModel) MarshalJSONForUpdate(state AISearchInstanceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AISearchInstanceCustomMetadataModel struct {
	DataType  types.String `tfsdk:"data_type" json:"data_type,required"`
	FieldName types.String `tfsdk:"field_name" json:"field_name,required"`
}

type AISearchInstanceMetadataModel struct {
	CreatedFromAISearchWizard types.Bool   `tfsdk:"created_from_aisearch_wizard" json:"created_from_aisearch_wizard,optional"`
	WorkerDomain              types.String `tfsdk:"worker_domain" json:"worker_domain,optional"`
}

type AISearchInstanceIndexMethodModel struct {
	Keyword types.Bool `tfsdk:"keyword" json:"keyword,required"`
	Vector  types.Bool `tfsdk:"vector" json:"vector,required"`
}

type AISearchInstanceIndexingOptionsModel struct {
	KeywordTokenizer types.String `tfsdk:"keyword_tokenizer" json:"keyword_tokenizer,computed_optional"`
}

type AISearchInstancePublicEndpointParamsModel struct {
	AuthorizedHosts         *[]types.String                                                                            `tfsdk:"authorized_hosts" json:"authorized_hosts,optional"`
	ChatCompletionsEndpoint customfield.NestedObject[AISearchInstancePublicEndpointParamsChatCompletionsEndpointModel] `tfsdk:"chat_completions_endpoint" json:"chat_completions_endpoint,computed_optional"`
	Enabled                 types.Bool                                                                                 `tfsdk:"enabled" json:"enabled,computed_optional"`
	Mcp                     customfield.NestedObject[AISearchInstancePublicEndpointParamsMcpModel]                     `tfsdk:"mcp" json:"mcp,computed_optional"`
	RateLimit               *AISearchInstancePublicEndpointParamsRateLimitModel                                        `tfsdk:"rate_limit" json:"rate_limit,optional"`
	SearchEndpoint          customfield.NestedObject[AISearchInstancePublicEndpointParamsSearchEndpointModel]          `tfsdk:"search_endpoint" json:"search_endpoint,computed_optional"`
}

type AISearchInstancePublicEndpointParamsChatCompletionsEndpointModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed_optional"`
}

type AISearchInstancePublicEndpointParamsMcpModel struct {
	Description types.String `tfsdk:"description" json:"description,computed_optional"`
	Disabled    types.Bool   `tfsdk:"disabled" json:"disabled,computed_optional"`
}

type AISearchInstancePublicEndpointParamsRateLimitModel struct {
	PeriodMs  types.Int64  `tfsdk:"period_ms" json:"period_ms,optional"`
	Requests  types.Int64  `tfsdk:"requests" json:"requests,optional"`
	Technique types.String `tfsdk:"technique" json:"technique,optional"`
}

type AISearchInstancePublicEndpointParamsSearchEndpointModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed_optional"`
}

type AISearchInstanceRetrievalOptionsModel struct {
	BoostBy          *[]*AISearchInstanceRetrievalOptionsBoostByModel `tfsdk:"boost_by" json:"boost_by,optional"`
	KeywordMatchMode types.String                                     `tfsdk:"keyword_match_mode" json:"keyword_match_mode,computed_optional"`
}

type AISearchInstanceRetrievalOptionsBoostByModel struct {
	Field     types.String `tfsdk:"field" json:"field,required"`
	Direction types.String `tfsdk:"direction" json:"direction,optional"`
}

type AISearchInstanceSourceParamsModel struct {
	ExcludeItems   *[]types.String                                                       `tfsdk:"exclude_items" json:"exclude_items,optional"`
	IncludeItems   *[]types.String                                                       `tfsdk:"include_items" json:"include_items,optional"`
	Prefix         types.String                                                          `tfsdk:"prefix" json:"prefix,optional"`
	R2Jurisdiction types.String                                                          `tfsdk:"r2_jurisdiction" json:"r2_jurisdiction,computed_optional"`
	WebCrawler     customfield.NestedObject[AISearchInstanceSourceParamsWebCrawlerModel] `tfsdk:"web_crawler" json:"web_crawler,computed_optional"`
}

type AISearchInstanceSourceParamsWebCrawlerModel struct {
	CrawlOptions *AISearchInstanceSourceParamsWebCrawlerCrawlOptionsModel `tfsdk:"crawl_options" json:"crawl_options,optional"`
	ParseOptions *AISearchInstanceSourceParamsWebCrawlerParseOptionsModel `tfsdk:"parse_options" json:"parse_options,optional"`
	ParseType    types.String                                             `tfsdk:"parse_type" json:"parse_type,computed_optional"`
	StoreOptions *AISearchInstanceSourceParamsWebCrawlerStoreOptionsModel `tfsdk:"store_options" json:"store_options,optional"`
}

type AISearchInstanceSourceParamsWebCrawlerCrawlOptionsModel struct {
	Depth                types.Float64 `tfsdk:"depth" json:"depth,optional"`
	IncludeExternalLinks types.Bool    `tfsdk:"include_external_links" json:"include_external_links,computed_optional"`
	IncludeSubdomains    types.Bool    `tfsdk:"include_subdomains" json:"include_subdomains,computed_optional"`
	MaxAge               types.Float64 `tfsdk:"max_age" json:"max_age,optional"`
	Source               types.String  `tfsdk:"source" json:"source,computed_optional"`
}

type AISearchInstanceSourceParamsWebCrawlerParseOptionsModel struct {
	ContentSelector     *[]*AISearchInstanceSourceParamsWebCrawlerParseOptionsContentSelectorModel `tfsdk:"content_selector" json:"content_selector,optional"`
	IncludeHeaders      *map[string]types.String                                                   `tfsdk:"include_headers" json:"include_headers,optional"`
	IncludeImages       types.Bool                                                                 `tfsdk:"include_images" json:"include_images,computed_optional"`
	SpecificSitemaps    *[]types.String                                                            `tfsdk:"specific_sitemaps" json:"specific_sitemaps,optional"`
	UseBrowserRendering types.Bool                                                                 `tfsdk:"use_browser_rendering" json:"use_browser_rendering,computed_optional"`
}

type AISearchInstanceSourceParamsWebCrawlerParseOptionsContentSelectorModel struct {
	Path     types.String `tfsdk:"path" json:"path,required"`
	Selector types.String `tfsdk:"selector" json:"selector,required"`
}

type AISearchInstanceSourceParamsWebCrawlerStoreOptionsModel struct {
	StorageID      types.String `tfsdk:"storage_id" json:"storage_id,required"`
	R2Jurisdiction types.String `tfsdk:"r2_jurisdiction" json:"r2_jurisdiction,computed_optional"`
	StorageType    types.String `tfsdk:"storage_type" json:"storage_type,optional"`
}
