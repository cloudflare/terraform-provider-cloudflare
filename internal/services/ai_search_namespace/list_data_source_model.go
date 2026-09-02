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

type AISearchNamespacesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AISearchNamespacesResultDataSourceModel] `json:"result,computed"`
}

type AISearchNamespacesDataSourceModel struct {
	AccountID types.String                                                          `tfsdk:"account_id" path:"account_id,required"`
	Search    types.String                                                          `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                           `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AISearchNamespacesResultDataSourceModel] `tfsdk:"result"`
}

func (m *AISearchNamespacesDataSourceModel) toListParams(_ context.Context) (params ai_search.NamespaceListParams, diags diag.Diagnostics) {
	params = ai_search.NamespaceListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type AISearchNamespacesResultDataSourceModel struct {
	CreatedAt            timetypes.RFC3339                                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name                 types.String                                                                    `tfsdk:"name" json:"name,computed"`
	Description          types.String                                                                    `tfsdk:"description" json:"description,computed"`
	PublicEndpointID     types.String                                                                    `tfsdk:"public_endpoint_id" json:"public_endpoint_id,computed"`
	PublicEndpointParams customfield.NestedObject[AISearchNamespacesPublicEndpointParamsDataSourceModel] `tfsdk:"public_endpoint_params" json:"public_endpoint_params,computed"`
}

type AISearchNamespacesPublicEndpointParamsDataSourceModel struct {
	AuthorizedHosts         customfield.List[types.String]                                                                         `tfsdk:"authorized_hosts" json:"authorized_hosts,computed"`
	ChatCompletionsEndpoint customfield.NestedObject[AISearchNamespacesPublicEndpointParamsChatCompletionsEndpointDataSourceModel] `tfsdk:"chat_completions_endpoint" json:"chat_completions_endpoint,computed"`
	CustomDomains           customfield.List[types.String]                                                                         `tfsdk:"custom_domains" json:"custom_domains,computed"`
	DefaultDomainEnabled    types.Bool                                                                                             `tfsdk:"default_domain_enabled" json:"default_domain_enabled,computed"`
	Enabled                 types.Bool                                                                                             `tfsdk:"enabled" json:"enabled,computed"`
	InstancesAllowed        customfield.List[types.String]                                                                         `tfsdk:"instances_allowed" json:"instances_allowed,computed"`
	Mcp                     customfield.NestedObject[AISearchNamespacesPublicEndpointParamsMcpDataSourceModel]                     `tfsdk:"mcp" json:"mcp,computed"`
	RateLimit               customfield.NestedObject[AISearchNamespacesPublicEndpointParamsRateLimitDataSourceModel]               `tfsdk:"rate_limit" json:"rate_limit,computed"`
	SearchEndpoint          customfield.NestedObject[AISearchNamespacesPublicEndpointParamsSearchEndpointDataSourceModel]          `tfsdk:"search_endpoint" json:"search_endpoint,computed"`
}

type AISearchNamespacesPublicEndpointParamsChatCompletionsEndpointDataSourceModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed"`
}

type AISearchNamespacesPublicEndpointParamsMcpDataSourceModel struct {
	Description types.String `tfsdk:"description" json:"description,computed"`
	Disabled    types.Bool   `tfsdk:"disabled" json:"disabled,computed"`
}

type AISearchNamespacesPublicEndpointParamsRateLimitDataSourceModel struct {
	PeriodMs  types.Int64  `tfsdk:"period_ms" json:"period_ms,computed"`
	Requests  types.Int64  `tfsdk:"requests" json:"requests,computed"`
	Technique types.String `tfsdk:"technique" json:"technique,computed"`
}

type AISearchNamespacesPublicEndpointParamsSearchEndpointDataSourceModel struct {
	Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed"`
}
