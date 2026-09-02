// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/hostnames"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HostnameTLSSettingsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[HostnameTLSSettingsResultDataSourceModel] `json:"result,computed"`
}

type HostnameTLSSettingsDataSourceModel struct {
	SettingID types.String                                                           `tfsdk:"setting_id" path:"setting_id,required"`
	ZoneID    types.String                                                           `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems  types.Int64                                                            `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[HostnameTLSSettingsResultDataSourceModel] `tfsdk:"result"`
}

func (m *HostnameTLSSettingsDataSourceModel) toListParams(_ context.Context) (params hostnames.SettingTLSListParams, diags diag.Diagnostics) {
	params = hostnames.SettingTLSListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type HostnameTLSSettingsResultDataSourceModel struct {
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Hostname  types.String      `tfsdk:"hostname" json:"hostname,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Value     types.String      `tfsdk:"value" json:"value,computed"`
}
