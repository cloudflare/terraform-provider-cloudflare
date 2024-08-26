// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessShortLivedCertificateResultEnvelope struct {
	Result ZeroTrustAccessShortLivedCertificateModel `json:"result"`
}

type ZeroTrustAccessShortLivedCertificateModel struct {
	ID        types.String `tfsdk:"id" json:"-,computed"`
	AppID     types.String `tfsdk:"app_id" path:"app_id"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
