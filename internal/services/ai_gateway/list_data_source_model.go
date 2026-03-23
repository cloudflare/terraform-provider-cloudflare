// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ai_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AIGatewaysResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AIGatewaysResultDataSourceModel] `json:"result,computed"`
}

type AIGatewaysDataSourceModel struct {
	AccountID types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	Search    types.String                                                  `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                   `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AIGatewaysResultDataSourceModel] `tfsdk:"result"`
}

func (m *AIGatewaysDataSourceModel) toListParams(_ context.Context) (params ai_gateway.AIGatewayListParams, diags diag.Diagnostics) {
	params = ai_gateway.AIGatewayListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type AIGatewaysResultDataSourceModel struct {
	ID                      types.String                                                `tfsdk:"id" json:"id,computed"`
	CacheInvalidateOnUpdate types.Bool                                                  `tfsdk:"cache_invalidate_on_update" json:"cache_invalidate_on_update,computed"`
	CacheTTL                types.Int64                                                 `tfsdk:"cache_ttl" json:"cache_ttl,computed"`
	CollectLogs             types.Bool                                                  `tfsdk:"collect_logs" json:"collect_logs,computed"`
	CreatedAt               timetypes.RFC3339                                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt              timetypes.RFC3339                                           `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	RateLimitingInterval    types.Int64                                                 `tfsdk:"rate_limiting_interval" json:"rate_limiting_interval,computed"`
	RateLimitingLimit       types.Int64                                                 `tfsdk:"rate_limiting_limit" json:"rate_limiting_limit,computed"`
	Authentication          types.Bool                                                  `tfsdk:"authentication" json:"authentication,computed"`
	DLP                     customfield.NestedObject[AIGatewaysDLPDataSourceModel]      `tfsdk:"dlp" json:"dlp,computed"`
	IsDefault               types.Bool                                                  `tfsdk:"is_default" json:"is_default,computed"`
	LogManagement           types.Int64                                                 `tfsdk:"log_management" json:"log_management,computed"`
	LogManagementStrategy   types.String                                                `tfsdk:"log_management_strategy" json:"log_management_strategy,computed"`
	Logpush                 types.Bool                                                  `tfsdk:"logpush" json:"logpush,computed"`
	LogpushPublicKey        types.String                                                `tfsdk:"logpush_public_key" json:"logpush_public_key,computed"`
	Otel                    customfield.NestedObjectList[AIGatewaysOtelDataSourceModel] `tfsdk:"otel" json:"otel,computed"`
	RateLimitingTechnique   types.String                                                `tfsdk:"rate_limiting_technique" json:"rate_limiting_technique,computed"`
	RetryBackoff            types.String                                                `tfsdk:"retry_backoff" json:"retry_backoff,computed"`
	RetryDelay              types.Int64                                                 `tfsdk:"retry_delay" json:"retry_delay,computed"`
	RetryMaxAttempts        types.Int64                                                 `tfsdk:"retry_max_attempts" json:"retry_max_attempts,computed"`
	StoreID                 types.String                                                `tfsdk:"store_id" json:"store_id,computed"`
	Stripe                  customfield.NestedObject[AIGatewaysStripeDataSourceModel]   `tfsdk:"stripe" json:"stripe,computed"`
	WorkersAIBillingMode    types.String                                                `tfsdk:"workers_ai_billing_mode" json:"workers_ai_billing_mode,computed"`
	Zdr                     types.Bool                                                  `tfsdk:"zdr" json:"zdr,computed"`
}

type AIGatewaysDLPDataSourceModel struct {
	Action   types.String                                                       `tfsdk:"action" json:"action,computed"`
	Enabled  types.Bool                                                         `tfsdk:"enabled" json:"enabled,computed"`
	Profiles customfield.List[types.String]                                     `tfsdk:"profiles" json:"profiles,computed"`
	Policies customfield.NestedObjectList[AIGatewaysDLPPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
}

type AIGatewaysDLPPoliciesDataSourceModel struct {
	ID       types.String                   `tfsdk:"id" json:"id,computed"`
	Action   types.String                   `tfsdk:"action" json:"action,computed"`
	Check    customfield.List[types.String] `tfsdk:"check" json:"check,computed"`
	Enabled  types.Bool                     `tfsdk:"enabled" json:"enabled,computed"`
	Profiles customfield.List[types.String] `tfsdk:"profiles" json:"profiles,computed"`
}

type AIGatewaysOtelDataSourceModel struct {
	Authorization types.String                  `tfsdk:"authorization" json:"authorization,computed"`
	Headers       customfield.Map[types.String] `tfsdk:"headers" json:"headers,computed"`
	URL           types.String                  `tfsdk:"url" json:"url,computed"`
	ContentType   types.String                  `tfsdk:"content_type" json:"content_type,computed"`
}

type AIGatewaysStripeDataSourceModel struct {
	Authorization types.String                                                             `tfsdk:"authorization" json:"authorization,computed"`
	UsageEvents   customfield.NestedObjectList[AIGatewaysStripeUsageEventsDataSourceModel] `tfsdk:"usage_events" json:"usage_events,computed"`
}

type AIGatewaysStripeUsageEventsDataSourceModel struct {
	Payload types.String `tfsdk:"payload" json:"payload,computed"`
}
