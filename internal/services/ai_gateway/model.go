// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AIGatewayResultEnvelope struct {
	Result AIGatewayModel `json:"result"`
}

type AIGatewayModel struct {
	ID                      types.String                                     `tfsdk:"id" json:"id,required"`
	AccountID               types.String                                     `tfsdk:"account_id" path:"account_id,required"`
	CacheInvalidateOnUpdate types.Bool                                       `tfsdk:"cache_invalidate_on_update" json:"cache_invalidate_on_update,required"`
	CacheTTL                types.Int64                                      `tfsdk:"cache_ttl" json:"cache_ttl,required"`
	CollectLogs             types.Bool                                       `tfsdk:"collect_logs" json:"collect_logs,required"`
	RateLimitingInterval    types.Int64                                      `tfsdk:"rate_limiting_interval" json:"rate_limiting_interval,required"`
	RateLimitingLimit       types.Int64                                      `tfsdk:"rate_limiting_limit" json:"rate_limiting_limit,required"`
	Authentication          types.Bool                                       `tfsdk:"authentication" json:"authentication,optional"`
	LogManagement           types.Int64                                      `tfsdk:"log_management" json:"log_management,optional"`
	LogManagementStrategy   types.String                                     `tfsdk:"log_management_strategy" json:"log_management_strategy,optional"`
	Logpush                 types.Bool                                       `tfsdk:"logpush" json:"logpush,optional"`
	LogpushPublicKey        types.String                                     `tfsdk:"logpush_public_key" json:"logpush_public_key,optional"`
	RateLimitingTechnique   types.String                                     `tfsdk:"rate_limiting_technique" json:"rate_limiting_technique,optional"`
	RetryBackoff            types.String                                     `tfsdk:"retry_backoff" json:"retry_backoff,optional"`
	RetryDelay              types.Int64                                      `tfsdk:"retry_delay" json:"retry_delay,optional"`
	RetryMaxAttempts        types.Int64                                      `tfsdk:"retry_max_attempts" json:"retry_max_attempts,optional"`
	StoreID                 types.String                                     `tfsdk:"store_id" json:"store_id,optional"`
	Zdr                     types.Bool                                       `tfsdk:"zdr" json:"zdr,optional"`
	DLP                     *AIGatewayDLPModel                               `tfsdk:"dlp" json:"dlp,optional"`
	Stripe                  *AIGatewayStripeModel                            `tfsdk:"stripe" json:"stripe,optional"`
	WorkersAIBillingMode    types.String                                     `tfsdk:"workers_ai_billing_mode" json:"workers_ai_billing_mode,computed_optional"`
	Otel                    customfield.NestedObjectList[AIGatewayOtelModel] `tfsdk:"otel" json:"otel,computed_optional"`
	CreatedAt               timetypes.RFC3339                                `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IsDefault               types.Bool                                       `tfsdk:"is_default" json:"is_default,computed"`
	ModifiedAt              timetypes.RFC3339                                `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
}

func (m AIGatewayModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AIGatewayModel) MarshalJSONForUpdate(state AIGatewayModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AIGatewayDLPModel struct {
	Action   types.String                  `tfsdk:"action" json:"action,optional"`
	Enabled  types.Bool                    `tfsdk:"enabled" json:"enabled,required"`
	Profiles *[]types.String               `tfsdk:"profiles" json:"profiles,optional"`
	Policies *[]*AIGatewayDLPPoliciesModel `tfsdk:"policies" json:"policies,optional"`
}

type AIGatewayDLPPoliciesModel struct {
	ID       types.String    `tfsdk:"id" json:"id,required"`
	Action   types.String    `tfsdk:"action" json:"action,required"`
	Check    *[]types.String `tfsdk:"check" json:"check,required"`
	Enabled  types.Bool      `tfsdk:"enabled" json:"enabled,required"`
	Profiles *[]types.String `tfsdk:"profiles" json:"profiles,required"`
}

type AIGatewayStripeModel struct {
	Authorization types.String                        `tfsdk:"authorization" json:"authorization,required"`
	UsageEvents   *[]*AIGatewayStripeUsageEventsModel `tfsdk:"usage_events" json:"usage_events,required"`
}

type AIGatewayStripeUsageEventsModel struct {
	Payload types.String `tfsdk:"payload" json:"payload,required"`
}

type AIGatewayOtelModel struct {
	Authorization types.String             `tfsdk:"authorization" json:"authorization,required"`
	Headers       *map[string]types.String `tfsdk:"headers" json:"headers,required"`
	URL           types.String             `tfsdk:"url" json:"url,required"`
	ContentType   types.String             `tfsdk:"content_type" json:"content_type,computed_optional"`
}
