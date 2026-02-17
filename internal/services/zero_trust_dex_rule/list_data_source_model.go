// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDEXRulesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[ZeroTrustDEXRulesResultDataSourceModel] `json:"items,computed"`
}

type ZeroTrustDEXRulesDataSourceModel struct {
	AccountID types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	Name      types.String                                                         `tfsdk:"name" query:"name,optional"`
	SortBy    types.String                                                         `tfsdk:"sort_by" query:"sort_by,computed_optional"`
	SortOrder types.String                                                         `tfsdk:"sort_order" query:"sort_order,computed_optional"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDEXRulesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDEXRulesDataSourceModel) toListParams(_ context.Context) (params zero_trust.DEXRuleListParams, diags diag.Diagnostics) {
	params = zero_trust.DEXRuleListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.SortBy.IsNull() {
		params.SortBy = cloudflare.F(zero_trust.DEXRuleListParamsSortBy(m.SortBy.ValueString()))
	}
	if !m.SortOrder.IsNull() {
		params.SortOrder = cloudflare.F(zero_trust.DEXRuleListParamsSortOrder(m.SortOrder.ValueString()))
	}

	return
}

type ZeroTrustDEXRulesResultDataSourceModel struct {
	Rules customfield.NestedObjectList[ZeroTrustDEXRulesRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
}

type ZeroTrustDEXRulesRulesDataSourceModel struct {
	ID            types.String                                                                     `tfsdk:"id" json:"id,computed"`
	CreatedAt     types.String                                                                     `tfsdk:"created_at" json:"created_at,computed"`
	Match         types.String                                                                     `tfsdk:"match" json:"match,computed"`
	Name          types.String                                                                     `tfsdk:"name" json:"name,computed"`
	Description   types.String                                                                     `tfsdk:"description" json:"description,computed"`
	TargetedTests customfield.NestedObjectList[ZeroTrustDEXRulesRulesTargetedTestsDataSourceModel] `tfsdk:"targeted_tests" json:"targeted_tests,computed"`
	UpdatedAt     types.String                                                                     `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustDEXRulesRulesTargetedTestsDataSourceModel struct {
	Data    customfield.NestedObject[ZeroTrustDEXRulesRulesTargetedTestsDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
	Enabled types.Bool                                                                       `tfsdk:"enabled" json:"enabled,computed"`
	Name    types.String                                                                     `tfsdk:"name" json:"name,computed"`
	TestID  types.String                                                                     `tfsdk:"test_id" json:"test_id,computed"`
}

type ZeroTrustDEXRulesRulesTargetedTestsDataDataSourceModel struct {
	Host   types.String `tfsdk:"host" json:"host,computed"`
	Kind   types.String `tfsdk:"kind" json:"kind,computed"`
	Method types.String `tfsdk:"method" json:"method,computed"`
}
