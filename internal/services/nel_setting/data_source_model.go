// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nel_setting

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NELSettingResultDataSourceEnvelope struct {
	Result NELSettingDataSourceModel `json:"result,computed"`
}

type NELSettingDataSourceModel struct {
	ID         types.String                                             `tfsdk:"id" path:"zone_id,computed"`
	ZoneID     types.String                                             `tfsdk:"zone_id" path:"zone_id,required"`
	Editable   types.Bool                                               `tfsdk:"editable" json:"editable,computed"`
	ModifiedOn timetypes.RFC3339                                        `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Value      customfield.NestedObject[NELSettingValueDataSourceModel] `tfsdk:"value" json:"value,computed"`
}

func (m *NELSettingDataSourceModel) toReadParams(_ context.Context) (params zones.NELGetParams, diags diag.Diagnostics) {
	params = zones.NELGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type NELSettingValueDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}
