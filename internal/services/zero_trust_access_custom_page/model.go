// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_page

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessCustomPageResultEnvelope struct {
	Result ZeroTrustAccessCustomPageModel `json:"result"`
}

type ZeroTrustAccessCustomPageModel struct {
	ID         types.String `tfsdk:"id" json:"-,computed"`
	UID        types.String `tfsdk:"uid" json:"uid,computed"`
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	CustomHTML types.String `tfsdk:"custom_html" json:"custom_html,required"`
	Name       types.String `tfsdk:"name" json:"name,required"`
	Type       types.String `tfsdk:"type" json:"type,required"`
}

func (m ZeroTrustAccessCustomPageModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessCustomPageModel) MarshalJSONForUpdate(state ZeroTrustAccessCustomPageModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
