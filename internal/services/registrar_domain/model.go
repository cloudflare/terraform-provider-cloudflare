// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegistrarDomainResultEnvelope struct {
	Result RegistrarDomainModel `json:"result"`
}

type RegistrarDomainModel struct {
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	DomainName types.String `tfsdk:"domain_name" path:"domain_name,required"`
	AutoRenew  types.Bool   `tfsdk:"auto_renew" json:"auto_renew,optional"`
	Locked     types.Bool   `tfsdk:"locked" json:"locked,optional"`
	Privacy    types.Bool   `tfsdk:"privacy" json:"privacy,optional"`
}
