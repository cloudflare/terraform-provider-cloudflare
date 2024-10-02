// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDEXTestResultEnvelope struct {
	Result ZeroTrustDEXTestModel `json:"result"`
}

type ZeroTrustDEXTestModel struct {
	ID             types.String                                                      `tfsdk:"id" json:"-,computed"`
	Name           types.String                                                      `tfsdk:"name" json:"name,required"`
	AccountID      types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	Enabled        types.Bool                                                        `tfsdk:"enabled" json:"enabled,required"`
	Interval       types.String                                                      `tfsdk:"interval" json:"interval,required"`
	Data           *ZeroTrustDEXTestDataModel                                        `tfsdk:"data" json:"data,required"`
	Description    types.String                                                      `tfsdk:"description" json:"description,optional"`
	Targeted       types.Bool                                                        `tfsdk:"targeted" json:"targeted,optional"`
	TargetPolicies customfield.NestedObjectList[ZeroTrustDEXTestTargetPoliciesModel] `tfsdk:"target_policies" json:"target_policies,computed_optional"`
}

func (m ZeroTrustDEXTestModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDEXTestModel) MarshalJSONForUpdate(state ZeroTrustDEXTestModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDEXTestDataModel struct {
	Host   types.String `tfsdk:"host" json:"host,optional"`
	Kind   types.String `tfsdk:"kind" json:"kind,optional"`
	Method types.String `tfsdk:"method" json:"method,optional"`
}

type ZeroTrustDEXTestTargetPoliciesModel struct {
	ID      types.String `tfsdk:"id" json:"id,optional"`
	Default types.Bool   `tfsdk:"default" json:"default,optional"`
	Name    types.String `tfsdk:"name" json:"name,optional"`
}
