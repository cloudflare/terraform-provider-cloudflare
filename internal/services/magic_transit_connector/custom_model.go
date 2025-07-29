package magic_transit_connector

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomMagicTransitConnectorResultEnvelope struct {
	Result CustomMagicTransitConnectorModel `json:"result"`
}

type CustomMagicTransitConnectorModel struct {
	ID                           types.String                            `tfsdk:"id" json:"id,computed"`
	AccountID                    types.String                            `tfsdk:"account_id" path:"account_id,required"`
	Device                       *CustomMagicTransitConnectorDeviceModel `tfsdk:"device" json:"device,required"`
	Activated                    types.Bool                              `tfsdk:"activated" json:"activated,computed_optional"`
	InterruptWindowDurationHours types.Float64                           `tfsdk:"interrupt_window_duration_hours" json:"interrupt_window_duration_hours,computed_optional"`
	InterruptWindowHourOfDay     types.Float64                           `tfsdk:"interrupt_window_hour_of_day" json:"interrupt_window_hour_of_day,computed_optional"`
	Notes                        types.String                            `tfsdk:"notes" json:"notes,computed_optional"`
	Timezone                     types.String                            `tfsdk:"timezone" json:"timezone,computed_optional"`
}

func (m CustomMagicTransitConnectorModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomMagicTransitConnectorModel) MarshalJSONForUpdate(state CustomMagicTransitConnectorModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type CustomMagicTransitConnectorDeviceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed_optional"`
	SerialNumber types.String `tfsdk:"serial_number" json:"serial_number,computed_optional"`
}
