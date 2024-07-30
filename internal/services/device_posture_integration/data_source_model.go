// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_integration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DevicePostureIntegrationResultDataSourceEnvelope struct {
	Result DevicePostureIntegrationDataSourceModel `json:"result,computed"`
}

type DevicePostureIntegrationResultListDataSourceEnvelope struct {
	Result *[]*DevicePostureIntegrationDataSourceModel `json:"result,computed"`
}

type DevicePostureIntegrationDataSourceModel struct {
	AccountID     types.String                                      `tfsdk:"account_id" path:"account_id"`
	IntegrationID types.String                                      `tfsdk:"integration_id" path:"integration_id"`
	ID            types.String                                      `tfsdk:"id" json:"id"`
	Config        *DevicePostureIntegrationConfigDataSourceModel    `tfsdk:"config" json:"config"`
	Interval      types.String                                      `tfsdk:"interval" json:"interval"`
	Name          types.String                                      `tfsdk:"name" json:"name"`
	Type          types.String                                      `tfsdk:"type" json:"type"`
	Filter        *DevicePostureIntegrationFindOneByDataSourceModel `tfsdk:"filter"`
}

type DevicePostureIntegrationConfigDataSourceModel struct {
	APIURL   types.String `tfsdk:"api_url" json:"api_url,computed"`
	AuthURL  types.String `tfsdk:"auth_url" json:"auth_url,computed"`
	ClientID types.String `tfsdk:"client_id" json:"client_id,computed"`
}

type DevicePostureIntegrationFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
