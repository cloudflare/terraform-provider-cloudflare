// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_devices_posture_integrations

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDevicesPostureIntegrationsResultEnvelope struct {
	Result ZeroTrustDevicesPostureIntegrationsModel `json:"result,computed"`
}

type ZeroTrustDevicesPostureIntegrationsModel struct {
	ID        types.String                                    `tfsdk:"id" json:"id,computed"`
	AccountID types.String                                    `tfsdk:"account_id" path:"account_id"`
	Config    *ZeroTrustDevicesPostureIntegrationsConfigModel `tfsdk:"config" json:"config"`
	Interval  types.String                                    `tfsdk:"interval" json:"interval"`
	Name      types.String                                    `tfsdk:"name" json:"name"`
	Type      types.String                                    `tfsdk:"type" json:"type"`
}

type ZeroTrustDevicesPostureIntegrationsConfigModel struct {
	APIURL             types.String `tfsdk:"api_url" json:"api_url"`
	AuthURL            types.String `tfsdk:"auth_url" json:"auth_url"`
	ClientID           types.String `tfsdk:"client_id" json:"client_id"`
	ClientSecret       types.String `tfsdk:"client_secret" json:"client_secret"`
	CustomerID         types.String `tfsdk:"customer_id" json:"customer_id"`
	ClientKey          types.String `tfsdk:"client_key" json:"client_key"`
	AccessClientID     types.String `tfsdk:"access_client_id" json:"access_client_id"`
	AccessClientSecret types.String `tfsdk:"access_client_secret" json:"access_client_secret"`
}
