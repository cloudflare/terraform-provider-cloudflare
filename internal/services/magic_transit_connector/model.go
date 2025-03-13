// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_connector

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitConnectorResultEnvelope struct {
Result MagicTransitConnectorModel `json:"result"`
}

type MagicTransitConnectorModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
ConnectorID types.String `tfsdk:"connector_id" path:"connector_id,required"`
Activated types.Bool `tfsdk:"activated" json:"activated,optional"`
InterruptWindowDurationHours types.Float64 `tfsdk:"interrupt_window_duration_hours" json:"interrupt_window_duration_hours,optional"`
InterruptWindowHourOfDay types.Float64 `tfsdk:"interrupt_window_hour_of_day" json:"interrupt_window_hour_of_day,optional"`
Notes types.String `tfsdk:"notes" json:"notes,optional"`
Timezone types.String `tfsdk:"timezone" json:"timezone,optional"`
LastHeartbeat types.String `tfsdk:"last_heartbeat" json:"last_heartbeat,computed"`
LastSeenVersion types.String `tfsdk:"last_seen_version" json:"last_seen_version,computed"`
LastUpdated types.String `tfsdk:"last_updated" json:"last_updated,computed"`
Device customfield.NestedObject[MagicTransitConnectorDeviceModel] `tfsdk:"device" json:"device,computed"`
}

func (m MagicTransitConnectorModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m MagicTransitConnectorModel) MarshalJSONForUpdate(state MagicTransitConnectorModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

type MagicTransitConnectorDeviceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
SerialNumber types.String `tfsdk:"serial_number" json:"serial_number,computed"`
}
