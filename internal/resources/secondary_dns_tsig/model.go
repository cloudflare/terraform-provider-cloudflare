// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_tsig

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSTSIGResultEnvelope struct {
	Result SecondaryDNSTSIGModel `json:"result,computed"`
}

type SecondaryDNSTSIGModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	Algo      types.String `tfsdk:"algo" json:"algo"`
	Name      types.String `tfsdk:"name" json:"name"`
	Secret    types.String `tfsdk:"secret" json:"secret"`
}
