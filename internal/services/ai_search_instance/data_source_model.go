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

type AISearchInstanceResultDataSourceEnvelope struct {
	Result AISearchInstanceDataSourceModel `json:"result,computed"`
}

type AISearchInstanceDataSourceModel struct {
	ID                   types.String                                                                  `tfsdk:"id" path:"id,computed_optional"`
	AccountID            types.String                                                                  `tfsdk:"account_id" path:"account_id,required"`
	AIGatewayID          types.String                                                                  `tfsdk:"ai_gateway_id" json:"ai_gateway_id,computed"`
	AISearchModel        types.String                                                                  `tfsdk:"aisearch_model" json:"ai_search_model,computed"`
	Cache                types.Bool                                                                    `tfsdk:"cache" json:"cache,computed"`
	CacheThreshold       types.String                                                                  `tfsdk:"cache_threshold" json:"cache_threshold,computed"`
	ChunkOverlap         types.Int64                                                                   `tfsdk:"chunk_overlap" json:"chunk_overlap,computed"`
	ChunkSize            types.Int64                                                                   `tfsdk:"chunk_size" json:"chunk_size,computed"`
	CreatedAt            timetypes.RFC3339                                                             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy            types.String                                                                  `tfsdk:"created_by" json:"created_by,computed"`
	EmbeddingModel       types.String                                                                  `tfsdk:"embedding_model" json:"embedding_model,computed"`
	Enable               types.Bool                                                                    `tfsdk:"enable" json:"enable,computed"`
	EngineVersion        types.Float64                                                                 `tfsdk:"engine_version" json:"engine_version,computed"`
	FusionMethod         types.String                                                                  `tfsdk:"fusion_method" json:"fusion_method,computed"`
	HybridSearchEnabled  types.Bool                                                                    `tfsdk:"hybrid_search_enabled" json:"hybrid_search_enabled,computed"`
	LastActivity         timetypes.RFC3339                                                             `tfsdk:"last_activity" json:"last_activity,computed" format:"date-time"`
	MaxNumResults        types.Int64                                                                   `tfsdk:"max_num_results" json:"max_num_results,computed"`
	ModifiedAt           timetypes.RFC3339                                                             `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	ModifiedBy           types.String                                                                  `tfsdk:"modified_by" json:"modified_by,computed"`
	Namespace            types.String                                                                  `tfsdk:"namespace" json:"namespace,computed"`
	Paused               types.Bool                                                                    `tfsdk:"paused" json:"paused,computed"`
	PublicEndpointID     types.String                                                                  `tfsdk:"public_endpoint_id" json:"public_endpoint_id,computed"`
	Reranking            types.Bool                                                                    `tfsdk:"reranking" json:"reranking,computed"`
	RerankingModel       types.String                                                                  `tfsdk:"reranking_model" json:"reranking_model,computed"`
	RewriteModel         types.String                                                                  `tfsdk:"rewrite_model" json:"rewrite_model,computed"`
	RewriteQuery         types.Bool                                                                    `tfsdk:"rewrite_query" json:"rewrite_query,computed"`
	ScoreThreshold       types.Float64                                                                 `tfsdk:"score_threshold" json:"score_threshold,computed"`
	Source               types.String                                                                  `tfsdk:"source" json:"source,computed"`
	Status               types.String                                                                  `tfsdk:"status" json:"status,computed"`
	TokenID              types.String                                                                  `tfsdk:"token_id" json:"token_id,computed"`
	Type                 types.String                                                                  `tfsdk:"type" json:"type,computed"`
	CustomMetadata       customfield.NestedObjectList[AISearchInstanceCustomMetadataDataSourceModel]   `tfsdk:"custom_metadata" json:"custom_metadata,computed"`
	IndexMethod          customfield.NestedObject[AISearchInstanceIndexMethodDataSourceModel]          `tfsdk:"index_method" json:"index_method,computed"`
	IndexingOptions      customfield.NestedObject[AISearchInstanceIndexingOptionsDataSourceModel]      `tfsdk:"indexing_options" json:"indexing_options,computed"`
	Metadata             customfield.NestedObject[AISearchInstanceMetadataDataSourceModel]             `tfsdk:"metadata" json:"metadata,computed"`
	PublicEndpointParams customfield.NestedObject[AISearchInstancePublicEndpointParamsDataSourceModel] `tfsdk:"public_endpoint_params" json:"public_endpoint_params,computed"`
	RetrievalOptions     customfield.NestedObject[AISearchInstanceRetrievalOptionsDataSourceModel]     `tfsdk:"retrieval_options" json:"retrieval_options,computed"`
	SourceParams         customfield.NestedObject[AISearchInstanceSourceParamsDataSourceModel]         `tfsdk:"source_params" json:"source_params,computed"`
	Filter               *AISearchInstanceFindOneByDataSourceModel                                     `tfsdk:"filter"`
}

func (m *AISearchInstanceDataSourceModel) toReadParams(_ context.Context) (params ai_search.InstanceReadParams, diags diag.Diagnostics) {
	params = ai_search.InstanceReadParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *AISearchInstanceDataSourceModel) toListParams(_ context.Context) (params ai_search.InstanceListParams, diags diag.Diagnostics) {
	params = ai_search.InstanceListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Namespace.IsNull() {
		params.Namespace = cloudflare.F(m.Filter.Namespace.ValueString())
	}
	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(ai_search.InstanceListParamsOrderBy(m.Filter.OrderBy.ValueString()))
	}
	if !m.Filter.OrderByDirection.IsNull() {
		params.OrderByDirection = cloudflare.F(ai_search.InstanceListParamsOrderByDirection(m.Filter.OrderByDirection.ValueString()))
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type AISearchInstanceCustomMetadataDataSourceModel struct {
	DataType  types.String `tfsdk:"data_type" json:"data_type,computed"`
	FieldName types.String `tfsdk:"field_name" json:"field_name,computed"`
}

type AISearchInstanceIndexMethodDataSourceModel struct {
	Keyword types.Bool `tfsdk:"keyword" json:"keyword,computed"`
	Vector  types.Bool `tfsdk:"vector" json:"vector,computed"`
}

type AISearchInstanceIndexingOptionsDataSourceModel struct {
	KeywordTokenizer types.String `tfsdk:"keyword_tokenizer" json:"keyword_tokenizer,computed"`
}

type AISearchInstanceMetadataDataSourceModel struct {
	CreatedFromAISearchWizard types.Bool   `tfsdk:"created_from_aisearch_wizard" json:"created_from_aisearch_wizard,computed"`
	WorkerDomain              types.String `tfsdk:"worker_domain" json:"worker_domain,computed"`
}

type AISearchInstancePublicEndpointParamsDataSourceModel struct {
	AuthorizedHosts         customfield.List[types.String]                                                                       `tfsdk:"authorized_hosts" json:"authorized_hosts,computed"`
	ChatCompletionsEndpoint customfield.NestedObject[AISearchInstancePublicEndpointParamsChatCompletionsEndpointDataSourceModel] `tfsdk:"chat_completions_endpoint" json:"chat_completions_endpoint,computed"`
	Enabled                 types.Bool                                                                                           `tfsdk:"enabled" json:"enabled,computed"`
	Mcp                     customfield.NestedObject[AISearchInstancePublicEndpointParamsMcpDataSourceModel]                     `tfsdk:"mcp" json:"mcp,computed"`
	RateLimit               customfield.NestedObject[AISearchInstancePublicEndpointParamsRateLimitDataSourceModel]               `tfsdk:"rate_limit" json:"rate_limit,computed"`
	SearchEndpoint          customfield.NestedObject[AISearchInstancePublicEndpointParamsSearchEndpointDataSourceModel]          `tfsdk:"search_endpoint" json:"search_endpoint,computed"`
}

type AISearchInstancePublicEndpointParamsChatCompletionsEndpointDataSourceModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed"`
}

