// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_static_route

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANStaticRouteResultEnvelope struct {
	Result MagicWANStaticRouteModel `json:"result"`
}

type MagicWANStaticRouteModel struct {
	ID            types.String                                                    `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	Nexthop       types.String                                                    `tfsdk:"nexthop" json:"nexthop,required,no_refresh"`
	Prefix        types.String                                                    `tfsdk:"prefix" json:"prefix,required,no_refresh"`
	Priority      types.Int64                                                     `tfsdk:"priority" json:"priority,required,no_refresh"`
	Description   types.String                                                    `tfsdk:"description" json:"description,optional,no_refresh"`
	Weight        types.Int64                                                     `tfsdk:"weight" json:"weight,optional,no_refresh"`
	Scope         *MagicWANStaticRouteScopeModel                                  `tfsdk:"scope" json:"scope,optional,no_refresh"`
	CreatedOn     timetypes.RFC3339                                               `tfsdk:"created_on" json:"created_on,computed,no_refresh" format:"date-time"`
	Modified      types.Bool                                                      `tfsdk:"modified" json:"modified,computed,no_refresh"`
	ModifiedOn    timetypes.RFC3339                                               `tfsdk:"modified_on" json:"modified_on,computed,no_refresh" format:"date-time"`
	ModifiedRoute customfield.NestedObject[MagicWANStaticRouteModifiedRouteModel] `tfsdk:"modified_route" json:"modified_route,computed,no_refresh"`
	Route         customfield.NestedObject[MagicWANStaticRouteRouteModel]         `tfsdk:"route" json:"route,computed_optional"`
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
	ID          types.String                                                         `tfsdk:"id" json:"id,computed"`
	Nexthop     types.String                                                         `tfsdk:"nexthop" json:"nexthop,computed"`
	Prefix      types.String                                                         `tfsdk:"prefix" json:"prefix,computed"`
	Priority    types.Int64                                                          `tfsdk:"priority" json:"priority,computed"`
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
	ID          types.String                                                 `tfsdk:"id" json:"id,computed"`
	Nexthop     types.String                                                 `tfsdk:"nexthop" json:"nexthop,computed_optional"`
	Prefix      types.String                                                 `tfsdk:"prefix" json:"prefix,computed_optional"`
	Priority    types.Int64                                                  `tfsdk:"priority" json:"priority,computed_optional"`
	CreatedOn   timetypes.RFC3339                                            `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String                                                 `tfsdk:"description" json:"description,computed_optional"`
	ModifiedOn  timetypes.RFC3339                                            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Scope       customfield.NestedObject[MagicWANStaticRouteRouteScopeModel] `tfsdk:"scope" json:"scope,computed_optional"`
	Weight      types.Int64                                                  `tfsdk:"weight" json:"weight,computed_optional"`
}

type MagicWANStaticRouteRouteScopeModel struct {
	ColoNames   customfield.List[types.String] `tfsdk:"colo_names" json:"colo_names,computed_optional"`
	ColoRegions customfield.List[types.String] `tfsdk:"colo_regions" json:"colo_regions,computed_optional"`
}
