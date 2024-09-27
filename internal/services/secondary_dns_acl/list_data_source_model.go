// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_acl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/secondary_dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSACLsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SecondaryDNSACLsResultDataSourceModel] `json:"result,computed"`
}

type SecondaryDNSACLsDataSourceModel struct {
	AccountID types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                         `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[SecondaryDNSACLsResultDataSourceModel] `tfsdk:"result"`
}

func (m *SecondaryDNSACLsDataSourceModel) toListParams(_ context.Context) (params secondary_dns.ACLListParams, diags diag.Diagnostics) {
	params = secondary_dns.ACLListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type SecondaryDNSACLsResultDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	IPRange types.String `tfsdk:"ip_range" json:"ip_range,computed"`
	Name    types.String `tfsdk:"name" json:"name,computed"`
}
