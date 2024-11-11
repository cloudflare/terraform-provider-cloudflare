// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile_certificates

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceDefaultProfileCertificatesResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceDefaultProfileCertificatesDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceDefaultProfileCertificatesDataSourceModel struct {
	ZoneTag types.String `tfsdk:"zone_tag" path:"zone_tag,required"`
}
