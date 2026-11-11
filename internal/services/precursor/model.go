// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package precursor

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PrecursorResultEnvelope struct {
	Result PrecursorModel `json:"result"`
}

type PrecursorModel struct {
	ID               types.String                                      `tfsdk:"id" json:"-,computed"`
	ZoneID           types.String                                      `tfsdk:"zone_id" path:"zone_id,required"`
	ShowInterstitial types.Bool                                        `tfsdk:"show_interstitial" json:"show_interstitial,optional"`
	Mode             types.String                                      `tfsdk:"mode" json:"mode,computed_optional"`
	Rules            customfield.NestedObjectList[PrecursorRulesModel] `tfsdk:"rules" json:"rules,computed_optional"`
}

func (m PrecursorModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PrecursorModel) MarshalJSONForUpdate(state PrecursorModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type PrecursorRulesModel struct {
	Expression       types.String `tfsdk:"expression" json:"expression,required"`
	ShowInterstitial types.Bool   `tfsdk:"show_interstitial" json:"show_interstitial,required"`
	ID               types.String `tfsdk:"id" json:"id,computed"`
	Description      types.String `tfsdk:"description" json:"description,computed_optional"`
	Enabled          types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
}
