// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_certificates

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceCertificatesResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceCertificatesDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceCertificatesDataSourceModel struct {
	ZoneTag types.String `tfsdk:"zone_tag" path:"zone_tag,required"`
}
