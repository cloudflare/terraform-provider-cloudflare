// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessShortLivedCertificateResultEnvelope struct {
	Result ZeroTrustAccessShortLivedCertificateModel `json:"result"`
}

type ZeroTrustAccessShortLivedCertificateModel struct {
	ID        types.String `tfsdk:"id" json:"-,computed"`
	AppID     types.String `tfsdk:"app_id" path:"app_id,required"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,optional"`
	AUD       types.String `tfsdk:"aud" json:"aud,computed"`
	PublicKey types.String `tfsdk:"public_key" json:"public_key,computed"`
}

func (m ZeroTrustAccessShortLivedCertificateModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessShortLivedCertificateModel) MarshalJSONForUpdate(state ZeroTrustAccessShortLivedCertificateModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
