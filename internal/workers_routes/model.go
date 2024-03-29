// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_routes

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersRoutesResultEnvelope struct {
	Result WorkersRoutesModel `json:"result,computed"`
}

type WorkersRoutesModel struct {
	ZoneID  types.String `tfsdk:"zone_id" path:"zone_id"`
	RouteID types.String `tfsdk:"route_id" path:"route_id"`
	Pattern types.String `tfsdk:"pattern" json:"pattern"`
	Script  types.String `tfsdk:"script" json:"script"`
}
