// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_acl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersACLsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSZoneTransfersACLsResultDataSourceModel] `json:"result,computed"`
}

type DNSZoneTransfersACLsDataSourceModel struct {
	AccountID types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[DNSZoneTransfersACLsResultDataSourceModel] `tfsdk:"result"`
}

func (m *DNSZoneTransfersACLsDataSourceModel) toListParams(_ context.Context) (params dns.ZoneTransferACLListParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferACLListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type DNSZoneTransfersACLsResultDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	IPRange types.String `tfsdk:"ip_range" json:"ip_range,computed"`
	Name    types.String `tfsdk:"name" json:"name,computed"`
}
