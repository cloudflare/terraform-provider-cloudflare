// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HostnameTLSSettingResultEnvelope struct {
	Result HostnameTLSSettingModel `json:"result"`
}

type HostnameTLSSettingModel struct {
	ID        types.String      `tfsdk:"id" json:"-,computed"`
	SettingID types.String      `tfsdk:"setting_id" path:"setting_id,required"`
	ZoneID    types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Hostname  types.String      `tfsdk:"hostname" path:"hostname,optional"`
	Value     types.Float64     `tfsdk:"value" json:"value,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
