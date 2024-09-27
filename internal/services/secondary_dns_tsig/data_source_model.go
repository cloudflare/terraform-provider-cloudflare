// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_tsig

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/secondary_dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSTSIGResultDataSourceEnvelope struct {
	Result SecondaryDNSTSIGDataSourceModel `json:"result,computed"`
}

type SecondaryDNSTSIGResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SecondaryDNSTSIGDataSourceModel] `json:"result,computed"`
}

type SecondaryDNSTSIGDataSourceModel struct {
	AccountID types.String                              `tfsdk:"account_id" path:"account_id,optional"`
	TSIGID    types.String                              `tfsdk:"tsig_id" path:"tsig_id,optional"`
	Algo      types.String                              `tfsdk:"algo" json:"algo,computed"`
	ID        types.String                              `tfsdk:"id" json:"id,computed"`
	Name      types.String                              `tfsdk:"name" json:"name,computed"`
	Secret    types.String                              `tfsdk:"secret" json:"secret,computed"`
	Filter    *SecondaryDNSTSIGFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *SecondaryDNSTSIGDataSourceModel) toReadParams(_ context.Context) (params secondary_dns.TSIGGetParams, diags diag.Diagnostics) {
	params = secondary_dns.TSIGGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *SecondaryDNSTSIGDataSourceModel) toListParams(_ context.Context) (params secondary_dns.TSIGListParams, diags diag.Diagnostics) {
	params = secondary_dns.TSIGListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type SecondaryDNSTSIGFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
