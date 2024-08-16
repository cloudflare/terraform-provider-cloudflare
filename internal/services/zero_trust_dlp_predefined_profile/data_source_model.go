// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPPredefinedProfileResultDataSourceEnvelope struct {
	Result ZeroTrustDLPPredefinedProfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPPredefinedProfileDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ProfileID types.String `tfsdk:"profile_id" path:"profile_id"`
}
