// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_peer

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSPeerResultEnvelope struct {
	Result SecondaryDNSPeerModel `json:"result,computed"`
}

type SecondaryDNSPeerModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
