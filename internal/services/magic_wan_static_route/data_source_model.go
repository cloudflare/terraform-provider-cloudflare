// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_static_route

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/magic_transit"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANStaticRouteResultDataSourceEnvelope struct {
	Result MagicWANStaticRouteDataSourceModel `json:"result,computed"`
}

type MagicWANStaticRouteDataSourceModel struct {
	AccountID types.String                             `tfsdk:"account_id" path:"account_id"`
	RouteID   types.String                             `tfsdk:"route_id" path:"route_id"`
	Route     *MagicWANStaticRouteRouteDataSourceModel `tfsdk:"route" json:"route"`
}

func (m *MagicWANStaticRouteDataSourceModel) toReadParams() (params magic_transit.RouteGetParams, diags diag.Diagnostics) {
	params = magic_transit.RouteGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicWANStaticRouteRouteDataSourceModel struct {
	Nexthop     types.String                                  `tfsdk:"nexthop" json:"nexthop,computed"`
	Prefix      types.String                                  `tfsdk:"prefix" json:"prefix,computed"`
	Priority    types.Int64                                   `tfsdk:"priority" json:"priority,computed"`
	ID          types.String                                  `tfsdk:"id" json:"id,computed_optional"`
	CreatedOn   timetypes.RFC3339                             `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String                                  `tfsdk:"description" json:"description,computed_optional"`
	ModifiedOn  timetypes.RFC3339                             `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Scope       *MagicWANStaticRouteRouteScopeDataSourceModel `tfsdk:"scope" json:"scope,computed_optional"`
	Weight      types.Int64                                   `tfsdk:"weight" json:"weight,computed_optional"`
}

type MagicWANStaticRouteRouteScopeDataSourceModel struct {
	ColoNames   *[]types.String `tfsdk:"colo_names" json:"colo_names,computed_optional"`
	ColoRegions *[]types.String `tfsdk:"colo_regions" json:"colo_regions,computed_optional"`
}
