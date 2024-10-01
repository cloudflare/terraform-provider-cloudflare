// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_normalization_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type URLNormalizationSettingsModel struct {
	ID     types.String `tfsdk:"id" json:"-,computed"`
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Scope  types.String `tfsdk:"scope" json:"scope,optional"`
	Type   types.String `tfsdk:"type" json:"type,optional"`
}

func (m URLNormalizationSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m URLNormalizationSettingsModel) MarshalJSONForUpdate(state URLNormalizationSettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
