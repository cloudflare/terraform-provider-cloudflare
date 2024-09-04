// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceManagedNetworksResultEnvelope struct {
	Result ZeroTrustDeviceManagedNetworksModel `json:"result"`
}

type ZeroTrustDeviceManagedNetworksModel struct {
	ID        types.String                                                        `tfsdk:"id" json:"-,computed"`
	NetworkID types.String                                                        `tfsdk:"network_id" json:"network_id,computed"`
	AccountID types.String                                                        `tfsdk:"account_id" path:"account_id"`
	Name      types.String                                                        `tfsdk:"name" json:"name,computed_optional"`
	Type      types.String                                                        `tfsdk:"type" json:"type,computed_optional"`
	Config    customfield.NestedObject[ZeroTrustDeviceManagedNetworksConfigModel] `tfsdk:"config" json:"config,computed_optional"`
}

type ZeroTrustDeviceManagedNetworksConfigModel struct {
	TLSSockaddr types.String `tfsdk:"tls_sockaddr" json:"tls_sockaddr,computed_optional"`
	Sha256      types.String `tfsdk:"sha256" json:"sha256,computed_optional"`
}
