// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_dex_test

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DeviceDEXTestResultEnvelope struct {
	Result DeviceDEXTestModel `json:"result,computed"`
}

type DeviceDEXTestModel struct {
	ID             types.String                         `tfsdk:"id" json:"-,computed"`
	Name           types.String                         `tfsdk:"name" json:"name"`
	AccountID      types.String                         `tfsdk:"account_id" path:"account_id"`
	Enabled        types.Bool                           `tfsdk:"enabled" json:"enabled"`
	Interval       types.String                         `tfsdk:"interval" json:"interval"`
	Data           *DeviceDEXTestDataModel              `tfsdk:"data" json:"data"`
	Description    types.String                         `tfsdk:"description" json:"description"`
	Targeted       types.Bool                           `tfsdk:"targeted" json:"targeted"`
	TargetPolicies *[]*DeviceDEXTestTargetPoliciesModel `tfsdk:"target_policies" json:"target_policies"`
}

type DeviceDEXTestDataModel struct {
	Host   types.String `tfsdk:"host" json:"host"`
	Kind   types.String `tfsdk:"kind" json:"kind"`
	Method types.String `tfsdk:"method" json:"method"`
}

type DeviceDEXTestTargetPoliciesModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Default types.Bool   `tfsdk:"default" json:"default"`
	Name    types.String `tfsdk:"name" json:"name"`
}
