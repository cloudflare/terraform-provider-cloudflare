// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_settings

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsSettingsResultEnvelope struct {
Result AuthenticatedOriginPullsSettingsModel `json:"result"`
}

type AuthenticatedOriginPullsSettingsModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,required"`
}

func (m AuthenticatedOriginPullsSettingsModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m AuthenticatedOriginPullsSettingsModel) MarshalJSONForUpdate(state AuthenticatedOriginPullsSettingsModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
