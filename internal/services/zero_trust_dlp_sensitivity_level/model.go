// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_sensitivity_level

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPSensitivityLevelResultEnvelope struct {
	Result ZeroTrustDLPSensitivityLevelModel `json:"result"`
}

type ZeroTrustDLPSensitivityLevelModel struct {
	ID                 types.String      `tfsdk:"id" json:"id,computed"`
	AccountID          types.String      `tfsdk:"account_id" path:"account_id,required"`
	SensitivityGroupID types.String      `tfsdk:"sensitivity_group_id" path:"sensitivity_group_id,required"`
	Name               types.String      `tfsdk:"name" json:"name,required"`
	Description        types.String      `tfsdk:"description" json:"description,optional"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt          timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustDLPSensitivityLevelModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPSensitivityLevelModel) MarshalJSONForUpdate(state ZeroTrustDLPSensitivityLevelModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
