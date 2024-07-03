// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_managed_networks

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DeviceManagedNetworksResultDataSourceEnvelope struct {
	Result DeviceManagedNetworksDataSourceModel `json:"result,computed"`
}

type DeviceManagedNetworksResultListDataSourceEnvelope struct {
	Result *[]*DeviceManagedNetworksDataSourceModel `json:"result,computed"`
}

type DeviceManagedNetworksDataSourceModel struct {
	AccountID types.String                                   `tfsdk:"account_id" path:"account_id"`
	NetworkID types.String                                   `tfsdk:"network_id" path:"network_id"`
	Config    *DeviceManagedNetworksConfigDataSourceModel    `tfsdk:"config" json:"config"`
	Name      types.String                                   `tfsdk:"name" json:"name"`
	Type      types.String                                   `tfsdk:"type" json:"type"`
	FindOneBy *DeviceManagedNetworksFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type DeviceManagedNetworksConfigDataSourceModel struct {
	TLSSockaddr types.String `tfsdk:"tls_sockaddr" json:"tls_sockaddr"`
	Sha256      types.String `tfsdk:"sha256" json:"sha256"`
}

type DeviceManagedNetworksFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
