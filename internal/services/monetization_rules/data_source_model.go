// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monetization_rules

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/monetization"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MonetizationRulesResultDataSourceEnvelope struct {
	Result MonetizationRulesDataSourceModel `json:"result,computed"`
}

type MonetizationRulesDataSourceModel struct {
	ID     types.String                                                        `tfsdk:"id" path:"zone_id,computed"`
	ZoneID types.String                                                        `tfsdk:"zone_id" path:"zone_id,required"`
	Rules  customfield.NestedObjectList[MonetizationRulesRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
}

func (m *MonetizationRulesDataSourceModel) toReadParams(_ context.Context) (params monetization.RuleGetParams, diags diag.Diagnostics) {
	params = monetization.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type MonetizationRulesRulesDataSourceModel struct {
	Address     types.String `tfsdk:"address" json:"address,computed"`
	Expression  types.String `tfsdk:"expression" json:"expression,computed"`
	Price       types.String `tfsdk:"price" json:"price,computed"`
	Scheme      types.String `tfsdk:"scheme" json:"scheme,computed"`
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}
