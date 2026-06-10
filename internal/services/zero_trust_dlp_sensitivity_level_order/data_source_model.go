// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_sensitivity_level_order

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPSensitivityLevelOrderResultDataSourceEnvelope struct {
	Result ZeroTrustDLPSensitivityLevelOrderDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPSensitivityLevelOrderDataSourceModel struct {
	ID                 types.String                   `tfsdk:"id" path:"sensitivity_group_id,computed"`
	SensitivityGroupID types.String                   `tfsdk:"sensitivity_group_id" path:"sensitivity_group_id,required"`
	AccountID          types.String                   `tfsdk:"account_id" path:"account_id,required"`
	LevelIDs           customfield.List[types.String] `tfsdk:"level_ids" json:"level_ids,computed"`
}

func (m *ZeroTrustDLPSensitivityLevelOrderDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPSensitivityGroupLevelOrderGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPSensitivityGroupLevelOrderGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
