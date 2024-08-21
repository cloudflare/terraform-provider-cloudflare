// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDEXTestResultDataSourceEnvelope struct {
	Result ZeroTrustDEXTestDataSourceModel `json:"result,computed"`
}

type ZeroTrustDEXTestResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustDEXTestDataSourceModel `json:"result,computed"`
}

type ZeroTrustDEXTestDataSourceModel struct {
	AccountID      types.String                                                  `tfsdk:"account_id" path:"account_id"`
	DEXTestID      types.String                                                  `tfsdk:"dex_test_id" path:"dex_test_id"`
	Enabled        types.Bool                                                    `tfsdk:"enabled" json:"enabled,computed"`
	Interval       types.String                                                  `tfsdk:"interval" json:"interval,computed"`
	Name           types.String                                                  `tfsdk:"name" json:"name,computed"`
	Data           customfield.NestedObject[ZeroTrustDEXTestDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
	Description    types.String                                                  `tfsdk:"description" json:"description"`
	Targeted       types.Bool                                                    `tfsdk:"targeted" json:"targeted"`
	TargetPolicies *[]*ZeroTrustDEXTestTargetPoliciesDataSourceModel             `tfsdk:"target_policies" json:"target_policies"`
	Filter         *ZeroTrustDEXTestFindOneByDataSourceModel                     `tfsdk:"filter"`
}

func (m *ZeroTrustDEXTestDataSourceModel) toReadParams() (params zero_trust.DeviceDEXTestGetParams, diags diag.Diagnostics) {
	params = zero_trust.DeviceDEXTestGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustDEXTestDataSourceModel) toListParams() (params zero_trust.DeviceDEXTestListParams, diags diag.Diagnostics) {
	params = zero_trust.DeviceDEXTestListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDEXTestDataDataSourceModel struct {
	Host   types.String `tfsdk:"host" json:"host"`
	Kind   types.String `tfsdk:"kind" json:"kind"`
	Method types.String `tfsdk:"method" json:"method"`
}

type ZeroTrustDEXTestTargetPoliciesDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Default types.Bool   `tfsdk:"default" json:"default"`
	Name    types.String `tfsdk:"name" json:"name"`
}

type ZeroTrustDEXTestFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
