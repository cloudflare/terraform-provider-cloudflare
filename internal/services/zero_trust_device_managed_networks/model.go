// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDeviceManagedNetworksResultEnvelope struct {
Result ZeroTrustDeviceManagedNetworksModel `json:"result"`
}

type ZeroTrustDeviceManagedNetworksModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
NetworkID types.String `tfsdk:"network_id" json:"network_id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Name types.String `tfsdk:"name" json:"name,required"`
Type types.String `tfsdk:"type" json:"type,required"`
Config *ZeroTrustDeviceManagedNetworksConfigModel `tfsdk:"config" json:"config,required"`
}

func (m ZeroTrustDeviceManagedNetworksModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m ZeroTrustDeviceManagedNetworksModel) MarshalJSONForUpdate(state ZeroTrustDeviceManagedNetworksModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDeviceManagedNetworksConfigModel struct {
TLSSockaddr types.String `tfsdk:"tls_sockaddr" json:"tls_sockaddr,required"`
Sha256 types.String `tfsdk:"sha256" json:"sha256,optional"`
}
