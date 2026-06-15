// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_check

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LeakedCredentialCheckResultEnvelope struct {
	Result LeakedCredentialCheckModel `json:"result"`
}

type LeakedCredentialCheckModel struct {
	ZoneID  types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
}

func (m LeakedCredentialCheckModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m LeakedCredentialCheckModel) MarshalJSONForUpdate(state LeakedCredentialCheckModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
