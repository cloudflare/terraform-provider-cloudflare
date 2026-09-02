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
	ID               types.String                                                           `tfsdk:"id" path:"zone_id,computed"`
	ZoneID           types.String                                                           `tfsdk:"zone_id" path:"zone_id,required"`
	DefaultMode      types.String                                                           `tfsdk:"default_mode" json:"default_mode,computed"`
	EnforcementRules customfield.NestedObjectList[PrecursorEnforcementRulesDataSourceModel] `tfsdk:"enforcement_rules" json:"enforcement_rules,computed"`
}

func (m *PrecursorDataSourceModel) toReadParams(_ context.Context) (params precursor.PrecursorGetParams, diags diag.Diagnostics) {
	params = precursor.PrecursorGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type PrecursorEnforcementRulesDataSourceModel struct {
	Expression  types.String `tfsdk:"expression" json:"expression,computed"`
	Mode        types.String `tfsdk:"mode" json:"mode,computed"`
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}
