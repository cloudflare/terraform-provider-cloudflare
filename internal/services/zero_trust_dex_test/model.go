// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDEXTestResultEnvelope struct {
	Result ZeroTrustDEXTestModel `json:"result"`
}

type ZeroTrustDEXTestModel struct {
	ID             types.String                                                      `tfsdk:"id" json:"-,computed"`
	Name           types.String                                                      `tfsdk:"name" json:"name,computed_optional"`
	AccountID      types.String                                                      `tfsdk:"account_id" path:"account_id"`
	Description    types.String                                                      `tfsdk:"description" json:"description,computed_optional"`
	Enabled        types.Bool                                                        `tfsdk:"enabled" json:"enabled,computed_optional"`
	Interval       types.String                                                      `tfsdk:"interval" json:"interval,computed_optional"`
	Targeted       types.Bool                                                        `tfsdk:"targeted" json:"targeted,computed_optional"`
	Data           customfield.NestedObject[ZeroTrustDEXTestDataModel]               `tfsdk:"data" json:"data,computed_optional"`
	TargetPolicies customfield.NestedObjectList[ZeroTrustDEXTestTargetPoliciesModel] `tfsdk:"target_policies" json:"target_policies,computed_optional"`
}

type ZeroTrustDEXTestDataModel struct {
	Host   types.String `tfsdk:"host" json:"host,computed_optional"`
	Kind   types.String `tfsdk:"kind" json:"kind,computed_optional"`
	Method types.String `tfsdk:"method" json:"method,computed_optional"`
}

type ZeroTrustDEXTestTargetPoliciesModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed_optional"`
	Default types.Bool   `tfsdk:"default" json:"default,computed_optional"`
	Name    types.String `tfsdk:"name" json:"name,computed_optional"`
}
