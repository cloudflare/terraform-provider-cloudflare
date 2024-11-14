// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessApplicationsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessApplicationsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessApplicationsDataSourceModel struct {
	AccountID types.String                                                                   `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                                   `tfsdk:"zone_id" path:"zone_id,optional"`
	AUD       types.String                                                                   `tfsdk:"aud" query:"aud,optional"`
	Domain    types.String                                                                   `tfsdk:"domain" query:"domain,optional"`
	Name      types.String                                                                   `tfsdk:"name" query:"name,optional"`
	Search    types.String                                                                   `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                                    `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessApplicationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessApplicationsDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessApplicationListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationListParams{}

	if !m.AUD.IsNull() {
		params.AUD = cloudflare.F(m.AUD.ValueString())
	}
	if !m.Domain.IsNull() {
		params.Domain = cloudflare.F(m.Domain.ValueString())
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessApplicationsResultDataSourceModel struct {
}
