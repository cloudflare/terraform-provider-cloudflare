// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessShortLivedCertificateResultDataSourceEnvelope struct {
Result ZeroTrustAccessShortLivedCertificateDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessShortLivedCertificateDataSourceModel struct {
AppID types.String `tfsdk:"app_id" path:"app_id,required"`
AccountID types.String `tfsdk:"account_id" path:"account_id,optional"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,optional"`
AUD types.String `tfsdk:"aud" json:"aud,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
PublicKey types.String `tfsdk:"public_key" json:"public_key,computed"`
}

func (m *ZeroTrustAccessShortLivedCertificateDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessApplicationCAGetParams, diags diag.Diagnostics) {
  params = zero_trust.AccessApplicationCAGetParams{

  }

  if !m.AccountID.IsNull() {
    params.AccountID = cloudflare.F(m.AccountID.ValueString())
  } else {
    params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
  }

  return
}
