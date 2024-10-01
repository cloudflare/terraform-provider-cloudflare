// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_certificates

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceCertificatesResultEnvelope struct {
	Result ZeroTrustDeviceCertificatesModel `json:"result"`
}

type ZeroTrustDeviceCertificatesModel struct {
	ZoneTag types.String `tfsdk:"zone_tag" path:"zone_tag,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,required"`
}

func (m ZeroTrustDeviceCertificatesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDeviceCertificatesModel) MarshalJSONForUpdate(state ZeroTrustDeviceCertificatesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
