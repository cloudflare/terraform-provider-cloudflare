// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegistrarDomainResultEnvelope struct {
	Result RegistrarDomainModel `json:"result"`
}

type RegistrarDomainModel struct {
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	DomainName types.String `tfsdk:"domain_name" path:"domain_name,required"`
	AutoRenew  types.Bool   `tfsdk:"auto_renew" json:"auto_renew,optional,no_refresh"`
	Locked     types.Bool   `tfsdk:"locked" json:"locked,optional,no_refresh"`
	Privacy    types.Bool   `tfsdk:"privacy" json:"privacy,optional,no_refresh"`
}

func (m RegistrarDomainModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m RegistrarDomainModel) MarshalJSONForUpdate(state RegistrarDomainModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
