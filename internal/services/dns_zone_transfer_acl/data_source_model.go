// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfer_acl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneTransferACLResultDataSourceEnvelope struct {
	Result DNSZoneTransferACLDataSourceModel `json:"result,computed"`
}

type DNSZoneTransferACLResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSZoneTransferACLDataSourceModel] `json:"result,computed"`
}

type DNSZoneTransferACLDataSourceModel struct {
	AccountID types.String                                `tfsdk:"account_id" path:"account_id,optional"`
	ACLID     types.String                                `tfsdk:"acl_id" path:"acl_id,optional"`
	ID        types.String                                `tfsdk:"id" json:"id,computed"`
	IPRange   types.String                                `tfsdk:"ip_range" json:"ip_range,computed"`
	Name      types.String                                `tfsdk:"name" json:"name,computed"`
	Filter    *DNSZoneTransferACLFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *DNSZoneTransferACLDataSourceModel) toReadParams(_ context.Context) (params dns.ZoneTransferACLGetParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferACLGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *DNSZoneTransferACLDataSourceModel) toListParams(_ context.Context) (params dns.ZoneTransferACLListParams, diags diag.Diagnostics) {
	params = dns.ZoneTransferACLListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type DNSZoneTransferACLFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
