// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_sensitivity_level_order

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPSensitivityLevelOrderResultEnvelope struct {
	Result ZeroTrustDLPSensitivityLevelOrderModel `json:"result"`
}

type ZeroTrustDLPSensitivityLevelOrderModel struct {
	ID                 types.String    `tfsdk:"id" json:"-,computed"`
	SensitivityGroupID types.String    `tfsdk:"sensitivity_group_id" path:"sensitivity_group_id,required"`
	AccountID          types.String    `tfsdk:"account_id" path:"account_id,required"`
	LevelIDs           *[]types.String `tfsdk:"level_ids" json:"level_ids,required"`
}

func (m ZeroTrustDLPSensitivityLevelOrderModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPSensitivityLevelOrderModel) MarshalJSONForUpdate(state ZeroTrustDLPSensitivityLevelOrderModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
