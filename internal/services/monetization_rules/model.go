// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monetization_rules

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MonetizationRulesResultEnvelope struct {
	Result MonetizationRulesModel `json:"result"`
}

type MonetizationRulesModel struct {
	ID     types.String                    `tfsdk:"id" json:"-,computed"`
	ZoneID types.String                    `tfsdk:"zone_id" path:"zone_id,required"`
	Rules  *[]*MonetizationRulesRulesModel `tfsdk:"rules" json:"rules,required"`
}

func (m MonetizationRulesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MonetizationRulesModel) MarshalJSONForUpdate(state MonetizationRulesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type MonetizationRulesRulesModel struct {
	Address     types.String `tfsdk:"address" json:"address,required"`
	Expression  types.String `tfsdk:"expression" json:"expression,required"`
	Price       types.String `tfsdk:"price" json:"price,optional"`
	Scheme      types.String `tfsdk:"scheme" json:"scheme,required"`
	ID          types.String `tfsdk:"id" json:"id,optional"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
}
