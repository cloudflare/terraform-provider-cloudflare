// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_tsig

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/secondary_dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSTSIGsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SecondaryDNSTSIGsResultDataSourceModel] `json:"result,computed"`
}

type SecondaryDNSTSIGsDataSourceModel struct {
	AccountID types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[SecondaryDNSTSIGsResultDataSourceModel] `tfsdk:"result"`
}

func (m *SecondaryDNSTSIGsDataSourceModel) toListParams(_ context.Context) (params secondary_dns.TSIGListParams, diags diag.Diagnostics) {
	params = secondary_dns.TSIGListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type SecondaryDNSTSIGsResultDataSourceModel struct {
	ID     types.String `tfsdk:"id" json:"id,computed"`
	Algo   types.String `tfsdk:"algo" json:"algo,computed"`
	Name   types.String `tfsdk:"name" json:"name,computed"`
	Secret types.String `tfsdk:"secret" json:"secret,computed"`
}
