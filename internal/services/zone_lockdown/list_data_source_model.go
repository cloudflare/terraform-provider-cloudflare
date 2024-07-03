// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneLockdownsResultListDataSourceEnvelope struct {
	Result *[]*ZoneLockdownsItemsDataSourceModel `json:"result,computed"`
}

type ZoneLockdownsDataSourceModel struct {
	ZoneIdentifier    types.String                          `tfsdk:"zone_identifier" path:"zone_identifier"`
	CreatedOn         types.String                          `tfsdk:"created_on" query:"created_on"`
	Description       types.String                          `tfsdk:"description" query:"description"`
	DescriptionSearch types.String                          `tfsdk:"description_search" query:"description_search"`
	IP                types.String                          `tfsdk:"ip" query:"ip"`
	IPRangeSearch     types.String                          `tfsdk:"ip_range_search" query:"ip_range_search"`
	IPSearch          types.String                          `tfsdk:"ip_search" query:"ip_search"`
	ModifiedOn        types.String                          `tfsdk:"modified_on" query:"modified_on"`
	Page              types.Float64                         `tfsdk:"page" query:"page"`
	PerPage           types.Float64                         `tfsdk:"per_page" query:"per_page"`
	Priority          types.Float64                         `tfsdk:"priority" query:"priority"`
	URISearch         types.String                          `tfsdk:"uri_search" query:"uri_search"`
	MaxItems          types.Int64                           `tfsdk:"max_items"`
	Items             *[]*ZoneLockdownsItemsDataSourceModel `tfsdk:"items"`
}

type ZoneLockdownsItemsDataSourceModel struct {
	ID          types.String    `tfsdk:"id" json:"id,computed"`
	CreatedOn   types.String    `tfsdk:"created_on" json:"created_on,computed"`
	Description types.String    `tfsdk:"description" json:"description,computed"`
	ModifiedOn  types.String    `tfsdk:"modified_on" json:"modified_on,computed"`
	Paused      types.Bool      `tfsdk:"paused" json:"paused,computed"`
	URLs        *[]types.String `tfsdk:"urls" json:"urls,computed"`
}

type ZoneLockdownsItemsConfigurationsDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target,computed"`
	Value  types.String `tfsdk:"value" json:"value,computed"`
}
