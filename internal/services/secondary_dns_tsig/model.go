// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_tsig

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSTSIGResultEnvelope struct {
	Result SecondaryDNSTSIGModel `json:"result"`
}

type SecondaryDNSTSIGModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Algo      types.String `tfsdk:"algo" json:"algo,required"`
	Name      types.String `tfsdk:"name" json:"name,required"`
	Secret    types.String `tfsdk:"secret" json:"secret,required"`
}
