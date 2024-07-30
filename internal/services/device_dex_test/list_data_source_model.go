// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_dex_test

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DeviceDEXTestsResultListDataSourceEnvelope struct {
	Result *[]*DeviceDEXTestsResultDataSourceModel `json:"result,computed"`
}

type DeviceDEXTestsDataSourceModel struct {
	AccountID types.String                            `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                             `tfsdk:"max_items"`
	Result    *[]*DeviceDEXTestsResultDataSourceModel `tfsdk:"result"`
}

type DeviceDEXTestsResultDataSourceModel struct {
	Data           customfield.NestedObject[DeviceDEXTestsDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
	Enabled        types.Bool                                                  `tfsdk:"enabled" json:"enabled,computed"`
	Interval       types.String                                                `tfsdk:"interval" json:"interval,computed"`
	Name           types.String                                                `tfsdk:"name" json:"name,computed"`
	Description    types.String                                                `tfsdk:"description" json:"description"`
	TargetPolicies *[]*DeviceDEXTestsTargetPoliciesDataSourceModel             `tfsdk:"target_policies" json:"target_policies"`
	Targeted       types.Bool                                                  `tfsdk:"targeted" json:"targeted"`
}

type DeviceDEXTestsDataDataSourceModel struct {
	Host   types.String `tfsdk:"host" json:"host"`
	Kind   types.String `tfsdk:"kind" json:"kind"`
	Method types.String `tfsdk:"method" json:"method"`
}

type DeviceDEXTestsTargetPoliciesDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Default types.Bool   `tfsdk:"default" json:"default"`
	Name    types.String `tfsdk:"name" json:"name"`
}
