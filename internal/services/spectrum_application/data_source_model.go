// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationResultDataSourceEnvelope struct {
	Result SpectrumApplicationDataSourceModel `json:"result,computed"`
}

type SpectrumApplicationResultListDataSourceEnvelope struct {
	Result *[]*SpectrumApplicationDataSourceModel `json:"result,computed"`
}

type SpectrumApplicationDataSourceModel struct {
	AppID  types.String                                 `tfsdk:"app_id" path:"app_id"`
	Zone   types.String                                 `tfsdk:"zone" path:"zone"`
	Filter *SpectrumApplicationFindOneByDataSourceModel `tfsdk:"filter"`
}

type SpectrumApplicationFindOneByDataSourceModel struct {
	Zone      types.String  `tfsdk:"zone" path:"zone"`
	Direction types.String  `tfsdk:"direction" query:"direction"`
	Order     types.String  `tfsdk:"order" query:"order"`
	Page      types.Float64 `tfsdk:"page" query:"page"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page"`
}
