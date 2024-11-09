// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_static_route

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANStaticRouteResultDataSourceEnvelope struct {
	Result MagicWANStaticRouteDataSourceModel `json:"result,computed"`
}

type MagicWANStaticRouteDataSourceModel struct {
	AccountID types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	RouteID   types.String                                                      `tfsdk:"route_id" path:"route_id,required"`
	Route     customfield.NestedObject[MagicWANStaticRouteRouteDataSourceModel] `tfsdk:"route" json:"route,computed"`
}

func (m *MagicWANStaticRouteDataSourceModel) toReadParams(_ context.Context) (params magic_transit.RouteGetParams, diags diag.Diagnostics) {
	params = magic_transit.RouteGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicWANStaticRouteRouteDataSourceModel struct {
	Nexthop     types.String                                                           `tfsdk:"nexthop" json:"nexthop,computed"`
	Prefix      types.String                                                           `tfsdk:"prefix" json:"prefix,computed"`
	Priority    types.Int64                                                            `tfsdk:"priority" json:"priority,computed"`
	ID          types.String                                                           `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                                                      `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String                                                           `tfsdk:"description" json:"description,computed"`
	ModifiedOn  timetypes.RFC3339                                                      `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Scope       customfield.NestedObject[MagicWANStaticRouteRouteScopeDataSourceModel] `tfsdk:"scope" json:"scope,computed"`
	Weight      types.Int64                                                            `tfsdk:"weight" json:"weight,computed"`
}

type MagicWANStaticRouteRouteScopeDataSourceModel struct {
	ColoNames   customfield.List[types.String] `tfsdk:"colo_names" json:"colo_names,computed"`
	ColoRegions customfield.List[types.String] `tfsdk:"colo_regions" json:"colo_regions,computed"`
}
