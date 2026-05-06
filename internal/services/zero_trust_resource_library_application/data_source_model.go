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

type ZeroTrustResourceLibraryApplicationResultDataSourceEnvelope struct {
	Result ZeroTrustResourceLibraryApplicationDataSourceModel `json:"result,computed"`
}

type ZeroTrustResourceLibraryApplicationDataSourceModel struct {
	AccountID                   types.String                   `tfsdk:"account_id" path:"account_id,required"`
	ID                          types.String                   `tfsdk:"id" path:"id,required"`
	ApplicationConfidenceScore  types.Float64                  `tfsdk:"application_confidence_score" json:"application_confidence_score,computed"`
	ApplicationSource           types.String                   `tfsdk:"application_source" json:"application_source,computed"`
	ApplicationType             types.String                   `tfsdk:"application_type" json:"application_type,computed"`
	ApplicationTypeDescription  types.String                   `tfsdk:"application_type_description" json:"application_type_description,computed"`
	CreatedAt                   types.String                   `tfsdk:"created_at" json:"created_at,computed"`
	GenAIScore                  types.Float64                  `tfsdk:"gen_ai_score" json:"gen_ai_score,computed"`
	HumanID                     types.String                   `tfsdk:"human_id" json:"human_id,computed"`
	IntelID                     types.Int64                    `tfsdk:"intel_id" json:"intel_id,computed"`
	Name                        types.String                   `tfsdk:"name" json:"name,computed"`
	UpdatedAt                   types.String                   `tfsdk:"updated_at" json:"updated_at,computed"`
	Version                     types.String                   `tfsdk:"version" json:"version,computed"`
	Hostnames                   customfield.List[types.String] `tfsdk:"hostnames" json:"hostnames,computed"`
	IPSubnets                   customfield.List[types.String] `tfsdk:"ip_subnets" json:"ip_subnets,computed"`
	PortProtocols               customfield.List[types.String] `tfsdk:"port_protocols" json:"port_protocols,computed"`
	SupportDomains              customfield.List[types.String] `tfsdk:"support_domains" json:"support_domains,computed"`
	Supported                   customfield.List[types.String] `tfsdk:"supported" json:"supported,computed"`
	ApplicationScoreComposition jsontypes.Normalized           `tfsdk:"application_score_composition" json:"application_score_composition,computed"`
}

func (m *ZeroTrustResourceLibraryApplicationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.ResourceLibraryApplicationGetParams, diags diag.Diagnostics) {
	params = zero_trust.ResourceLibraryApplicationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
