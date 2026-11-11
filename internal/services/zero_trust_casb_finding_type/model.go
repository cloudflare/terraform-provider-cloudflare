// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_finding_type

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCasbFindingTypeResultEnvelope struct {
	Result ZeroTrustCasbFindingTypeModel `json:"result"`
}

type ZeroTrustCasbFindingTypeModel struct {
	ID                       types.String                                                    `tfsdk:"id" json:"id,computed"`
	AccountID                types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	AssetClass               types.String                                                    `tfsdk:"asset_class" json:"asset_class,required"`
	AssetVendor              types.String                                                    `tfsdk:"asset_vendor" json:"asset_vendor,required"`
	Expression               types.String                                                    `tfsdk:"expression" json:"expression,required"`
	AppliesToAllIntegrations types.Bool                                                      `tfsdk:"applies_to_all_integrations" json:"applies_to_all_integrations,required"`
	Description              types.String                                                    `tfsdk:"description" json:"description,required"`
	Name                     types.String                                                    `tfsdk:"name" json:"name,required"`
	Severity                 types.String                                                    `tfsdk:"severity" json:"severity,required"`
	IntegrationIDs           customfield.List[types.String]                                  `tfsdk:"integration_ids" json:"integration_ids,computed_optional"`
	CreatedAt                timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	MatchCount               types.Int64                                                     `tfsdk:"match_count" json:"match_count,computed"`
	Origin                   types.String                                                    `tfsdk:"origin" json:"origin,computed"`
	Quarantined              types.Bool                                                      `tfsdk:"quarantined" json:"quarantined,computed"`
	UpdatedAt                timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Vendor                   types.String                                                    `tfsdk:"vendor" json:"vendor,computed"`
	Version                  types.Int64                                                     `tfsdk:"version" json:"version,computed"`
	Category                 customfield.NestedObject[ZeroTrustCasbFindingTypeCategoryModel] `tfsdk:"category" json:"category,computed"`
}

func (m ZeroTrustCasbFindingTypeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCasbFindingTypeModel) MarshalJSONForUpdate(state ZeroTrustCasbFindingTypeModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustCasbFindingTypeCategoryModel struct {
	Observation types.String `tfsdk:"observation" json:"observation,computed"`
	Product     types.String `tfsdk:"product" json:"product,computed"`
	Type        types.String `tfsdk:"type" json:"type,computed"`
}
