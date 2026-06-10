// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_sensitivity_level

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPSensitivityLevelResultDataSourceEnvelope struct {
	Result ZeroTrustDLPSensitivityLevelDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPSensitivityLevelDataSourceModel struct {
	ID                 types.String      `tfsdk:"id" path:"sensitivity_level_id,computed"`
	SensitivityLevelID types.String      `tfsdk:"sensitivity_level_id" path:"sensitivity_level_id,required"`
	AccountID          types.String      `tfsdk:"account_id" path:"account_id,required"`
	SensitivityGroupID types.String      `tfsdk:"sensitivity_group_id" path:"sensitivity_group_id,required"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description        types.String      `tfsdk:"description" json:"description,computed"`
	Name               types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt          timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m *ZeroTrustDLPSensitivityLevelDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPSensitivityGroupLevelGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPSensitivityGroupLevelGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
