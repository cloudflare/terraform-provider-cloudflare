// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_connector

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitConnectorsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicTransitConnectorsResultDataSourceModel] `json:"result,computed"`
}

type MagicTransitConnectorsDataSourceModel struct {
	AccountID types.String                                                              `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                               `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[MagicTransitConnectorsResultDataSourceModel] `tfsdk:"result"`
}

func (m *MagicTransitConnectorsDataSourceModel) toListParams(_ context.Context) (params magic_transit.ConnectorListParams, diags diag.Diagnostics) {
	params = magic_transit.ConnectorListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicTransitConnectorsResultDataSourceModel struct {
	ID                           types.String                                                          `tfsdk:"id" json:"id,computed"`
	Activated                    types.Bool                                                            `tfsdk:"activated" json:"activated,computed"`
	InterruptWindowDurationHours types.Float64                                                         `tfsdk:"interrupt_window_duration_hours" json:"interrupt_window_duration_hours,computed"`
	InterruptWindowHourOfDay     types.Float64                                                         `tfsdk:"interrupt_window_hour_of_day" json:"interrupt_window_hour_of_day,computed"`
	LastUpdated                  types.String                                                          `tfsdk:"last_updated" json:"last_updated,computed"`
	Notes                        types.String                                                          `tfsdk:"notes" json:"notes,computed"`
	Timezone                     types.String                                                          `tfsdk:"timezone" json:"timezone,computed"`
	Device                       customfield.NestedObject[MagicTransitConnectorsDeviceDataSourceModel] `tfsdk:"device" json:"device,computed"`
	LastHeartbeat                types.String                                                          `tfsdk:"last_heartbeat" json:"last_heartbeat,computed"`
	LastSeenVersion              types.String                                                          `tfsdk:"last_seen_version" json:"last_seen_version,computed"`
}

type MagicTransitConnectorsDeviceDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	SerialNumber types.String `tfsdk:"serial_number" json:"serial_number,computed"`
}
