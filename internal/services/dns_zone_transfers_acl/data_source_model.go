// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_acl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransfersACLResultDataSourceEnvelope struct {
	Result DNSZoneTransfersACLDataSourceModel `json:"result,computed"`
}

type DNSZoneTransfersACLDataSourceModel struct {
	ID        types.String `tfsdk:"id" path:"acl_id,computed"`
	ACLID     types.String `tfsdk:"acl_id" path:"acl_id,optional"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	IPRange   types.String `tfsdk:"ip_range" json:"ip_range,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
}

func (m *DNSZoneTransfersACLDataSourceModel) toReadParams(_ context.Context) (params dns.ZoneTransferACLGetParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferACLGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
