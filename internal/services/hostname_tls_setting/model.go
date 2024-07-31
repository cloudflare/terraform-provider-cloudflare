// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HostnameTLSSettingResultEnvelope struct {
	Result HostnameTLSSettingModel `json:"result,computed"`
}

type HostnameTLSSettingModel struct {
	ID        types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID    types.String      `tfsdk:"zone_id" path:"zone_id"`
	SettingID types.String      `tfsdk:"setting_id" path:"setting_id"`
	Hostname  types.String      `tfsdk:"hostname" path:"hostname"`
	Value     types.Float64     `tfsdk:"value" json:"value"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}
