// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_security

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISecurityResultEnvelope struct {
	Result AISecurityModel `json:"result"`
}

type AISecurityModel struct {
	ID      types.String `tfsdk:"id" json:"-,computed"`
	ZoneID  types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
}

func (m AISecurityModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AISecurityModel) MarshalJSONForUpdate(state AISecurityModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
