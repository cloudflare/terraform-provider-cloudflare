// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessShortLivedCertificateResultDataSourceEnvelope struct {
	Result ZeroTrustAccessShortLivedCertificateDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessShortLivedCertificateResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessShortLivedCertificateDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessShortLivedCertificateDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	AppID     types.String `tfsdk:"app_id" path:"app_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
	AUD       types.String `tfsdk:"aud" json:"aud"`
	ID        types.String `tfsdk:"id" json:"id"`
	PublicKey types.String `tfsdk:"public_key" json:"public_key"`
}
