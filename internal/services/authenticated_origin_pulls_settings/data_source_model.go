// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/origin_tls_client_auth"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsSettingsResultDataSourceEnvelope struct {
	Result AuthenticatedOriginPullsSettingsDataSourceModel `json:"result,computed"`
}

type AuthenticatedOriginPullsSettingsDataSourceModel struct {
	ZoneID  types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}

func (m *AuthenticatedOriginPullsSettingsDataSourceModel) toReadParams(_ context.Context) (params origin_tls_client_auth.SettingGetParams, diags diag.Diagnostics) {
	params = origin_tls_client_auth.SettingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
