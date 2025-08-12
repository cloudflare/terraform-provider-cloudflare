// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile_certificates

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceDefaultProfileCertificatesResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceDefaultProfileCertificatesDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceDefaultProfileCertificatesDataSourceModel struct {
	ZoneID  types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}

func (m *ZeroTrustDeviceDefaultProfileCertificatesDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DevicePolicyDefaultCertificateGetParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyDefaultCertificateGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
