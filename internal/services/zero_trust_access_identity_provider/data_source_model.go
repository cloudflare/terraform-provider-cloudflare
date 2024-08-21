// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessIdentityProviderResultDataSourceEnvelope struct {
	Result ZeroTrustAccessIdentityProviderDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessIdentityProviderResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessIdentityProviderDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessIdentityProviderDataSourceModel struct {
	AccountID          types.String                                             `tfsdk:"account_id" path:"account_id"`
	IdentityProviderID types.String                                             `tfsdk:"identity_provider_id" path:"identity_provider_id"`
	ZoneID             types.String                                             `tfsdk:"zone_id" path:"zone_id"`
	Filter             *ZeroTrustAccessIdentityProviderFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustAccessIdentityProviderDataSourceModel) toReadParams() (params zero_trust.IdentityProviderGetParams, diags diag.Diagnostics) {
	params = zero_trust.IdentityProviderGetParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

func (m *ZeroTrustAccessIdentityProviderDataSourceModel) toListParams() (params zero_trust.IdentityProviderListParams, diags diag.Diagnostics) {
	params = zero_trust.IdentityProviderListParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessIdentityProviderFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
