// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_application

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustResourceLibraryApplicationResultEnvelope struct {
	Result ZeroTrustResourceLibraryApplicationModel `json:"result"`
}

type ZeroTrustResourceLibraryApplicationModel struct {
	ID                          types.Int64                   `tfsdk:"id" json:"id,computed"`
	AccountID                   types.String                  `tfsdk:"account_id" path:"account_id,required"`
	CategoryID                  types.Int64                   `tfsdk:"category_id" json:"category_id,required"`
	HumanID                     types.String                  `tfsdk:"human_id" json:"human_id,required"`
	Name                        types.String                  `tfsdk:"name" json:"name,required"`
	Hostnames                   *[]types.String               `tfsdk:"hostnames" json:"hostnames,optional"`
	IPSubnets                   *[]types.String               `tfsdk:"ip_subnets" json:"ip_subnets,optional"`
	PortProtocols               *[]types.String               `tfsdk:"port_protocols" json:"port_protocols,optional"`
	SupportDomains              *[]types.String               `tfsdk:"support_domains" json:"support_domains,optional"`
	ApplicationConfidenceScore  types.Float64                 `tfsdk:"application_confidence_score" json:"application_confidence_score,computed"`
	ApplicationSource           types.String                  `tfsdk:"application_source" json:"application_source,computed"`
	ApplicationType             types.String                  `tfsdk:"application_type" json:"application_type,computed"`
	ApplicationTypeDescription  types.String                  `tfsdk:"application_type_description" json:"application_type_description,computed"`
	CreatedAt                   types.String                  `tfsdk:"created_at" json:"created_at,computed"`
	GenAIScore                  types.Float64                 `tfsdk:"gen_ai_score" json:"gen_ai_score,computed"`
	UpdatedAt                   types.String                  `tfsdk:"updated_at" json:"updated_at,computed"`
	Version                     types.String                  `tfsdk:"version" json:"version,computed"`
	Supported                   customfield.Set[types.String] `tfsdk:"supported" json:"supported,computed"`
	ApplicationScoreComposition jsontypes.Normalized          `tfsdk:"application_score_composition" json:"application_score_composition,computed"`
}

func (m ZeroTrustResourceLibraryApplicationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustResourceLibraryApplicationModel) MarshalJSONForUpdate(state ZeroTrustResourceLibraryApplicationModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
