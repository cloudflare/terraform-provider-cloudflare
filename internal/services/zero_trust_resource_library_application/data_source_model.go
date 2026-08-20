// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_application

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustResourceLibraryApplicationResultDataSourceEnvelope struct {
	Result ZeroTrustResourceLibraryApplicationDataSourceModel `json:"result,computed"`
}

type ZeroTrustResourceLibraryApplicationDataSourceModel struct {
	ID                          types.Int64                                                  `tfsdk:"id" path:"id,computed_optional"`
	AccountID                   types.String                                                 `tfsdk:"account_id" path:"account_id,required"`
	ApplicationConfidenceScore  types.Float64                                                `tfsdk:"application_confidence_score" json:"application_confidence_score,computed"`
	ApplicationSource           types.String                                                 `tfsdk:"application_source" json:"application_source,computed"`
	ApplicationType             types.String                                                 `tfsdk:"application_type" json:"application_type,computed"`
	ApplicationTypeDescription  types.String                                                 `tfsdk:"application_type_description" json:"application_type_description,computed"`
	CategoryID                  types.Int64                                                  `tfsdk:"category_id" json:"category_id,computed"`
	CreatedAt                   types.String                                                 `tfsdk:"created_at" json:"created_at,computed"`
	GenAIScore                  types.Float64                                                `tfsdk:"gen_ai_score" json:"gen_ai_score,computed"`
	HumanID                     types.String                                                 `tfsdk:"human_id" json:"human_id,computed"`
	Name                        types.String                                                 `tfsdk:"name" json:"name,computed"`
	UpdatedAt                   types.String                                                 `tfsdk:"updated_at" json:"updated_at,computed"`
	Version                     types.String                                                 `tfsdk:"version" json:"version,computed"`
	Hostnames                   customfield.Set[types.String]                                `tfsdk:"hostnames" json:"hostnames,computed"`
	IPSubnets                   customfield.Set[types.String]                                `tfsdk:"ip_subnets" json:"ip_subnets,computed"`
	PortProtocols               customfield.Set[types.String]                                `tfsdk:"port_protocols" json:"port_protocols,computed"`
	SupportDomains              customfield.Set[types.String]                                `tfsdk:"support_domains" json:"support_domains,computed"`
	Supported                   customfield.Set[types.String]                                `tfsdk:"supported" json:"supported,computed"`
	ApplicationScoreComposition jsontypes.Normalized                                         `tfsdk:"application_score_composition" json:"application_score_composition,computed"`
	Filter                      *ZeroTrustResourceLibraryApplicationFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustResourceLibraryApplicationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.ResourceLibraryApplicationGetParams, diags diag.Diagnostics) {
	params = zero_trust.ResourceLibraryApplicationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustResourceLibraryApplicationDataSourceModel) toListParams(_ context.Context) (params zero_trust.ResourceLibraryApplicationListParams, diags diag.Diagnostics) {
	params = zero_trust.ResourceLibraryApplicationListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Filter.IsNull() {
		params.Filter = cloudflare.F(m.Filter.Filter.ValueString())
	}
	if !m.Filter.Limit.IsNull() {
		params.Limit = cloudflare.F(m.Filter.Limit.ValueInt64())
	}
	if !m.Filter.Offset.IsNull() {
		params.Offset = cloudflare.F(m.Filter.Offset.ValueInt64())
	}
	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(m.Filter.OrderBy.ValueString())
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type ZeroTrustResourceLibraryApplicationFindOneByDataSourceModel struct {
	Filter  types.String `tfsdk:"filter" query:"filter,optional"`
	Limit   types.Int64  `tfsdk:"limit" query:"limit,computed_optional"`
	Offset  types.Int64  `tfsdk:"offset" query:"offset,computed_optional"`
	OrderBy types.String `tfsdk:"order_by" query:"order_by,optional"`
	Search  types.String `tfsdk:"search" query:"search,optional"`
}
