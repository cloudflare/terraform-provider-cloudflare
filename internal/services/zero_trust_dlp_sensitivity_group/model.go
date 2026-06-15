// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_sensitivity_group

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPSensitivityGroupResultEnvelope struct {
	Result ZeroTrustDLPSensitivityGroupModel `json:"result"`
}

type ZeroTrustDLPSensitivityGroupModel struct {
	ID          types.String                                                          `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                                                          `tfsdk:"account_id" path:"account_id,required"`
	Name        types.String                                                          `tfsdk:"name" json:"name,required"`
	Description types.String                                                          `tfsdk:"description" json:"description,optional"`
	CreatedAt   timetypes.RFC3339                                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	TemplateID  types.String                                                          `tfsdk:"template_id" json:"template_id,computed"`
	UpdatedAt   timetypes.RFC3339                                                     `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Levels      customfield.NestedObjectList[ZeroTrustDLPSensitivityGroupLevelsModel] `tfsdk:"levels" json:"levels,computed"`
}

func (m ZeroTrustDLPSensitivityGroupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPSensitivityGroupModel) MarshalJSONForUpdate(state ZeroTrustDLPSensitivityGroupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDLPSensitivityGroupLevelsModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
}
