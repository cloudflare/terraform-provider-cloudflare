// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package static_route

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StaticRouteResultEnvelope struct {
	Result StaticRouteModel `json:"result,computed"`
}

type StaticRouteModel struct {
	AccountID     types.String               `tfsdk:"account_id" path:"account_id"`
	RouteID       types.String               `tfsdk:"route_id" path:"route_id"`
	Routes        *[]*StaticRouteRoutesModel `tfsdk:"routes" json:"routes,computed"`
	Modified      types.Bool                 `tfsdk:"modified" json:"modified,computed"`
	ModifiedRoute types.String               `tfsdk:"modified_route" json:"modified_route,computed"`
	Route         types.String               `tfsdk:"route" json:"route,computed"`
	Deleted       types.Bool                 `tfsdk:"deleted" json:"deleted,computed"`
	DeletedRoute  types.String               `tfsdk:"deleted_route" json:"deleted_route,computed"`
}

type StaticRouteRoutesModel struct {
	Nexthop     types.String                 `tfsdk:"nexthop" json:"nexthop"`
	Prefix      types.String                 `tfsdk:"prefix" json:"prefix"`
	Priority    types.Int64                  `tfsdk:"priority" json:"priority"`
	ID          types.String                 `tfsdk:"id" json:"id,computed"`
	CreatedOn   types.String                 `tfsdk:"created_on" json:"created_on,computed"`
	Description types.String                 `tfsdk:"description" json:"description"`
	ModifiedOn  types.String                 `tfsdk:"modified_on" json:"modified_on,computed"`
	Scope       *StaticRouteRoutesScopeModel `tfsdk:"scope" json:"scope"`
	Weight      types.Int64                  `tfsdk:"weight" json:"weight"`
}

type StaticRouteRoutesScopeModel struct {
	ColoNames   *[]types.String `tfsdk:"colo_names" json:"colo_names"`
	ColoRegions *[]types.String `tfsdk:"colo_regions" json:"colo_regions"`
}
