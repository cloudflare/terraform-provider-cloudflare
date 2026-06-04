// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AIGatewaysResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AIGatewaysResultDataSourceModel] `json:"result,computed"`
}

type AIGatewaysDataSourceModel struct {
	AccountID types.String                                                  `tfsdk:"account_id" path:"account_id,optional"`
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
	ID                      types.String                                                   `tfsdk:"id" json:"id,computed"`
	CacheInvalidateOnUpdate types.Bool                                                     `tfsdk:"cache_invalidate_on_update" json:"cache_invalidate_on_update,computed"`
	CacheTTL                types.Int64                                                    `tfsdk:"cache_ttl" json:"cache_ttl,computed"`
	CollectLogs             types.Bool                                                     `tfsdk:"collect_logs" json:"collect_logs,computed"`
	CreatedAt               timetypes.RFC3339                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt              timetypes.RFC3339                                              `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	RateLimitingInterval    types.Int64                                                    `tfsdk:"rate_limiting_interval" json:"rate_limiting_interval,computed"`
	RateLimitingLimit       types.Int64                                                    `tfsdk:"rate_limiting_limit" json:"rate_limiting_limit,computed"`
	Authentication          types.Bool                                                     `tfsdk:"authentication" json:"authentication,computed"`
	DLP                     customfield.NestedObject[AIGatewaysDLPDataSourceModel]         `tfsdk:"dlp" json:"dlp,computed"`
	Guardrails              customfield.NestedObject[AIGatewaysGuardrailsDataSourceModel]  `tfsdk:"guardrails" json:"guardrails,computed"`
	IsDefault               types.Bool                                                     `tfsdk:"is_default" json:"is_default,computed"`
	LogManagement           types.Int64                                                    `tfsdk:"log_management" json:"log_management,computed"`
	LogManagementStrategy   types.String                                                   `tfsdk:"log_management_strategy" json:"log_management_strategy,computed"`
	Logpush                 types.Bool                                                     `tfsdk:"logpush" json:"logpush,computed"`
	LogpushPublicKey        types.String                                                   `tfsdk:"logpush_public_key" json:"logpush_public_key,computed"`
	Otel                    customfield.NestedObjectList[AIGatewaysOtelDataSourceModel]    `tfsdk:"otel" json:"otel,computed"`
	RateLimitingTechnique   types.String                                                   `tfsdk:"rate_limiting_technique" json:"rate_limiting_technique,computed"`
	RetryBackoff            types.String                                                   `tfsdk:"retry_backoff" json:"retry_backoff,computed"`
	RetryDelay              types.Int64                                                    `tfsdk:"retry_delay" json:"retry_delay,computed"`
	RetryMaxAttempts        types.Int64                                                    `tfsdk:"retry_max_attempts" json:"retry_max_attempts,computed"`
	SpendLimits             customfield.NestedObject[AIGatewaysSpendLimitsDataSourceModel] `tfsdk:"spend_limits" json:"spend_limits,computed"`
	StoreID                 types.String                                                   `tfsdk:"store_id" json:"store_id,computed"`
	Stripe                  customfield.NestedObject[AIGatewaysStripeDataSourceModel]      `tfsdk:"stripe" json:"stripe,computed"`
	WorkersAIBillingMode    types.String                                                   `tfsdk:"workers_ai_billing_mode" json:"workers_ai_billing_mode,computed"`
	Zdr                     types.Bool                                                     `tfsdk:"zdr" json:"zdr,computed"`
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

type AIGatewaysGuardrailsDataSourceModel struct {
	Prompt   customfield.NestedObject[AIGatewaysGuardrailsPromptDataSourceModel]   `tfsdk:"prompt" json:"prompt,computed"`
	Response customfield.NestedObject[AIGatewaysGuardrailsResponseDataSourceModel] `tfsdk:"response" json:"response,computed"`
}

type AIGatewaysGuardrailsPromptDataSourceModel struct {
	P1  types.String `tfsdk:"p1" json:"P1,computed"`
	S1  types.String `tfsdk:"s1" json:"S1,computed"`
	S10 types.String `tfsdk:"s10" json:"S10,computed"`
	S11 types.String `tfsdk:"s11" json:"S11,computed"`
	S12 types.String `tfsdk:"s12" json:"S12,computed"`
	S13 types.String `tfsdk:"s13" json:"S13,computed"`
	S2  types.String `tfsdk:"s2" json:"S2,computed"`
	S3  types.String `tfsdk:"s3" json:"S3,computed"`
	S4  types.String `tfsdk:"s4" json:"S4,computed"`
	S5  types.String `tfsdk:"s5" json:"S5,computed"`
	S6  types.String `tfsdk:"s6" json:"S6,computed"`
	S7  types.String `tfsdk:"s7" json:"S7,computed"`
	S8  types.String `tfsdk:"s8" json:"S8,computed"`
	S9  types.String `tfsdk:"s9" json:"S9,computed"`
}

type AIGatewaysGuardrailsResponseDataSourceModel struct {
	P1  types.String `tfsdk:"p1" json:"P1,computed"`
	S1  types.String `tfsdk:"s1" json:"S1,computed"`
	S10 types.String `tfsdk:"s10" json:"S10,computed"`
	S11 types.String `tfsdk:"s11" json:"S11,computed"`
	S12 types.String `tfsdk:"s12" json:"S12,computed"`
	S13 types.String `tfsdk:"s13" json:"S13,computed"`
	S2  types.String `tfsdk:"s2" json:"S2,computed"`
	S3  types.String `tfsdk:"s3" json:"S3,computed"`
	S4  types.String `tfsdk:"s4" json:"S4,computed"`
	S5  types.String `tfsdk:"s5" json:"S5,computed"`
	S6  types.String `tfsdk:"s6" json:"S6,computed"`
	S7  types.String `tfsdk:"s7" json:"S7,computed"`
	S8  types.String `tfsdk:"s8" json:"S8,computed"`
	S9  types.String `tfsdk:"s9" json:"S9,computed"`
}

type AIGatewaysOtelDataSourceModel struct {
	Headers       customfield.Map[types.String] `tfsdk:"headers" json:"headers,computed"`
	URL           types.String                  `tfsdk:"url" json:"url,computed"`
	Authorization types.String                  `tfsdk:"authorization" json:"authorization,computed"`
	ContentType   types.String                  `tfsdk:"content_type" json:"content_type,computed"`
}

type AIGatewaysSpendLimitsDataSourceModel struct {
	Enabled types.Bool                                                              `tfsdk:"enabled" json:"enabled,computed"`
	Rules   customfield.NestedObjectList[AIGatewaysSpendLimitsRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
}

type AIGatewaysSpendLimitsRulesDataSourceModel struct {
	ID                types.String                                                                   `tfsdk:"id" json:"id,computed"`
	Limit             types.Float64                                                                  `tfsdk:"limit" json:"limit,computed"`
	LimitType         types.String                                                                   `tfsdk:"limit_type" json:"limitType,computed"`
	Window            types.Int64                                                                    `tfsdk:"window" json:"window,computed"`
	Enabled           types.Bool                                                                     `tfsdk:"enabled" json:"enabled,computed"`
	Metadata          customfield.NestedObjectMap[AIGatewaysSpendLimitsRulesMetadataDataSourceModel] `tfsdk:"metadata" json:"metadata,computed"`
	Model             types.String                                                                   `tfsdk:"model" json:"model,computed"`
	AIGatewayProvider types.String                                                                   `tfsdk:"ai_gateway_provider" json:"provider,computed"`
	Technique         types.String                                                                   `tfsdk:"technique" json:"technique,computed"`
}

type AIGatewaysSpendLimitsRulesMetadataDataSourceModel struct {
	Mode  types.String `tfsdk:"mode" json:"mode,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type AIGatewaysStripeDataSourceModel struct {
	Authorization types.String                                                             `tfsdk:"authorization" json:"authorization,computed"`
	UsageEvents   customfield.NestedObjectList[AIGatewaysStripeUsageEventsDataSourceModel] `tfsdk:"usage_events" json:"usage_events,computed"`
}

type AIGatewaysStripeUsageEventsDataSourceModel struct {
	Payload types.String `tfsdk:"payload" json:"payload,computed"`
}
