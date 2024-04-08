// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_devices_dex_tests

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDevicesDEXTestsResultEnvelope struct {
	Result ZeroTrustDevicesDEXTestsModel `json:"result,computed"`
}

type ZeroTrustDevicesDEXTestsModel struct {
	AccountID      types.String                                   `tfsdk:"account_id" path:"account_id"`
	DEXTestID      types.String                                   `tfsdk:"dex_test_id" path:"dex_test_id"`
	Data           *ZeroTrustDevicesDEXTestsDataModel             `tfsdk:"data" json:"data"`
	Enabled        types.Bool                                     `tfsdk:"enabled" json:"enabled"`
	Interval       types.String                                   `tfsdk:"interval" json:"interval"`
	Name           types.String                                   `tfsdk:"name" json:"name"`
	Description    types.String                                   `tfsdk:"description" json:"description"`
	TargetPolicies []*ZeroTrustDevicesDEXTestsTargetPoliciesModel `tfsdk:"target_policies" json:"target_policies"`
	Targeted       types.Bool                                     `tfsdk:"targeted" json:"targeted"`
}

type ZeroTrustDevicesDEXTestsDataModel struct {
	Host   types.String `tfsdk:"host" json:"host"`
	Kind   types.String `tfsdk:"kind" json:"kind"`
	Method types.String `tfsdk:"method" json:"method"`
}

type ZeroTrustDevicesDEXTestsTargetPoliciesModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Default types.Bool   `tfsdk:"default" json:"default"`
	Name    types.String `tfsdk:"name" json:"name"`
}
