// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_acl

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSACLResultEnvelope struct {
	Result SecondaryDNSACLModel `json:"result,computed"`
}

type SecondaryDNSACLModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
