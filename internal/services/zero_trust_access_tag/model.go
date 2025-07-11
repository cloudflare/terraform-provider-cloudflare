// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tag

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessTagResultEnvelope struct {
	Result ZeroTrustAccessTagModel `json:"result"`
}

type ZeroTrustAccessTagModel struct {
	ID        types.String `tfsdk:"id" json:"-,computed"`
	Name      types.String `tfsdk:"name" json:"name,required"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

func (m ZeroTrustAccessTagModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessTagModel) MarshalJSONForUpdate(state ZeroTrustAccessTagModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
