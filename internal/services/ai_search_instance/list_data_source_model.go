// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_instance

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ai_search"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISearchInstancesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AISearchInstancesResultDataSourceModel] `json:"result,computed"`
}

type AISearchInstancesDataSourceModel struct {
	AccountID types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	Search    types.String                                                         `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AISearchInstancesResultDataSourceModel] `tfsdk:"result"`
}

func (m *AISearchInstancesDataSourceModel) toListParams(_ context.Context) (params ai_search.InstanceListParams, diags diag.Diagnostics) {
	params = ai_search.InstanceListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type AISearchInstancesResultDataSourceModel struct {
	ID                             types.String                                                                   `tfsdk:"id" json:"id,computed"`
	AccountID                      types.String                                                                   `tfsdk:"account_id" json:"account_id,computed"`
	AccountTag                     types.String                                                                   `tfsdk:"account_tag" json:"account_tag,computed"`
	CreatedAt                      timetypes.RFC3339                                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	InternalID                     types.String                                                                   `tfsdk:"internal_id" json:"internal_id,computed"`
	ModifiedAt                     timetypes.RFC3339                                                              `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Source                         types.String                                                                   `tfsdk:"source" json:"source,computed"`
	Type                           types.String                                                                   `tfsdk:"type" json:"type,computed"`
	VectorizeName                  types.String                                                                   `tfsdk:"vectorize_name" json:"vectorize_name,computed"`
	AIGatewayID                    types.String                                                                   `tfsdk:"ai_gateway_id" json:"ai_gateway_id,computed"`
	AISearchModel                  types.String                                                                   `tfsdk:"aisearch_model" json:"ai_search_model,computed"`
	Cache                          types.Bool                                                                     `tfsdk:"cache" json:"cache,computed"`
	CacheThreshold                 types.String                                                                   `tfsdk:"cache_threshold" json:"cache_threshold,computed"`
	Chunk                          types.Bool                                                                     `tfsdk:"chunk" json:"chunk,computed"`
	ChunkOverlap                   types.Int64                                                                    `tfsdk:"chunk_overlap" json:"chunk_overlap,computed"`
	ChunkSize                      types.Int64                                                                    `tfsdk:"chunk_size" json:"chunk_size,computed"`
	CreatedBy                      types.String                                                                   `tfsdk:"created_by" json:"created_by,computed"`
	CustomMetadata                 customfield.NestedObjectList[AISearchInstancesCustomMetadataDataSourceModel]   `tfsdk:"custom_metadata" json:"custom_metadata,computed"`
	EmbeddingModel                 types.String                                                                   `tfsdk:"embedding_model" json:"embedding_model,computed"`
	Enable                         types.Bool                                                                     `tfsdk:"enable" json:"enable,computed"`
	EngineVersion                  types.Float64                                                                  `tfsdk:"engine_version" json:"engine_version,computed"`
	HybridSearchEnabled            types.Bool                                                                     `tfsdk:"hybrid_search_enabled" json:"hybrid_search_enabled,computed"`
	LastActivity                   timetypes.RFC3339                                                              `tfsdk:"last_activity" json:"last_activity,computed" format:"date-time"`
	MaxNumResults                  types.Int64                                                                    `tfsdk:"max_num_results" json:"max_num_results,computed"`
	Metadata                       customfield.NestedObject[AISearchInstancesMetadataDataSourceModel]             `tfsdk:"metadata" json:"metadata,computed"`
	ModifiedBy                     types.String                                                                   `tfsdk:"modified_by" json:"modified_by,computed"`
	Paused                         types.Bool                                                                     `tfsdk:"paused" json:"paused,computed"`
	PublicEndpointID               types.String                                                                   `tfsdk:"public_endpoint_id" json:"public_endpoint_id,computed"`
	PublicEndpointParams           customfield.NestedObject[AISearchInstancesPublicEndpointParamsDataSourceModel] `tfsdk:"public_endpoint_params" json:"public_endpoint_params,computed"`
	Reranking                      types.Bool                                                                     `tfsdk:"reranking" json:"reranking,computed"`
	RerankingModel                 types.String                                                                   `tfsdk:"reranking_model" json:"reranking_model,computed"`
	RewriteModel                   types.String                                                                   `tfsdk:"rewrite_model" json:"rewrite_model,computed"`
	RewriteQuery                   types.Bool                                                                     `tfsdk:"rewrite_query" json:"rewrite_query,computed"`
	ScoreThreshold                 types.Float64                                                                  `tfsdk:"score_threshold" json:"score_threshold,computed"`
	SourceParams                   customfield.NestedObject[AISearchInstancesSourceParamsDataSourceModel]         `tfsdk:"source_params" json:"source_params,computed"`
	Status                         types.String                                                                   `tfsdk:"status" json:"status,computed"`
	Summarization                  types.Bool                                                                     `tfsdk:"summarization" json:"summarization,computed"`
	SummarizationModel             types.String                                                                   `tfsdk:"summarization_model" json:"summarization_model,computed"`
	SystemPromptAISearch           types.String                                                                   `tfsdk:"system_prompt_aisearch" json:"system_prompt_ai_search,computed"`
	SystemPromptIndexSummarization types.String                                                                   `tfsdk:"system_prompt_index_summarization" json:"system_prompt_index_summarization,computed"`
	SystemPromptRewriteQuery       types.String                                                                   `tfsdk:"system_prompt_rewrite_query" json:"system_prompt_rewrite_query,computed"`
	TokenID                        types.String                                                                   `tfsdk:"token_id" json:"token_id,computed"`
	VectorizeActiveNamespace       types.String                                                                   `tfsdk:"vectorize_active_namespace" json:"vectorize_active_namespace,computed"`
}

