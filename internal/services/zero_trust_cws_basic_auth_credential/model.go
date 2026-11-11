// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_basic_auth_credential

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSBasicAuthCredentialResultEnvelope struct {
	Result ZeroTrustCWSBasicAuthCredentialModel `json:"result"`
}

type ZeroTrustCWSBasicAuthCredentialModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	Username  types.String      `tfsdk:"username" json:"username,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustCWSBasicAuthCredentialModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCWSBasicAuthCredentialModel) MarshalJSONForUpdate(state ZeroTrustCWSBasicAuthCredentialModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
