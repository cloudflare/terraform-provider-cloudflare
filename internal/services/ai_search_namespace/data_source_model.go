// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_namespace

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_search"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISearchNamespaceResultDataSourceEnvelope struct {
	Result AISearchNamespaceDataSourceModel `json:"result,computed"`
}

type AISearchNamespaceDataSourceModel struct {
	AccountID            types.String                                                                   `tfsdk:"account_id" path:"account_id,required"`
	Name                 types.String                                                                   `tfsdk:"name" path:"name,required"`
	CreatedAt            timetypes.RFC3339                                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description          types.String                                                                   `tfsdk:"description" json:"description,computed"`
	PublicEndpointID     types.String                                                                   `tfsdk:"public_endpoint_id" json:"public_endpoint_id,computed"`
	PublicEndpointParams customfield.NestedObject[AISearchNamespacePublicEndpointParamsDataSourceModel] `tfsdk:"public_endpoint_params" json:"public_endpoint_params,computed"`
}

func (m *AISearchNamespaceDataSourceModel) toReadParams(_ context.Context) (params ai_search.NamespaceReadParams, diags diag.Diagnostics) {
	params = ai_search.NamespaceReadParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type AISearchNamespacePublicEndpointParamsDataSourceModel struct {
	AuthorizedHosts         customfield.List[types.String]                                                                        `tfsdk:"authorized_hosts" json:"authorized_hosts,computed"`
	ChatCompletionsEndpoint customfield.NestedObject[AISearchNamespacePublicEndpointParamsChatCompletionsEndpointDataSourceModel] `tfsdk:"chat_completions_endpoint" json:"chat_completions_endpoint,computed"`
	CustomDomains           customfield.List[types.String]                                                                        `tfsdk:"custom_domains" json:"custom_domains,computed"`
	DefaultDomainEnabled    types.Bool                                                                                            `tfsdk:"default_domain_enabled" json:"default_domain_enabled,computed"`
	Enabled                 types.Bool                                                                                            `tfsdk:"enabled" json:"enabled,computed"`
	InstancesAllowed        customfield.List[types.String]                                                                        `tfsdk:"instances_allowed" json:"instances_allowed,computed"`
	Mcp                     customfield.NestedObject[AISearchNamespacePublicEndpointParamsMcpDataSourceModel]                     `tfsdk:"mcp" json:"mcp,computed"`
	RateLimit               customfield.NestedObject[AISearchNamespacePublicEndpointParamsRateLimitDataSourceModel]               `tfsdk:"rate_limit" json:"rate_limit,computed"`
	SearchEndpoint          customfield.NestedObject[AISearchNamespacePublicEndpointParamsSearchEndpointDataSourceModel]          `tfsdk:"search_endpoint" json:"search_endpoint,computed"`
}

type AISearchNamespacePublicEndpointParamsChatCompletionsEndpointDataSourceModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed"`
}

type AISearchNamespacePublicEndpointParamsMcpDataSourceModel struct {
	Description types.String `tfsdk:"description" json:"description,computed"`
	Disabled    types.Bool   `tfsdk:"disabled" json:"disabled,computed"`
}

type AISearchNamespacePublicEndpointParamsRateLimitDataSourceModel struct {
	PeriodMs  types.Int64  `tfsdk:"period_ms" json:"period_ms,computed"`
	Requests  types.Int64  `tfsdk:"requests" json:"requests,computed"`
	Technique types.String `tfsdk:"technique" json:"technique,computed"`
}

type AISearchNamespacePublicEndpointParamsSearchEndpointDataSourceModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed"`
}
