// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/rulesets"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RulesetsResultDataSourceModel] `json:"result,computed"`
}

type RulesetsDataSourceModel struct {
	AccountID types.String                                                `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                `tfsdk:"zone_id" path:"zone_id,optional"`
	MaxItems  types.Int64                                                 `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[RulesetsResultDataSourceModel] `tfsdk:"result"`
}

func (m *RulesetsDataSourceModel) toListParams(_ context.Context) (params rulesets.RulesetListParams, diags diag.Diagnostics) {
	params = rulesets.RulesetListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type RulesetsResultDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Kind        types.String `tfsdk:"kind" json:"kind,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Phase       types.String `tfsdk:"phase" json:"phase,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
}
