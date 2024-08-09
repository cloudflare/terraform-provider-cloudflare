// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceManagedNetworksResultDataSourceEnvelope struct {
	Result ZeroTrustDeviceManagedNetworksDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceManagedNetworksResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustDeviceManagedNetworksDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceManagedNetworksDataSourceModel struct {
	AccountID types.String                                            `tfsdk:"account_id" path:"account_id"`
	NetworkID types.String                                            `tfsdk:"network_id" path:"network_id"`
	Name      types.String                                            `tfsdk:"name" json:"name"`
	Type      types.String                                            `tfsdk:"type" json:"type"`
	Config    *ZeroTrustDeviceManagedNetworksConfigDataSourceModel    `tfsdk:"config" json:"config"`
	Filter    *ZeroTrustDeviceManagedNetworksFindOneByDataSourceModel `tfsdk:"filter"`
}

type ZeroTrustDeviceManagedNetworksConfigDataSourceModel struct {
	TLSSockaddr types.String `tfsdk:"tls_sockaddr" json:"tls_sockaddr,computed"`
	Sha256      types.String `tfsdk:"sha256" json:"sha256"`
}

type ZeroTrustDeviceManagedNetworksFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