type AISearchInstancePublicEndpointParamsMcpDataSourceModel struct {
	Description types.String `tfsdk:"description" json:"description,computed"`
	Disabled    types.Bool   `tfsdk:"disabled" json:"disabled,computed"`
}

type AISearchInstancePublicEndpointParamsRateLimitDataSourceModel struct {
	PeriodMs  types.Int64  `tfsdk:"period_ms" json:"period_ms,computed"`
	Requests  types.Int64  `tfsdk:"requests" json:"requests,computed"`
	Technique types.String `tfsdk:"technique" json:"technique,computed"`
}

type AISearchInstancePublicEndpointParamsSearchEndpointDataSourceModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed"`
}

type AISearchInstanceRetrievalOptionsDataSourceModel struct {
	BoostBy          customfield.NestedObjectList[AISearchInstanceRetrievalOptionsBoostByDataSourceModel] `tfsdk:"boost_by" json:"boost_by,computed"`
	KeywordMatchMode types.String                                                                         `tfsdk:"keyword_match_mode" json:"keyword_match_mode,computed"`
}

type AISearchInstanceRetrievalOptionsBoostByDataSourceModel struct {
	Field     types.String `tfsdk:"field" json:"field,computed"`
	Direction types.String `tfsdk:"direction" json:"direction,computed"`
}

