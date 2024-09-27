// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_acls

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSACLsResultEnvelope struct {
	Result SecondaryDNSACLsModel `json:"result"`
}

type SecondaryDNSACLsModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	IPRange   types.String `tfsdk:"ip_range" json:"ip_range,required"`
	Name      types.String `tfsdk:"name" json:"name,required"`
}
