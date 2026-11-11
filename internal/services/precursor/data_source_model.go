// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package precursor

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/precursor"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PrecursorResultDataSourceEnvelope struct {
	Result PrecursorDataSourceModel `json:"result,computed"`
}

type PrecursorDataSourceModel struct {
	ID               types.String                                                `tfsdk:"id" path:"zone_id,computed"`
	ZoneID           types.String                                                `tfsdk:"zone_id" path:"zone_id,required"`
	Mode             types.String                                                `tfsdk:"mode" json:"mode,computed"`
	ShowInterstitial types.Bool                                                  `tfsdk:"show_interstitial" json:"show_interstitial,computed"`
	Rules            customfield.NestedObjectList[PrecursorRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
}

func (m *PrecursorDataSourceModel) toReadParams(_ context.Context) (params precursor.PrecursorGetParams, diags diag.Diagnostics) {
	params = precursor.PrecursorGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type PrecursorRulesDataSourceModel struct {
	Expression       types.String `tfsdk:"expression" json:"expression,computed"`
	ShowInterstitial types.Bool   `tfsdk:"show_interstitial" json:"show_interstitial,computed"`
	ID               types.String `tfsdk:"id" json:"id,computed"`
	Description      types.String `tfsdk:"description" json:"description,computed"`
	Enabled          types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}
