// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_managed_networks

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DeviceManagedNetworksResultEnvelope struct {
	Result DeviceManagedNetworksModel `json:"result,computed"`
}

type DeviceManagedNetworksModel struct {
	NetworkID types.String                      `tfsdk:"network_id" json:"network_id,computed"`
	AccountID types.String                      `tfsdk:"account_id" path:"account_id"`
	Config    *DeviceManagedNetworksConfigModel `tfsdk:"config" json:"config"`
	Name      types.String                      `tfsdk:"name" json:"name"`
	Type      types.String                      `tfsdk:"type" json:"type"`
}

type DeviceManagedNetworksConfigModel struct {
	TLSSockaddr types.String `tfsdk:"tls_sockaddr" json:"tls_sockaddr"`
	Sha256      types.String `tfsdk:"sha256" json:"sha256"`
}
