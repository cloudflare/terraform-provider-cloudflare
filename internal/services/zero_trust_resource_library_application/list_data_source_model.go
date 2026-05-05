// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_application

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustResourceLibraryApplicationsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustResourceLibraryApplicationsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustResourceLibraryApplicationsDataSourceModel struct {
	AccountID types.String                                                                            `tfsdk:"account_id" path:"account_id,required"`
	Filter    types.String                                                                            `tfsdk:"filter" query:"filter,optional"`
	OrderBy   types.String                                                                            `tfsdk:"order_by" query:"order_by,optional"`
	Limit     types.Int64                                                                             `tfsdk:"limit" query:"limit,computed_optional"`
	Offset    types.Int64                                                                             `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems  types.Int64                                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustResourceLibraryApplicationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustResourceLibraryApplicationsDataSourceModel) toListParams(_ context.Context) (params zero_trust.ResourceLibraryApplicationListParams, diags diag.Diagnostics) {
	params = zero_trust.ResourceLibraryApplicationListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.IsNull() {
		params.Filter = cloudflare.F(m.Filter.ValueString())
	}
	if !m.Limit.IsNull() {
		params.Limit = cloudflare.F(m.Limit.ValueInt64())
	}
	if !m.Offset.IsNull() {
		params.Offset = cloudflare.F(m.Offset.ValueInt64())
	}
	if !m.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(m.OrderBy.ValueString())
	}

	return
}

type ZeroTrustResourceLibraryApplicationsResultDataSourceModel struct {
	ID                          types.String                   `tfsdk:"id" json:"id,computed"`
	ApplicationConfidenceScore  types.Float64                  `tfsdk:"application_confidence_score" json:"application_confidence_score,computed"`
	ApplicationSource           types.String                   `tfsdk:"application_source" json:"application_source,computed"`
	ApplicationType             types.String                   `tfsdk:"application_type" json:"application_type,computed"`
	ApplicationTypeDescription  types.String                   `tfsdk:"application_type_description" json:"application_type_description,computed"`
	CreatedAt                   types.String                   `tfsdk:"created_at" json:"created_at,computed"`
	GenAIScore                  types.Float64                  `tfsdk:"gen_ai_score" json:"gen_ai_score,computed"`
	Hostnames                   customfield.List[types.String] `tfsdk:"hostnames" json:"hostnames,computed"`
	HumanID                     types.String                   `tfsdk:"human_id" json:"human_id,computed"`
	IPSubnets                   customfield.List[types.String] `tfsdk:"ip_subnets" json:"ip_subnets,computed"`
	Name                        types.String                   `tfsdk:"name" json:"name,computed"`
	PortProtocols               customfield.List[types.String] `tfsdk:"port_protocols" json:"port_protocols,computed"`
	SupportDomains              customfield.List[types.String] `tfsdk:"support_domains" json:"support_domains,computed"`
	UpdatedAt                   types.String                   `tfsdk:"updated_at" json:"updated_at,computed"`
	Version                     types.String                   `tfsdk:"version" json:"version,computed"`
	ApplicationScoreComposition jsontypes.Normalized           `tfsdk:"application_score_composition" json:"application_score_composition,computed"`
	IntelID                     types.Int64                    `tfsdk:"intel_id" json:"intel_id,computed"`
}
