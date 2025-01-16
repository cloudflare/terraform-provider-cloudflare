// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessShortLivedCertificateResultDataSourceEnvelope struct {
	Result ZeroTrustAccessShortLivedCertificateDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessShortLivedCertificateResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessShortLivedCertificateDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessShortLivedCertificateDataSourceModel struct {
	AccountID types.String                                                  `tfsdk:"account_id" path:"account_id,optional"`
	AppID     types.String                                                  `tfsdk:"app_id" path:"app_id,optional"`
	ZoneID    types.String                                                  `tfsdk:"zone_id" path:"zone_id,optional"`
	AUD       types.String                                                  `tfsdk:"aud" json:"aud,computed"`
	ID        types.String                                                  `tfsdk:"id" json:"id,computed"`
	PublicKey types.String                                                  `tfsdk:"public_key" json:"public_key,computed"`
	Filter    *ZeroTrustAccessShortLivedCertificateFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustAccessShortLivedCertificateDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessApplicationCAGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationCAGetParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

func (m *ZeroTrustAccessShortLivedCertificateDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessApplicationCAListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationCAListParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessShortLivedCertificateFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,optional"`
}
