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

type ZeroTrustDeviceManagedNetworksListResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[ZeroTrustDeviceManagedNetworksListResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDeviceManagedNetworksListDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[ZeroTrustDeviceManagedNetworksListResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDeviceManagedNetworksListDataSourceModel) toListParams(_ context.Context) (params zero_trust.DeviceNetworkListParams, diags diag.Diagnostics) {
  params = zero_trust.DeviceNetworkListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type ZeroTrustDeviceManagedNetworksListResultDataSourceModel struct {
Config customfield.NestedObject[ZeroTrustDeviceManagedNetworksListConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
NetworkID types.String `tfsdk:"network_id" json:"network_id,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type ZeroTrustDeviceManagedNetworksListConfigDataSourceModel struct {
TLSSockaddr types.String `tfsdk:"tls_sockaddr" json:"tls_sockaddr,computed"`
Sha256 types.String `tfsdk:"sha256" json:"sha256,computed"`
}
