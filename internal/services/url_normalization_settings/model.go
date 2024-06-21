// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_normalization_settings

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type URLNormalizationSettingsModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Scope  types.String `tfsdk:"scope" json:"scope"`
	Type   types.String `tfsdk:"type" json:"type"`
}
