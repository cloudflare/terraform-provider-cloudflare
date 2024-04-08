// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filters

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FiltersResultEnvelope struct {
	Result FiltersModel `json:"result,computed"`
}

type FiltersModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String `tfsdk:"id" path:"id"`
	Expression     types.String `tfsdk:"expression" json:"expression"`
	Paused         types.Bool   `tfsdk:"paused" json:"paused"`
	Description    types.String `tfsdk:"description" json:"description"`
	Ref            types.String `tfsdk:"ref" json:"ref"`
}
