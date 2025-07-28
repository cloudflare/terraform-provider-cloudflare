package fallback_domain

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FallbackDomainResultEnvelope struct {
	Result *[]*FallbackDomainDomainsModel `json:"result"`
}

type FallbackDomainModel struct {
	ID        types.String                   `tfsdk:"id" json:"-,computed"`
	AccountID types.String                   `tfsdk:"account_id" path:"account_id,required"`
	Domains   *[]*FallbackDomainDomainsModel `tfsdk:"domains" json:"domains,required"`
}

func (m FallbackDomainModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Domains)
}

func (m FallbackDomainModel) MarshalJSONForUpdate(state FallbackDomainModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Domains, state.Domains)
}

type FallbackDomainDomainsModel struct {
	Suffix      types.String    `tfsdk:"suffix" json:"suffix,required"`
	Description types.String    `tfsdk:"description" json:"description,optional"`
	DNSServer   *[]types.String `tfsdk:"dns_server" json:"dns_server,optional"`
}
