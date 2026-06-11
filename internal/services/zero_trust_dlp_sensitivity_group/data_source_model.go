// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_sensitivity_group

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPSensitivityGroupResultDataSourceEnvelope struct {
	Result ZeroTrustDLPSensitivityGroupDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPSensitivityGroupDataSourceModel struct {
	ID                 types.String                                                                    `tfsdk:"id" path:"sensitivity_group_id,computed"`
	SensitivityGroupID types.String                                                                    `tfsdk:"sensitivity_group_id" path:"sensitivity_group_id,required"`
	AccountID          types.String                                                                    `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt          timetypes.RFC3339                                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description        types.String                                                                    `tfsdk:"description" json:"description,computed"`
	Name               types.String                                                                    `tfsdk:"name" json:"name,computed"`
	TemplateID         types.String                                                                    `tfsdk:"template_id" json:"template_id,computed"`
	UpdatedAt          timetypes.RFC3339                                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Levels             customfield.NestedObjectList[ZeroTrustDLPSensitivityGroupLevelsDataSourceModel] `tfsdk:"levels" json:"levels,computed"`
}

func (m *ZeroTrustDLPSensitivityGroupDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPSensitivityGroupGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPSensitivityGroupGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPSensitivityGroupLevelsDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
}
