// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_data_class

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPDataClassResultEnvelope struct {
	Result ZeroTrustDLPDataClassModel `json:"result"`
}

type ZeroTrustDLPDataClassModel struct {
	ID                types.String                                    `tfsdk:"id" json:"id,computed"`
	AccountID         types.String                                    `tfsdk:"account_id" path:"account_id,required"`
	Expression        types.String                                    `tfsdk:"expression" json:"expression,required"`
	Name              types.String                                    `tfsdk:"name" json:"name,required"`
	DataTags          *[]types.String                                 `tfsdk:"data_tags" json:"data_tags,required"`
	SensitivityLevels *[]*ZeroTrustDLPDataClassSensitivityLevelsModel `tfsdk:"sensitivity_levels" json:"sensitivity_levels,required"`
	Description       types.String                                    `tfsdk:"description" json:"description,optional"`
	CreatedAt         timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt         timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustDLPDataClassModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPDataClassModel) MarshalJSONForUpdate(state ZeroTrustDLPDataClassModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDLPDataClassSensitivityLevelsModel struct {
	GroupID types.String `tfsdk:"group_id" json:"group_id,required"`
	LevelID types.String `tfsdk:"level_id" json:"level_id,required"`
}
