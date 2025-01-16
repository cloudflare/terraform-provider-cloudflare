// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_connector

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitConnectorResultDataSourceEnvelope struct {
	Result MagicTransitConnectorDataSourceModel `json:"result,computed"`
}

type MagicTransitConnectorResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicTransitConnectorDataSourceModel] `json:"result,computed"`
}

type MagicTransitConnectorDataSourceModel struct {
	AccountID                    types.String                                                         `tfsdk:"account_id" path:"account_id,optional"`
	ConnectorID                  types.String                                                         `tfsdk:"connector_id" path:"connector_id,optional"`
	Activated                    types.Bool                                                           `tfsdk:"activated" json:"activated,computed"`
	ID                           types.String                                                         `tfsdk:"id" json:"id,computed"`
	InterruptWindowDurationHours types.Float64                                                        `tfsdk:"interrupt_window_duration_hours" json:"interrupt_window_duration_hours,computed"`
	InterruptWindowHourOfDay     types.Float64                                                        `tfsdk:"interrupt_window_hour_of_day" json:"interrupt_window_hour_of_day,computed"`
	LastHeartbeat                types.String                                                         `tfsdk:"last_heartbeat" json:"last_heartbeat,computed"`
	LastSeenVersion              types.String                                                         `tfsdk:"last_seen_version" json:"last_seen_version,computed"`
	LastUpdated                  types.String                                                         `tfsdk:"last_updated" json:"last_updated,computed"`
	Notes                        types.String                                                         `tfsdk:"notes" json:"notes,computed"`
	Timezone                     types.String                                                         `tfsdk:"timezone" json:"timezone,computed"`
	Device                       customfield.NestedObject[MagicTransitConnectorDeviceDataSourceModel] `tfsdk:"device" json:"device,computed"`
	Filter                       *MagicTransitConnectorFindOneByDataSourceModel                       `tfsdk:"filter"`
}

func (m *MagicTransitConnectorDataSourceModel) toReadParams(_ context.Context) (params magic_transit.ConnectorGetParams, diags diag.Diagnostics) {
	params = magic_transit.ConnectorGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *MagicTransitConnectorDataSourceModel) toListParams(_ context.Context) (params magic_transit.ConnectorListParams, diags diag.Diagnostics) {
	params = magic_transit.ConnectorListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type MagicTransitConnectorDeviceDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	SerialNumber types.String `tfsdk:"serial_number" json:"serial_number,computed"`
}

type MagicTransitConnectorFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
