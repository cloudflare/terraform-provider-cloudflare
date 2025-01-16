// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_tsig

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersTSIGResultDataSourceEnvelope struct {
	Result DNSZoneTransfersTSIGDataSourceModel `json:"result,computed"`
}

type DNSZoneTransfersTSIGResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSZoneTransfersTSIGDataSourceModel] `json:"result,computed"`
}

type DNSZoneTransfersTSIGDataSourceModel struct {
	AccountID types.String                                  `tfsdk:"account_id" path:"account_id,optional"`
	TSIGID    types.String                                  `tfsdk:"tsig_id" path:"tsig_id,optional"`
	Algo      types.String                                  `tfsdk:"algo" json:"algo,computed"`
	ID        types.String                                  `tfsdk:"id" json:"id,computed"`
	Name      types.String                                  `tfsdk:"name" json:"name,computed"`
	Secret    types.String                                  `tfsdk:"secret" json:"secret,computed"`
	Filter    *DNSZoneTransfersTSIGFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *DNSZoneTransfersTSIGDataSourceModel) toReadParams(_ context.Context) (params dns.ZoneTransferTSIGGetParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferTSIGGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *DNSZoneTransfersTSIGDataSourceModel) toListParams(_ context.Context) (params dns.ZoneTransferTSIGListParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferTSIGListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type DNSZoneTransfersTSIGFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
