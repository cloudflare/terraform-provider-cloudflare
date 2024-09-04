// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_integration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDevicePostureIntegrationResultEnvelope struct {
	Result ZeroTrustDevicePostureIntegrationModel `json:"result"`
}

type ZeroTrustDevicePostureIntegrationModel struct {
	ID        types.String                                  `tfsdk:"id" json:"id,computed"`
	AccountID types.String                                  `tfsdk:"account_id" path:"account_id"`
	Config    *ZeroTrustDevicePostureIntegrationConfigModel `tfsdk:"config" json:"config"`
	Interval  types.String                                  `tfsdk:"interval" json:"interval,computed_optional"`
	Name      types.String                                  `tfsdk:"name" json:"name,computed_optional"`
	Type      types.String                                  `tfsdk:"type" json:"type,computed_optional"`
}

type ZeroTrustDevicePostureIntegrationConfigModel struct {
	APIURL             types.String `tfsdk:"api_url" json:"api_url"`
	AuthURL            types.String `tfsdk:"auth_url" json:"auth_url"`
	ClientID           types.String `tfsdk:"client_id" json:"client_id"`
	ClientSecret       types.String `tfsdk:"client_secret" json:"client_secret"`
	CustomerID         types.String `tfsdk:"customer_id" json:"customer_id"`
	ClientKey          types.String `tfsdk:"client_key" json:"client_key"`
	AccessClientID     types.String `tfsdk:"access_client_id" json:"access_client_id"`
	AccessClientSecret types.String `tfsdk:"access_client_secret" json:"access_client_secret"`
}
