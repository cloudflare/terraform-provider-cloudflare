// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPPredefinedProfileResultDataSourceEnvelope struct {
	Result ZeroTrustDLPPredefinedProfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPPredefinedProfileDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ProfileID types.String `tfsdk:"profile_id" path:"profile_id"`
}

func (m *ZeroTrustDLPPredefinedProfileDataSourceModel) toReadParams() (params zero_trust.DLPProfilePredefinedGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPProfilePredefinedGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
