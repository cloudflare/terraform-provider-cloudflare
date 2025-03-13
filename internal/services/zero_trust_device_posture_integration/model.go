// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_integration

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDevicePostureIntegrationResultEnvelope struct {
Result ZeroTrustDevicePostureIntegrationModel `json:"result"`
}

type ZeroTrustDevicePostureIntegrationModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Interval types.String `tfsdk:"interval" json:"interval,required"`
Name types.String `tfsdk:"name" json:"name,required"`
Type types.String `tfsdk:"type" json:"type,required"`
Config *ZeroTrustDevicePostureIntegrationConfigModel `tfsdk:"config" json:"config,required"`
}

func (m ZeroTrustDevicePostureIntegrationModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m ZeroTrustDevicePostureIntegrationModel) MarshalJSONForUpdate(state ZeroTrustDevicePostureIntegrationModel) (data []byte, err error) {
  return apijson.MarshalForPatch(m, state)
}

type ZeroTrustDevicePostureIntegrationConfigModel struct {
APIURL types.String `tfsdk:"api_url" json:"api_url,optional"`
AuthURL types.String `tfsdk:"auth_url" json:"auth_url,optional"`
ClientID types.String `tfsdk:"client_id" json:"client_id,optional"`
ClientSecret types.String `tfsdk:"client_secret" json:"client_secret,optional"`
CustomerID types.String `tfsdk:"customer_id" json:"customer_id,optional"`
ClientKey types.String `tfsdk:"client_key" json:"client_key,optional"`
AccessClientID types.String `tfsdk:"access_client_id" json:"access_client_id,optional"`
AccessClientSecret types.String `tfsdk:"access_client_secret" json:"access_client_secret,optional"`
}
