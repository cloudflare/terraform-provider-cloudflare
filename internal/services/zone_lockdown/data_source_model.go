// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneLockdownResultDataSourceEnvelope struct {
	Result ZoneLockdownDataSourceModel `json:"result,computed"`
}

type ZoneLockdownResultListDataSourceEnvelope struct {
	Result *[]*ZoneLockdownDataSourceModel `json:"result,computed"`
}

type ZoneLockdownDataSourceModel struct {
	ZoneIdentifier types.String                               `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                               `tfsdk:"id" path:"id"`
	Configurations *ZoneLockdownConfigurationsDataSourceModel `tfsdk:"configurations" json:"configurations"`
	CreatedOn      types.String                               `tfsdk:"created_on" json:"created_on"`
	Description    types.String                               `tfsdk:"description" json:"description"`
	ModifiedOn     types.String                               `tfsdk:"modified_on" json:"modified_on"`
	Paused         types.Bool                                 `tfsdk:"paused" json:"paused"`
	URLs           types.String                               `tfsdk:"urls" json:"urls"`
	FindOneBy      *ZoneLockdownFindOneByDataSourceModel      `tfsdk:"find_one_by"`
}

type ZoneLockdownConfigurationsDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target"`
	Value  types.String `tfsdk:"value" json:"value"`
}

type ZoneLockdownFindOneByDataSourceModel struct {
	ZoneIdentifier    types.String  `tfsdk:"zone_identifier" path:"zone_identifier"`
	CreatedOn         types.String  `tfsdk:"created_on" query:"created_on"`
	Description       types.String  `tfsdk:"description" query:"description"`
	DescriptionSearch types.String  `tfsdk:"description_search" query:"description_search"`
	IP                types.String  `tfsdk:"ip" query:"ip"`
	IPRangeSearch     types.String  `tfsdk:"ip_range_search" query:"ip_range_search"`
	IPSearch          types.String  `tfsdk:"ip_search" query:"ip_search"`
	ModifiedOn        types.String  `tfsdk:"modified_on" query:"modified_on"`
	Page              types.Float64 `tfsdk:"page" query:"page"`
	PerPage           types.Float64 `tfsdk:"per_page" query:"per_page"`
	Priority          types.Float64 `tfsdk:"priority" query:"priority"`
	URISearch         types.String  `tfsdk:"uri_search" query:"uri_search"`
}
