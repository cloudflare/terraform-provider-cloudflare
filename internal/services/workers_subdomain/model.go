// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_subdomain

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersSubdomainResultEnvelope struct {
	Result *WorkersSubdomainModel `json:"result"`
}

type WorkersSubdomainModel struct {
	ID        types.String `tfsdk:"id" json:"-,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Subdomain types.String `tfsdk:"subdomain" json:"subdomain,required"`
}

func (m WorkersSubdomainModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkersSubdomainModel) MarshalJSONForUpdate(state WorkersSubdomainModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
