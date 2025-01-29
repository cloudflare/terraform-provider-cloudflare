// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_tsig

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersTSIGsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSZoneTransfersTSIGsResultDataSourceModel] `json:"result,computed"`
}

type DNSZoneTransfersTSIGsDataSourceModel struct {
	AccountID types.String                                                             `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[DNSZoneTransfersTSIGsResultDataSourceModel] `tfsdk:"result"`
}

func (m *DNSZoneTransfersTSIGsDataSourceModel) toListParams(_ context.Context) (params dns.ZoneTransferTSIGListParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferTSIGListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type DNSZoneTransfersTSIGsResultDataSourceModel struct {
	ID     types.String `tfsdk:"id" json:"id,computed"`
	Algo   types.String `tfsdk:"algo" json:"algo,computed"`
	Name   types.String `tfsdk:"name" json:"name,computed"`
	Secret types.String `tfsdk:"secret" json:"secret,computed"`
}
