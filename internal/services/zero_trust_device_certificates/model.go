// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_certificates

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceCertificatesResultEnvelope struct {
	Result ZeroTrustDeviceCertificatesModel `json:"result"`
}

type ZeroTrustDeviceCertificatesModel struct {
	ZoneTag types.String `tfsdk:"zone_tag" path:"zone_tag"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
}
