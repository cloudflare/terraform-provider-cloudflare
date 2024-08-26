// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessApplicationsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessApplicationsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessApplicationsDataSourceModel struct {
	AccountID types.String                                                                   `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                                                   `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                                                    `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessApplicationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessApplicationsDataSourceModel) toListParams() (params zero_trust.AccessApplicationListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessApplicationsResultDataSourceModel struct {
}
