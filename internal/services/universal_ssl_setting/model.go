// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package universal_ssl_setting

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UniversalSSLSettingResultEnvelope struct {
	Result UniversalSSLSettingModel `json:"result"`
}

type UniversalSSLSettingModel struct {
	ZoneID  types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
}

func (m UniversalSSLSettingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UniversalSSLSettingModel) MarshalJSONForUpdate(state UniversalSSLSettingModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
