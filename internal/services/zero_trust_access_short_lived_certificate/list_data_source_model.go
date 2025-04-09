// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessShortLivedCertificatesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessShortLivedCertificatesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessShortLivedCertificatesDataSourceModel struct {
	AccountID types.String                                                                             `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                                             `tfsdk:"zone_id" path:"zone_id,optional"`
	MaxItems  types.Int64                                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessShortLivedCertificatesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessShortLivedCertificatesDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessApplicationCAListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationCAListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessShortLivedCertificatesResultDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AUD       types.String `tfsdk:"aud" json:"aud,computed"`
	PublicKey types.String `tfsdk:"public_key" json:"public_key,computed"`
}
