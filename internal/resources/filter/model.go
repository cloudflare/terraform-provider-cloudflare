// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterResultEnvelope struct {
	Result FilterModel `json:"result,computed"`
}

type FilterModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String `tfsdk:"id" path:"id"`
}
