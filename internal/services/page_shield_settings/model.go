// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldSettingsResultEnvelope struct {
	Result PageShieldSettingsModel `json:"result"`
}

type PageShieldSettingsModel struct {
	ID                             types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID                         types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled                        types.Bool        `tfsdk:"enabled" json:"enabled,computed_optional"`
	UpdatedAt                      timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UseCloudflareReportingEndpoint types.Bool        `tfsdk:"use_cloudflare_reporting_endpoint" json:"use_cloudflare_reporting_endpoint,computed_optional"`
	UseConnectionURLPath           types.Bool        `tfsdk:"use_connection_url_path" json:"use_connection_url_path,computed_optional"`
}

func (m PageShieldSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PageShieldSettingsModel) MarshalJSONForUpdate(state PageShieldSettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
