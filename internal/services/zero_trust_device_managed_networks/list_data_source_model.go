// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceManagedNetworksListResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustDeviceManagedNetworksListResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustDeviceManagedNetworksListDataSourceModel struct {
	AccountID types.String                                                `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                                 `tfsdk:"max_items"`
	Result    *[]*ZeroTrustDeviceManagedNetworksListResultDataSourceModel `tfsdk:"result"`
}

type ZeroTrustDeviceManagedNetworksListResultDataSourceModel struct {
	Config    *ZeroTrustDeviceManagedNetworksListConfigDataSourceModel `tfsdk:"config" json:"config"`
	Name      types.String                                             `tfsdk:"name" json:"name"`
	NetworkID types.String                                             `tfsdk:"network_id" json:"network_id"`
	Type      types.String                                             `tfsdk:"type" json:"type"`
}

type ZeroTrustDeviceManagedNetworksListConfigDataSourceModel struct {
	TLSSockaddr types.String `tfsdk:"tls_sockaddr" json:"tls_sockaddr,computed"`
	Sha256      types.String `tfsdk:"sha256" json:"sha256"`
}
