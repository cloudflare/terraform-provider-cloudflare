// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/page_shield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldSettingsResultDataSourceEnvelope struct {
	Result PageShieldSettingsDataSourceModel `json:"result,computed"`
}

type PageShieldSettingsDataSourceModel struct {
	ID                             types.String      `tfsdk:"id" path:"zone_id,computed"`
	ZoneID                         types.String      `tfsdk:"zone_id" path:"zone_id,optional"`
	Enabled                        types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	UpdatedAt                      timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UseCloudflareReportingEndpoint types.Bool        `tfsdk:"use_cloudflare_reporting_endpoint" json:"use_cloudflare_reporting_endpoint,computed"`
	UseConnectionURLPath           types.Bool        `tfsdk:"use_connection_url_path" json:"use_connection_url_path,computed"`
}

func (m *PageShieldSettingsDataSourceModel) toReadParams(_ context.Context) (params page_shield.PageShieldGetParams, diags diag.Diagnostics) {
	params = page_shield.PageShieldGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
