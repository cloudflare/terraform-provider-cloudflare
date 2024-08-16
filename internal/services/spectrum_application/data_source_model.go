// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationResultDataSourceEnvelope struct {
	Result SpectrumApplicationDataSourceModel `json:"result,computed"`
}

type SpectrumApplicationDataSourceModel struct {
	AppID  types.String `tfsdk:"app_id" path:"app_id"`
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
