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

type ZeroTrustAccessApplicationResultDataSourceEnvelope struct {
	Result ZeroTrustAccessApplicationDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessApplicationResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessApplicationDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessApplicationDataSourceModel struct {
	AccountID types.String                                        `tfsdk:"account_id" path:"account_id,optional"`
	AppID     types.String                                        `tfsdk:"app_id" path:"app_id,optional"`
	ZoneID    types.String                                        `tfsdk:"zone_id" path:"zone_id,optional"`
	Filter    *ZeroTrustAccessApplicationFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustAccessApplicationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessApplicationGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationGetParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

func (m *ZeroTrustAccessApplicationDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessApplicationListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationListParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessApplicationFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,optional"`
}
