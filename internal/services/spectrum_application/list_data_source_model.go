// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationsResultListDataSourceEnvelope struct {
	Result *[]*SpectrumApplicationsResultDataSourceModel `json:"result,computed"`
}

type SpectrumApplicationsDataSourceModel struct {
	Zone      types.String                                  `tfsdk:"zone" path:"zone"`
	Direction types.String                                  `tfsdk:"direction" query:"direction"`
	Order     types.String                                  `tfsdk:"order" query:"order"`
	MaxItems  types.Int64                                   `tfsdk:"max_items"`
	Result    *[]*SpectrumApplicationsResultDataSourceModel `tfsdk:"result"`
}

type SpectrumApplicationsResultDataSourceModel struct {
}
