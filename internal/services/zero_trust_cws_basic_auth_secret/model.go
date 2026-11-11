// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_basic_auth_secret

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSBasicAuthSecretResultEnvelope struct {
	Result ZeroTrustCWSBasicAuthSecretModel `json:"result"`
}

type ZeroTrustCWSBasicAuthSecretModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	AccountID    types.String      `tfsdk:"account_id" path:"account_id,required"`
	CredentialID types.String      `tfsdk:"credential_id" path:"credential_id,required"`
	ExpiresAt    timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,optional" format:"date-time"`
	Password     types.String      `tfsdk:"password" json:"password,optional,no_refresh"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
}

func (m ZeroTrustCWSBasicAuthSecretModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCWSBasicAuthSecretModel) MarshalJSONForUpdate(state ZeroTrustCWSBasicAuthSecretModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
