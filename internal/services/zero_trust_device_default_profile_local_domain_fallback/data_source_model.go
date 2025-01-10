// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile_local_domain_fallback

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceDefaultProfileLocalDomainFallbackResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceDefaultProfileLocalDomainFallbackDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceDefaultProfileLocalDomainFallbackDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

func (m *ZeroTrustDeviceDefaultProfileLocalDomainFallbackDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DevicePolicyDefaultFallbackDomainGetParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyDefaultFallbackDomainGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
