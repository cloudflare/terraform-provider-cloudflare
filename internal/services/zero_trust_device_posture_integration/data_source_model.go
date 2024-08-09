// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_integration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDevicePostureIntegrationResultDataSourceEnvelope struct {
	Result ZeroTrustDevicePostureIntegrationDataSourceModel `json:"result,computed"`
}

type ZeroTrustDevicePostureIntegrationResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustDevicePostureIntegrationDataSourceModel `json:"result,computed"`
}

type ZeroTrustDevicePostureIntegrationDataSourceModel struct {
	AccountID     types.String                                               `tfsdk:"account_id" path:"account_id"`
	IntegrationID types.String                                               `tfsdk:"integration_id" path:"integration_id"`
	ID            types.String                                               `tfsdk:"id" json:"id"`
	Interval      types.String                                               `tfsdk:"interval" json:"interval"`
	Name          types.String                                               `tfsdk:"name" json:"name"`
	Type          types.String                                               `tfsdk:"type" json:"type"`
	Config        *ZeroTrustDevicePostureIntegrationConfigDataSourceModel    `tfsdk:"config" json:"config"`
	Filter        *ZeroTrustDevicePostureIntegrationFindOneByDataSourceModel `tfsdk:"filter"`
}

type ZeroTrustDevicePostureIntegrationConfigDataSourceModel struct {
	APIURL   types.String `tfsdk:"api_url" json:"api_url,computed"`
	AuthURL  types.String `tfsdk:"auth_url" json:"auth_url,computed"`
	ClientID types.String `tfsdk:"client_id" json:"client_id,computed"`
}

type ZeroTrustDevicePostureIntegrationFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
