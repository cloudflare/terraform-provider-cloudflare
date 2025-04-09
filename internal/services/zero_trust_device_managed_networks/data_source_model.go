// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceManagedNetworksResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceManagedNetworksDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceManagedNetworksDataSourceModel struct {
	ID        types.String                                                                  `tfsdk:"id" path:"network_id,computed"`
	NetworkID types.String                                                                  `tfsdk:"network_id" path:"network_id,computed_optional"`
	AccountID types.String                                                                  `tfsdk:"account_id" path:"account_id,required"`
	Name      types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Type      types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Config    customfield.NestedObject[ZeroTrustDeviceManagedNetworksConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
}

func (m *ZeroTrustDeviceManagedNetworksDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DeviceNetworkGetParams, diags diag.Diagnostics) {
	params = zero_trust.DeviceNetworkGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDeviceManagedNetworksConfigDataSourceModel struct {
	TLSSockaddr types.String `tfsdk:"tls_sockaddr" json:"tls_sockaddr,computed"`
	Sha256      types.String `tfsdk:"sha256" json:"sha256,computed"`
}
