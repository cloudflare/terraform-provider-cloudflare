// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_static_route

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANStaticRouteResultEnvelope struct {
	Result MagicWANStaticRouteModel `json:"result"`
}

type MagicWANStaticRouteModel struct {
	AccountID     types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	RouteID       types.String                                                    `tfsdk:"route_id" path:"route_id,optional"`
	Description   types.String                                                    `tfsdk:"description" json:"description,optional"`
	Nexthop       types.String                                                    `tfsdk:"nexthop" json:"nexthop,optional"`
	Prefix        types.String                                                    `tfsdk:"prefix" json:"prefix,optional"`
	Priority      types.Int64                                                     `tfsdk:"priority" json:"priority,optional"`
	Weight        types.Int64                                                     `tfsdk:"weight" json:"weight,optional"`
	Scope         customfield.NestedObject[MagicWANStaticRouteScopeModel]         `tfsdk:"scope" json:"scope,computed_optional"`
	Modified      types.Bool                                                      `tfsdk:"modified" json:"modified,computed"`
	ModifiedRoute customfield.NestedObject[MagicWANStaticRouteModifiedRouteModel] `tfsdk:"modified_route" json:"modified_route,computed"`
	Route         customfield.NestedObject[MagicWANStaticRouteRouteModel]         `tfsdk:"route" json:"route,computed"`
	Routes        customfield.NestedObjectList[MagicWANStaticRouteRoutesModel]    `tfsdk:"routes" json:"routes,computed"`
}

func (m MagicWANStaticRouteModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MagicWANStaticRouteModel) MarshalJSONForUpdate(state MagicWANStaticRouteModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type MagicWANStaticRouteScopeModel struct {
	ColoNames   *[]types.String `tfsdk:"colo_names" json:"colo_names,optional"`
	ColoRegions *[]types.String `tfsdk:"colo_regions" json:"colo_regions,optional"`
}

type MagicWANStaticRouteModifiedRouteModel struct {
	Nexthop     types.String                                                         `tfsdk:"nexthop" json:"nexthop,computed"`
	Prefix      types.String                                                         `tfsdk:"prefix" json:"prefix,computed"`
	Priority    types.Int64                                                          `tfsdk:"priority" json:"priority,computed"`
	ID          types.String                                                         `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                                                    `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String                                                         `tfsdk:"description" json:"description,computed"`
	ModifiedOn  timetypes.RFC3339                                                    `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Scope       customfield.NestedObject[MagicWANStaticRouteModifiedRouteScopeModel] `tfsdk:"scope" json:"scope,computed"`
	Weight      types.Int64                                                          `tfsdk:"weight" json:"weight,computed"`
}

type MagicWANStaticRouteModifiedRouteScopeModel struct {
	ColoNames   customfield.List[types.String] `tfsdk:"colo_names" json:"colo_names,computed"`
	ColoRegions customfield.List[types.String] `tfsdk:"colo_regions" json:"colo_regions,computed"`
}

type MagicWANStaticRouteRouteModel struct {
	Nexthop     types.String                                                 `tfsdk:"nexthop" json:"nexthop,computed"`
	Prefix      types.String                                                 `tfsdk:"prefix" json:"prefix,computed"`
	Priority    types.Int64                                                  `tfsdk:"priority" json:"priority,computed"`
	ID          types.String                                                 `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                                            `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String                                                 `tfsdk:"description" json:"description,computed"`
	ModifiedOn  timetypes.RFC3339                                            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Scope       customfield.NestedObject[MagicWANStaticRouteRouteScopeModel] `tfsdk:"scope" json:"scope,computed"`
	Weight      types.Int64                                                  `tfsdk:"weight" json:"weight,computed"`
}

type MagicWANStaticRouteRouteScopeModel struct {
	ColoNames   customfield.List[types.String] `tfsdk:"colo_names" json:"colo_names,computed"`
	ColoRegions customfield.List[types.String] `tfsdk:"colo_regions" json:"colo_regions,computed"`
}

type MagicWANStaticRouteRoutesModel struct {
	Nexthop     types.String                                                  `tfsdk:"nexthop" json:"nexthop,computed"`
	Prefix      types.String                                                  `tfsdk:"prefix" json:"prefix,computed"`
	Priority    types.Int64                                                   `tfsdk:"priority" json:"priority,computed"`
	ID          types.String                                                  `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                                             `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String                                                  `tfsdk:"description" json:"description,computed"`
	ModifiedOn  timetypes.RFC3339                                             `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Scope       customfield.NestedObject[MagicWANStaticRouteRoutesScopeModel] `tfsdk:"scope" json:"scope,computed"`
	Weight      types.Int64                                                   `tfsdk:"weight" json:"weight,computed"`
}

type MagicWANStaticRouteRoutesScopeModel struct {
	ColoNames   customfield.List[types.String] `tfsdk:"colo_names" json:"colo_names,computed"`
	ColoRegions customfield.List[types.String] `tfsdk:"colo_regions" json:"colo_regions,computed"`
}
