// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile_certificates

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceDefaultProfileCertificatesResultEnvelope struct {
Result ZeroTrustDeviceDefaultProfileCertificatesModel `json:"result"`
}

type ZeroTrustDeviceDefaultProfileCertificatesModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,required"`
}

func (m ZeroTrustDeviceDefaultProfileCertificatesModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m ZeroTrustDeviceDefaultProfileCertificatesModel) MarshalJSONForUpdate(state ZeroTrustDeviceDefaultProfileCertificatesModel) (data []byte, err error) {
  return apijson.MarshalForPatch(m, state)
}
