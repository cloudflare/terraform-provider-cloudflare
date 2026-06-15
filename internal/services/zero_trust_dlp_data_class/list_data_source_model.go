// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_data_class

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPDataClassesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDLPDataClassesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDLPDataClassesDataSourceModel struct {
	AccountID types.String                                                               `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDLPDataClassesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDLPDataClassesDataSourceModel) toListParams(_ context.Context) (params zero_trust.DLPDataClassListParams, diags diag.Diagnostics) {
	params = zero_trust.DLPDataClassListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPDataClassesResultDataSourceModel struct {
	ID                types.String                                                                          `tfsdk:"id" json:"id,computed"`
	CreatedAt         timetypes.RFC3339                                                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DataTags          customfield.List[types.String]                                                        `tfsdk:"data_tags" json:"data_tags,computed"`
	Expression        types.String                                                                          `tfsdk:"expression" json:"expression,computed"`
	Name              types.String                                                                          `tfsdk:"name" json:"name,computed"`
	SensitivityLevels customfield.NestedObjectList[ZeroTrustDLPDataClassesSensitivityLevelsDataSourceModel] `tfsdk:"sensitivity_levels" json:"sensitivity_levels,computed"`
	UpdatedAt         timetypes.RFC3339                                                                     `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description       types.String                                                                          `tfsdk:"description" json:"description,computed"`
}

type ZeroTrustDLPDataClassesSensitivityLevelsDataSourceModel struct {
	GroupID types.String `tfsdk:"group_id" json:"group_id,computed"`
	LevelID types.String `tfsdk:"level_id" json:"level_id,computed"`
}
