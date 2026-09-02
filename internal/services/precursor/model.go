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
	ID               types.String                                                 `tfsdk:"id" json:"-,computed"`
	ZoneID           types.String                                                 `tfsdk:"zone_id" path:"zone_id,required"`
	DefaultMode      types.String                                                 `tfsdk:"default_mode" json:"default_mode,computed_optional"`
	EnforcementRules customfield.NestedObjectList[PrecursorEnforcementRulesModel] `tfsdk:"enforcement_rules" json:"enforcement_rules,computed_optional"`
}

func (m PrecursorModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PrecursorModel) MarshalJSONForUpdate(state PrecursorModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type PrecursorEnforcementRulesModel struct {
	Expression  types.String `tfsdk:"expression" json:"expression,required"`
	Mode        types.String `tfsdk:"mode" json:"mode,required"`
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed_optional"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
}
