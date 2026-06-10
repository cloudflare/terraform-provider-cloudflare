// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_sensitivity_level

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPSensitivityLevelsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDLPSensitivityLevelsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDLPSensitivityLevelsDataSourceModel struct {
	AccountID          types.String                                                                     `tfsdk:"account_id" path:"account_id,required"`
	SensitivityGroupID types.String                                                                     `tfsdk:"sensitivity_group_id" path:"sensitivity_group_id,required"`
	MaxItems           types.Int64                                                                      `tfsdk:"max_items"`
	Result             customfield.NestedObjectList[ZeroTrustDLPSensitivityLevelsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDLPSensitivityLevelsDataSourceModel) toListParams(_ context.Context) (params zero_trust.DLPSensitivityGroupLevelListParams, diags diag.Diagnostics) {
	params = zero_trust.DLPSensitivityGroupLevelListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPSensitivityLevelsResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
}
