// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDEXTestsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDEXTestsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDEXTestsDataSourceModel struct {
	AccountID types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDEXTestsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDEXTestsDataSourceModel) toListParams(_ context.Context) (params zero_trust.DeviceDEXTestListParams, diags diag.Diagnostics) {
	params = zero_trust.DeviceDEXTestListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDEXTestsResultDataSourceModel struct {
	Data           customfield.NestedObject[ZeroTrustDEXTestsDataDataSourceModel]               `tfsdk:"data" json:"data,computed"`
	Enabled        types.Bool                                                                   `tfsdk:"enabled" json:"enabled,computed"`
	Interval       types.String                                                                 `tfsdk:"interval" json:"interval,computed"`
	Name           types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Description    types.String                                                                 `tfsdk:"description" json:"description,computed"`
	TargetPolicies customfield.NestedObjectList[ZeroTrustDEXTestsTargetPoliciesDataSourceModel] `tfsdk:"target_policies" json:"target_policies,computed"`
	Targeted       types.Bool                                                                   `tfsdk:"targeted" json:"targeted,computed"`
}

type ZeroTrustDEXTestsDataDataSourceModel struct {
	Host   types.String `tfsdk:"host" json:"host,computed"`
	Kind   types.String `tfsdk:"kind" json:"kind,computed"`
	Method types.String `tfsdk:"method" json:"method,computed"`
}

type ZeroTrustDEXTestsTargetPoliciesDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Default types.Bool   `tfsdk:"default" json:"default,computed"`
	Name    types.String `tfsdk:"name" json:"name,computed"`
}
