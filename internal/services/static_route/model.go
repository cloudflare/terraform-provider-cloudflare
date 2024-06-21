// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package static_route

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StaticRouteResultEnvelope struct {
	Result StaticRouteModel `json:"result,computed"`
}

type StaticRouteModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	RouteID   types.String `tfsdk:"route_id" path:"route_id"`
}
