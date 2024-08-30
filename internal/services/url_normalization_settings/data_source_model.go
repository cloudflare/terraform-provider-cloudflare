// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_normalization_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/url_normalization"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type URLNormalizationSettingsDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Scope  types.String `tfsdk:"scope" json:"scope"`
	Type   types.String `tfsdk:"type" json:"type"`
}

func (m *URLNormalizationSettingsDataSourceModel) toReadParams(_ context.Context) (params url_normalization.URLNormalizationGetParams, diags diag.Diagnostics) {
	params = url_normalization.URLNormalizationGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
