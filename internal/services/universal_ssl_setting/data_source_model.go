// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package universal_ssl_setting

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/ssl"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UniversalSSLSettingResultDataSourceEnvelope struct {
	Result UniversalSSLSettingDataSourceModel `json:"result,computed"`
}

type UniversalSSLSettingDataSourceModel struct {
	ZoneID  types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}

func (m *UniversalSSLSettingDataSourceModel) toReadParams(_ context.Context) (params ssl.UniversalSettingGetParams, diags diag.Diagnostics) {
	params = ssl.UniversalSettingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
