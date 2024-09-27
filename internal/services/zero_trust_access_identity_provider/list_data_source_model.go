// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessIdentityProvidersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessIdentityProvidersResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessIdentityProvidersDataSourceModel struct {
	AccountID types.String                                                                        `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                                        `tfsdk:"zone_id" path:"zone_id,optional"`
	MaxItems  types.Int64                                                                         `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessIdentityProvidersResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessIdentityProvidersDataSourceModel) toListParams(_ context.Context) (params zero_trust.IdentityProviderListParams, diags diag.Diagnostics) {
	params = zero_trust.IdentityProviderListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessIdentityProvidersResultDataSourceModel struct {
}
