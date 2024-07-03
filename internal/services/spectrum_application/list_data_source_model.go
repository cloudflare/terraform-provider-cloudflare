// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationsResultListDataSourceEnvelope struct {
	Result *[]*SpectrumApplicationsItemsDataSourceModel `json:"result,computed"`
}

type SpectrumApplicationsDataSourceModel struct {
	Zone      types.String                                 `tfsdk:"zone" path:"zone"`
	Direction types.String                                 `tfsdk:"direction" query:"direction"`
	Order     types.String                                 `tfsdk:"order" query:"order"`
	Page      types.Float64                                `tfsdk:"page" query:"page"`
	PerPage   types.Float64                                `tfsdk:"per_page" query:"per_page"`
	MaxItems  types.Int64                                  `tfsdk:"max_items"`
	Items     *[]*SpectrumApplicationsItemsDataSourceModel `tfsdk:"items"`
}

type SpectrumApplicationsItemsDataSourceModel struct {
}
