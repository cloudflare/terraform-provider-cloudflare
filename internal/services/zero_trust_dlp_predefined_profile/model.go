// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPPredefinedProfileResultEnvelope struct {
	Result ZeroTrustDLPPredefinedProfileModel `json:"result"`
}

type ZeroTrustDLPPredefinedProfileModel struct {
	ID        types.String `tfsdk:"id" json:"-,computed"`
	ProfileID types.String `tfsdk:"profile_id" path:"profile_id"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
