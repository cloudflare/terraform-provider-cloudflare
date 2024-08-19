// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessShortLivedCertificatesResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessShortLivedCertificatesResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessShortLivedCertificatesDataSourceModel struct {
	AccountID types.String                                                   `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                                   `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                                    `tfsdk:"max_items"`
	Result    *[]*ZeroTrustAccessShortLivedCertificatesResultDataSourceModel `tfsdk:"result"`
}

type ZeroTrustAccessShortLivedCertificatesResultDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AUD       types.String `tfsdk:"aud" json:"aud,computed"`
	PublicKey types.String `tfsdk:"public_key" json:"public_key,computed"`
}
