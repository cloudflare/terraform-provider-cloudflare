// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_integration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DevicePostureIntegrationsResultListDataSourceEnvelope struct {
	Result *[]*DevicePostureIntegrationsItemsDataSourceModel `json:"result,computed"`
}

type DevicePostureIntegrationsDataSourceModel struct {
	AccountID types.String                                      `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                       `tfsdk:"max_items"`
	Items     *[]*DevicePostureIntegrationsItemsDataSourceModel `tfsdk:"items"`
}

type DevicePostureIntegrationsItemsDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	Interval types.String `tfsdk:"interval" json:"interval,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Type     types.String `tfsdk:"type" json:"type,computed"`
}

type DevicePostureIntegrationsItemsConfigDataSourceModel struct {
	APIURL   types.String `tfsdk:"api_url" json:"api_url,computed"`
	AuthURL  types.String `tfsdk:"auth_url" json:"auth_url,computed"`
	ClientID types.String `tfsdk:"client_id" json:"client_id,computed"`
}
