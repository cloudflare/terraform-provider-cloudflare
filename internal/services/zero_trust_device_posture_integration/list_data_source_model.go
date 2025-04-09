// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_integration

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDevicePostureIntegrationsResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[ZeroTrustDevicePostureIntegrationsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDevicePostureIntegrationsDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[ZeroTrustDevicePostureIntegrationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDevicePostureIntegrationsDataSourceModel) toListParams(_ context.Context) (params zero_trust.DevicePostureIntegrationListParams, diags diag.Diagnostics) {
  params = zero_trust.DevicePostureIntegrationListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type ZeroTrustDevicePostureIntegrationsResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Config customfield.NestedObject[ZeroTrustDevicePostureIntegrationsConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
Interval types.String `tfsdk:"interval" json:"interval,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type ZeroTrustDevicePostureIntegrationsConfigDataSourceModel struct {
APIURL types.String `tfsdk:"api_url" json:"api_url,computed"`
AuthURL types.String `tfsdk:"auth_url" json:"auth_url,computed"`
ClientID types.String `tfsdk:"client_id" json:"client_id,computed"`
}
