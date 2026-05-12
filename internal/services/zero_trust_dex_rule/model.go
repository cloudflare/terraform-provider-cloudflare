// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDEXRuleResultEnvelope struct {
	Result ZeroTrustDEXRuleModel `json:"result"`
}

type ZeroTrustDEXRuleModel struct {
	ID            types.String                                                     `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                                     `tfsdk:"account_id" path:"account_id,required"`
	Match         types.String                                                     `tfsdk:"match" json:"match,required"`
	Name          types.String                                                     `tfsdk:"name" json:"name,required"`
	Description   types.String                                                     `tfsdk:"description" json:"description,optional"`
	CreatedAt     types.String                                                     `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt     types.String                                                     `tfsdk:"updated_at" json:"updated_at,computed"`
	TargetedTests customfield.NestedObjectList[ZeroTrustDEXRuleTargetedTestsModel] `tfsdk:"targeted_tests" json:"targeted_tests,computed"`
}

func (m ZeroTrustDEXRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDEXRuleModel) MarshalJSONForUpdate(state ZeroTrustDEXRuleModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ZeroTrustDEXRuleTargetedTestsModel struct {
	Data    customfield.NestedObject[ZeroTrustDEXRuleTargetedTestsDataModel] `tfsdk:"data" json:"data,computed"`
	Enabled types.Bool                                                       `tfsdk:"enabled" json:"enabled,computed"`
	Name    types.String                                                     `tfsdk:"name" json:"name,computed"`
	TestID  types.String                                                     `tfsdk:"test_id" json:"test_id,computed"`
}

type ZeroTrustDEXRuleTargetedTestsDataModel struct {
	Host   types.String `tfsdk:"host" json:"host,computed"`
	Kind   types.String `tfsdk:"kind" json:"kind,computed"`
	Method types.String `tfsdk:"method" json:"method,computed"`
}
