// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_dex_test

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DeviceDEXTestsResultListDataSourceEnvelope struct {
	Result *[]*DeviceDEXTestsItemsDataSourceModel `json:"result,computed"`
}

type DeviceDEXTestsDataSourceModel struct {
	AccountID types.String                           `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                            `tfsdk:"max_items"`
	Items     *[]*DeviceDEXTestsItemsDataSourceModel `tfsdk:"items"`
}

type DeviceDEXTestsItemsDataSourceModel struct {
	Enabled        types.Bool                                           `tfsdk:"enabled" json:"enabled,computed"`
	Interval       types.String                                         `tfsdk:"interval" json:"interval,computed"`
	Name           types.String                                         `tfsdk:"name" json:"name,computed"`
	Description    types.String                                         `tfsdk:"description" json:"description,computed"`
	TargetPolicies *[]*DeviceDEXTestsItemsTargetPoliciesDataSourceModel `tfsdk:"target_policies" json:"target_policies,computed"`
	Targeted       types.Bool                                           `tfsdk:"targeted" json:"targeted,computed"`
}

type DeviceDEXTestsItemsDataDataSourceModel struct {
	Host   types.String `tfsdk:"host" json:"host,computed"`
	Kind   types.String `tfsdk:"kind" json:"kind,computed"`
	Method types.String `tfsdk:"method" json:"method,computed"`
}

type DeviceDEXTestsItemsTargetPoliciesDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Default types.Bool   `tfsdk:"default" json:"default,computed"`
	Name    types.String `tfsdk:"name" json:"name,computed"`
}
