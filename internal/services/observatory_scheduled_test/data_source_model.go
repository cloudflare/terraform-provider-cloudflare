// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ObservatoryScheduledTestResultDataSourceEnvelope struct {
	Result ObservatoryScheduledTestDataSourceModel `json:"result,computed"`
}

type ObservatoryScheduledTestDataSourceModel struct {
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
	URL       types.String `tfsdk:"url" path:"url,computed"`
	Region    types.String `tfsdk:"region" query:"region,computed"`
	Frequency types.String `tfsdk:"frequency" json:"frequency"`
}
