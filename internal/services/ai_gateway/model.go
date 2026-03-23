// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AIGatewayResultEnvelope struct {
	Result AIGatewayModel `json:"result"`
}

type AIGatewayModel struct {
	ID                          types.String                                                `tfsdk:"id" json:"id,required"`
	AccountID                   types.String                                                `tfsdk:"account_id" path:"account_id,required"`
	CacheInvalidateOnUpdate      types.Bool                                                  `tfsdk:"cache_invalidate_on_update" json:"cache_invalidate_on_update,optional"`
	CacheTTL                    types.Int64                                                 `tfsdk:"cache_ttl" json:"cache_ttl,optional"`
	CollectLogs                 types.Bool                                                  `tfsdk:"collect_logs" json:"collect_logs,optional"`
	CreatedAt                   timetypes.RFC3339                                           `tfsdk:"created_at" json:"created_at,computed"`
	ModifiedAt                  timetypes.RFC3339                                           `tfsdk:"modified_at" json:"modified_at,computed"`
	RateLimitingInterval        types.Int64                                                 `tfsdk:"rate_limiting_interval" json:"rate_limiting_interval,optional"`
	RateLimitingLimit           types.Int64                                                 `tfsdk:"rate_limiting_limit" json:"rate_limiting_limit,optional"`
	Authentication              types.Bool                                                  `tfsdk:"authentication" json:"authentication,optional"`
	DLP                         *AIGatewayDLPModel                                         `tfsdk:"dlp" json:"dlp,optional"`
	IsDefault                   types.Bool                                                  `tfsdk:"is_default" json:"is_default,computed_optional"`
	LogManagement               types.Int64                                                 `tfsdk:"log_management" json:"log_management,optional"`
	LogManagementStrategy       types.String                                                `tfsdk:"log_management_strategy" json:"log_management_strategy,optional"`
	Logpush                     types.Bool                                                  `tfsdk:"logpush" json:"logpush,optional"`
	LogpushPublicKey            types.String                                                `tfsdk:"logpush_public_key" json:"logpush_public_key,optional"`
	OTel                        []AIGatewayOTelModel                                        `tfsdk:"otel" json:"otel,optional"`
	RateLimitingTechnique       types.String                                                `tfsdk:"rate_limiting_technique" json:"rate_limiting_technique,optional"`
	StoreID                     types.String                                                `tfsdk:"store_id" json:"store_id,optional"`
	Stripe                      *AIGatewayStripeModel                                      `tfsdk:"stripe" json:"stripe,optional"`
	WorkersAIBillingMode        types.String                                                `tfsdk:"workers_ai_billing_mode" json:"workers_ai_billing_mode,optional"`
	ZDR                         types.Bool                                                  `tfsdk:"zdr" json:"zdr,optional"`
}

func (m AIGatewayModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AIGatewayModel) MarshalJSONForUpdate(state AIGatewayModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AIGatewayDLPModel struct {
	Action   types.String  `tfsdk:"action" json:"action,optional"`
	Enabled  types.Bool    `tfsdk:"enabled" json:"enabled,optional"`
	Profiles []types.String `tfsdk:"profiles" json:"profiles,optional"`
}

type AIGatewayOTelModel struct {
	Authorization types.String             `tfsdk:"authorization" json:"authorization,optional"`
	Headers       *map[string]types.String `tfsdk:"headers" json:"headers,optional"`
	URL          types.String              `tfsdk:"url" json:"url,optional"`
	ContentType  types.String              `tfsdk:"content_type" json:"content_type,optional"`
}

type AIGatewayStripeModel struct {
	Authorization types.String                     `tfsdk:"authorization" json:"authorization,optional"`
	UsageEvents   []AIGatewayStripeUsageEventModel `tfsdk:"usage_events" json:"usage_events,optional"`
}

type AIGatewayStripeUsageEventModel struct {
	Payload types.String `tfsdk:"payload" json:"payload,optional"`
}
