// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDEXTestsResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustDEXTestsResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustDEXTestsDataSourceModel struct {
	AccountID types.String                               `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                `tfsdk:"max_items"`
	Result    *[]*ZeroTrustDEXTestsResultDataSourceModel `tfsdk:"result"`
}

type ZeroTrustDEXTestsResultDataSourceModel struct {
	Data           customfield.NestedObject[ZeroTrustDEXTestsDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
	Enabled        types.Bool                                                     `tfsdk:"enabled" json:"enabled,computed"`
	Interval       types.String                                                   `tfsdk:"interval" json:"interval,computed"`
	Name           types.String                                                   `tfsdk:"name" json:"name,computed"`
	Description    types.String                                                   `tfsdk:"description" json:"description"`
	TargetPolicies *[]*ZeroTrustDEXTestsTargetPoliciesDataSourceModel             `tfsdk:"target_policies" json:"target_policies"`
	Targeted       types.Bool                                                     `tfsdk:"targeted" json:"targeted"`
}

type ZeroTrustDEXTestsDataDataSourceModel struct {
	Host   types.String `tfsdk:"host" json:"host"`
	Kind   types.String `tfsdk:"kind" json:"kind"`
	Method types.String `tfsdk:"method" json:"method"`
}

type ZeroTrustDEXTestsTargetPoliciesDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Default types.Bool   `tfsdk:"default" json:"default"`
	Name    types.String `tfsdk:"name" json:"name"`
}