type AISearchInstanceSourceParamsDataSourceModel struct {
	ExcludeItems   customfield.List[types.String]                                                  `tfsdk:"exclude_items" json:"exclude_items,computed"`
	IncludeItems   customfield.List[types.String]                                                  `tfsdk:"include_items" json:"include_items,computed"`
	Prefix         types.String                                                                    `tfsdk:"prefix" json:"prefix,computed"`
	R2Jurisdiction types.String                                                                    `tfsdk:"r2_jurisdiction" json:"r2_jurisdiction,computed"`
	WebCrawler     customfield.NestedObject[AISearchInstanceSourceParamsWebCrawlerDataSourceModel] `tfsdk:"web_crawler" json:"web_crawler,computed"`
}

type AISearchInstanceSourceParamsWebCrawlerDataSourceModel struct {
	CrawlOptions customfield.NestedObject[AISearchInstanceSourceParamsWebCrawlerCrawlOptionsDataSourceModel] `tfsdk:"crawl_options" json:"crawl_options,computed"`
	ParseOptions customfield.NestedObject[AISearchInstanceSourceParamsWebCrawlerParseOptionsDataSourceModel] `tfsdk:"parse_options" json:"parse_options,computed"`
	ParseType    types.String                                                                                `tfsdk:"parse_type" json:"parse_type,computed"`
	StoreOptions customfield.NestedObject[AISearchInstanceSourceParamsWebCrawlerStoreOptionsDataSourceModel] `tfsdk:"store_options" json:"store_options,computed"`
}

type AISearchInstanceSourceParamsWebCrawlerCrawlOptionsDataSourceModel struct {
	Depth                types.Float64 `tfsdk:"depth" json:"depth,computed"`
	IncludeExternalLinks types.Bool    `tfsdk:"include_external_links" json:"include_external_links,computed"`
	IncludeSubdomains    types.Bool    `tfsdk:"include_subdomains" json:"include_subdomains,computed"`
	MaxAge               types.Float64 `tfsdk:"max_age" json:"max_age,computed"`
	Source               types.String  `tfsdk:"source" json:"source,computed"`
}

type AISearchInstanceSourceParamsWebCrawlerParseOptionsDataSourceModel struct {
	ContentSelector     customfield.NestedObjectList[AISearchInstanceSourceParamsWebCrawlerParseOptionsContentSelectorDataSourceModel] `tfsdk:"content_selector" json:"content_selector,computed"`
	IncludeHeaders      customfield.Map[types.String]                                                                                  `tfsdk:"include_headers" json:"include_headers,computed"`
	IncludeImages       types.Bool                                                                                                     `tfsdk:"include_images" json:"include_images,computed"`
	SpecificSitemaps    customfield.List[types.String]                                                                                 `tfsdk:"specific_sitemaps" json:"specific_sitemaps,computed"`
	UseBrowserRendering types.Bool                                                                                                     `tfsdk:"use_browser_rendering" json:"use_browser_rendering,computed"`
}

type AISearchInstanceSourceParamsWebCrawlerParseOptionsContentSelectorDataSourceModel struct {
	Path     types.String `tfsdk:"path" json:"path,computed"`
	Selector types.String `tfsdk:"selector" json:"selector,computed"`
}

type AISearchInstanceSourceParamsWebCrawlerStoreOptionsDataSourceModel struct {
	StorageID      types.String `tfsdk:"storage_id" json:"storage_id,computed"`
	R2Jurisdiction types.String `tfsdk:"r2_jurisdiction" json:"r2_jurisdiction,computed"`
	StorageType    types.String `tfsdk:"storage_type" json:"storage_type,computed"`
}

type AISearchInstanceFindOneByDataSourceModel struct {
	Namespace        types.String `tfsdk:"namespace" query:"namespace,optional"`
	OrderBy          types.String `tfsdk:"order_by" query:"order_by,computed_optional"`
	OrderByDirection types.String `tfsdk:"order_by_direction" query:"order_by_direction,computed_optional"`
	Search           types.String `tfsdk:"search" query:"search,optional"`
}
