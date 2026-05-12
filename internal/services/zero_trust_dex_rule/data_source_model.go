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

type ZeroTrustDEXRuleResultDataSourceEnvelope struct {
	Result ZeroTrustDEXRuleDataSourceModel `json:"result,computed"`
}

type ZeroTrustDEXRuleDataSourceModel struct {
	ID            types.String                                                               `tfsdk:"id" path:"rule_id,computed"`
	RuleID        types.String                                                               `tfsdk:"rule_id" path:"rule_id,required"`
	AccountID     types.String                                                               `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt     types.String                                                               `tfsdk:"created_at" json:"created_at,computed"`
	Description   types.String                                                               `tfsdk:"description" json:"description,computed"`
	Match         types.String                                                               `tfsdk:"match" json:"match,computed"`
	Name          types.String                                                               `tfsdk:"name" json:"name,computed"`
	UpdatedAt     types.String                                                               `tfsdk:"updated_at" json:"updated_at,computed"`
	TargetedTests customfield.NestedObjectList[ZeroTrustDEXRuleTargetedTestsDataSourceModel] `tfsdk:"targeted_tests" json:"targeted_tests,computed"`
}

func (m *ZeroTrustDEXRuleDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DEXRuleGetParams, diags diag.Diagnostics) {
	params = zero_trust.DEXRuleGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDEXRuleTargetedTestsDataSourceModel struct {
	Data    customfield.NestedObject[ZeroTrustDEXRuleTargetedTestsDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
	Enabled types.Bool                                                                 `tfsdk:"enabled" json:"enabled,computed"`
	Name    types.String                                                               `tfsdk:"name" json:"name,computed"`
	TestID  types.String                                                               `tfsdk:"test_id" json:"test_id,computed"`
}

type ZeroTrustDEXRuleTargetedTestsDataDataSourceModel struct {
	Host   types.String `tfsdk:"host" json:"host,computed"`
	Kind   types.String `tfsdk:"kind" json:"kind,computed"`
	Method types.String `tfsdk:"method" json:"method,computed"`
}
