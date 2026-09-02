// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_namespace

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISearchNamespaceResultEnvelope struct {
	Result AISearchNamespaceModel `json:"result"`
}

type AISearchNamespaceModel struct {
	AccountID            types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	Name                 types.String                                                         `tfsdk:"name" json:"name,required"`
	Description          types.String                                                         `tfsdk:"description" json:"description,optional"`
	PublicEndpointParams customfield.NestedObject[AISearchNamespacePublicEndpointParamsModel] `tfsdk:"public_endpoint_params" json:"public_endpoint_params,computed_optional"`
	CreatedAt            timetypes.RFC3339                                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	PublicEndpointID     types.String                                                         `tfsdk:"public_endpoint_id" json:"public_endpoint_id,computed"`
}

func (m AISearchNamespaceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AISearchNamespaceModel) MarshalJSONForUpdate(state AISearchNamespaceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AISearchNamespacePublicEndpointParamsModel struct {
	AuthorizedHosts         *[]types.String                                                                             `tfsdk:"authorized_hosts" json:"authorized_hosts,optional"`
	ChatCompletionsEndpoint customfield.NestedObject[AISearchNamespacePublicEndpointParamsChatCompletionsEndpointModel] `tfsdk:"chat_completions_endpoint" json:"chat_completions_endpoint,computed_optional"`
	CustomDomains           *[]types.String                                                                             `tfsdk:"custom_domains" json:"custom_domains,optional"`
	DefaultDomainEnabled    types.Bool                                                                                  `tfsdk:"default_domain_enabled" json:"default_domain_enabled,computed_optional"`
	Enabled                 types.Bool                                                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	InstancesAllowed        customfield.List[types.String]                                                              `tfsdk:"instances_allowed" json:"instances_allowed,computed_optional"`
	Mcp                     customfield.NestedObject[AISearchNamespacePublicEndpointParamsMcpModel]                     `tfsdk:"mcp" json:"mcp,computed_optional"`
	RateLimit               *AISearchNamespacePublicEndpointParamsRateLimitModel                                        `tfsdk:"rate_limit" json:"rate_limit,optional"`
	SearchEndpoint          customfield.NestedObject[AISearchNamespacePublicEndpointParamsSearchEndpointModel]          `tfsdk:"search_endpoint" json:"search_endpoint,computed_optional"`
}

type AISearchNamespacePublicEndpointParamsChatCompletionsEndpointModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed_optional"`
}

type AISearchNamespacePublicEndpointParamsMcpModel struct {
	Description types.String `tfsdk:"description" json:"description,computed_optional"`
	Disabled    types.Bool   `tfsdk:"disabled" json:"disabled,computed_optional"`
}

type AISearchNamespacePublicEndpointParamsRateLimitModel struct {
	PeriodMs  types.Int64  `tfsdk:"period_ms" json:"period_ms,optional"`
	Requests  types.Int64  `tfsdk:"requests" json:"requests,optional"`
	Technique types.String `tfsdk:"technique" json:"technique,optional"`
}

type AISearchNamespacePublicEndpointParamsSearchEndpointModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed_optional"`
}