type AISearchInstancesCustomMetadataDataSourceModel struct {
	DataType  types.String `tfsdk:"data_type" json:"data_type,computed"`
	FieldName types.String `tfsdk:"field_name" json:"field_name,computed"`
}

type AISearchInstancesMetadataDataSourceModel struct {
	CreatedFromAISearchWizard types.Bool   `tfsdk:"created_from_aisearch_wizard" json:"created_from_aisearch_wizard,computed"`
	WorkerDomain              types.String `tfsdk:"worker_domain" json:"worker_domain,computed"`
}

type AISearchInstancesPublicEndpointParamsDataSourceModel struct {
	AuthorizedHosts         customfield.List[types.String]                                                                        `tfsdk:"authorized_hosts" json:"authorized_hosts,computed"`
	ChatCompletionsEndpoint customfield.NestedObject[AISearchInstancesPublicEndpointParamsChatCompletionsEndpointDataSourceModel] `tfsdk:"chat_completions_endpoint" json:"chat_completions_endpoint,computed"`
	Enabled                 types.Bool                                                                                            `tfsdk:"enabled" json:"enabled,computed"`
	Mcp                     customfield.NestedObject[AISearchInstancesPublicEndpointParamsMcpDataSourceModel]                     `tfsdk:"mcp" json:"mcp,computed"`
	RateLimit               customfield.NestedObject[AISearchInstancesPublicEndpointParamsRateLimitDataSourceModel]               `tfsdk:"rate_limit" json:"rate_limit,computed"`
	SearchEndpoint          customfield.NestedObject[AISearchInstancesPublicEndpointParamsSearchEndpointDataSourceModel]          `tfsdk:"search_endpoint" json:"search_endpoint,computed"`
}

type AISearchInstancesPublicEndpointParamsChatCompletionsEndpointDataSourceModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed"`
}

type AISearchInstancesPublicEndpointParamsMcpDataSourceModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed"`
}

type AISearchInstancesPublicEndpointParamsRateLimitDataSourceModel struct {
	PeriodMs  types.Int64  `tfsdk:"period_ms" json:"period_ms,computed"`
	Requests  types.Int64  `tfsdk:"requests" json:"requests,computed"`
	Technique types.String `tfsdk:"technique" json:"technique,computed"`
}

type AISearchInstancesPublicEndpointParamsSearchEndpointDataSourceModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed"`
}

type AISearchInstancesSourceParamsDataSourceModel struct {
	ExcludeItems   customfield.List[types.String]                                                   `tfsdk:"exclude_items" json:"exclude_items,computed"`
	IncludeItems   customfield.List[types.String]                                                   `tfsdk:"include_items" json:"include_items,computed"`
	Prefix         types.String                                                                     `tfsdk:"prefix" json:"prefix,computed"`
	R2Jurisdiction types.String                                                                     `tfsdk:"r2_jurisdiction" json:"r2_jurisdiction,computed"`
	WebCrawler     customfield.NestedObject[AISearchInstancesSourceParamsWebCrawlerDataSourceModel] `tfsdk:"web_crawler" json:"web_crawler,computed"`
}

type AISearchInstancesSourceParamsWebCrawlerDataSourceModel struct {
	ParseOptions customfield.NestedObject[AISearchInstancesSourceParamsWebCrawlerParseOptionsDataSourceModel] `tfsdk:"parse_options" json:"parse_options,computed"`
	ParseType    types.String                                                                                 `tfsdk:"parse_type" json:"parse_type,computed"`
	StoreOptions customfield.NestedObject[AISearchInstancesSourceParamsWebCrawlerStoreOptionsDataSourceModel] `tfsdk:"store_options" json:"store_options,computed"`
}

type AISearchInstancesSourceParamsWebCrawlerParseOptionsDataSourceModel struct {
	IncludeHeaders      customfield.Map[types.String]  `tfsdk:"include_headers" json:"include_headers,computed"`
	IncludeImages       types.Bool                     `tfsdk:"include_images" json:"include_images,computed"`
	SpecificSitemaps    customfield.List[types.String] `tfsdk:"specific_sitemaps" json:"specific_sitemaps,computed"`
	UseBrowserRendering types.Bool                     `tfsdk:"use_browser_rendering" json:"use_browser_rendering,computed"`
}

type AISearchInstancesSourceParamsWebCrawlerStoreOptionsDataSourceModel struct {
	StorageID      types.String `tfsdk:"storage_id" json:"storage_id,computed"`
	R2Jurisdiction types.String `tfsdk:"r2_jurisdiction" json:"r2_jurisdiction,computed"`
	StorageType    types.String `tfsdk:"storage_type" json:"storage_type,computed"`
}
