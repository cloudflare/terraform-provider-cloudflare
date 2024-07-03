// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_dex_test

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DeviceDEXTestResultDataSourceEnvelope struct {
	Result DeviceDEXTestDataSourceModel `json:"result,computed"`
}

type DeviceDEXTestResultListDataSourceEnvelope struct {
	Result *[]*DeviceDEXTestDataSourceModel `json:"result,computed"`
}

type DeviceDEXTestDataSourceModel struct {
	AccountID      types.String                                   `tfsdk:"account_id" path:"account_id"`
	DEXTestID      types.String                                   `tfsdk:"dex_test_id" path:"dex_test_id"`
	Data           *DeviceDEXTestDataDataSourceModel              `tfsdk:"data" json:"data"`
	Enabled        types.Bool                                     `tfsdk:"enabled" json:"enabled"`
	Interval       types.String                                   `tfsdk:"interval" json:"interval"`
	Name           types.String                                   `tfsdk:"name" json:"name"`
	Description    types.String                                   `tfsdk:"description" json:"description"`
	TargetPolicies *[]*DeviceDEXTestTargetPoliciesDataSourceModel `tfsdk:"target_policies" json:"target_policies"`
	Targeted       types.Bool                                     `tfsdk:"targeted" json:"targeted"`
	FindOneBy      *DeviceDEXTestFindOneByDataSourceModel         `tfsdk:"find_one_by"`
}

type DeviceDEXTestDataDataSourceModel struct {
	Host   types.String `tfsdk:"host" json:"host"`
	Kind   types.String `tfsdk:"kind" json:"kind"`
	Method types.String `tfsdk:"method" json:"method"`
}

type DeviceDEXTestTargetPoliciesDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Default types.Bool   `tfsdk:"default" json:"default"`
	Name    types.String `tfsdk:"name" json:"name"`
}

type DeviceDEXTestFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
