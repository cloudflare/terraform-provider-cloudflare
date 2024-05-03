// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ObservatoryScheduledTestResultEnvelope struct {
	Result ObservatoryScheduledTestModel `json:"result,computed"`
}

type ObservatoryScheduledTestModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	URL    types.String `tfsdk:"url" path:"url"`
}
