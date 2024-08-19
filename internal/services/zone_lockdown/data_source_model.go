// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneLockdownResultDataSourceEnvelope struct {
	Result ZoneLockdownDataSourceModel `json:"result,computed"`
}

type ZoneLockdownResultListDataSourceEnvelope struct {
	Result *[]*ZoneLockdownDataSourceModel `json:"result,computed"`
}

type ZoneLockdownDataSourceModel struct {
	ZoneIdentifier types.String                                                        `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                                                        `tfsdk:"id" path:"id"`
	CreatedOn      timetypes.RFC3339                                                   `tfsdk:"created_on" json:"created_on,computed"`
	Description    types.String                                                        `tfsdk:"description" json:"description,computed"`
	ModifiedOn     timetypes.RFC3339                                                   `tfsdk:"modified_on" json:"modified_on,computed"`
	Paused         types.Bool                                                          `tfsdk:"paused" json:"paused,computed"`
	URLs           *[]types.String                                                     `tfsdk:"urls" json:"urls,computed"`
	Configurations customfield.NestedObject[ZoneLockdownConfigurationsDataSourceModel] `tfsdk:"configurations" json:"configurations,computed"`
	Filter         *ZoneLockdownFindOneByDataSourceModel                               `tfsdk:"filter"`
}

type ZoneLockdownConfigurationsDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target"`
	Value  types.String `tfsdk:"value" json:"value"`
}

type ZoneLockdownFindOneByDataSourceModel struct {
	ZoneIdentifier    types.String      `tfsdk:"zone_identifier" path:"zone_identifier"`
	CreatedOn         timetypes.RFC3339 `tfsdk:"created_on" query:"created_on"`
	Description       types.String      `tfsdk:"description" query:"description"`
	DescriptionSearch types.String      `tfsdk:"description_search" query:"description_search"`
	IP                types.String      `tfsdk:"ip" query:"ip"`
	IPRangeSearch     types.String      `tfsdk:"ip_range_search" query:"ip_range_search"`
	IPSearch          types.String      `tfsdk:"ip_search" query:"ip_search"`
	ModifiedOn        timetypes.RFC3339 `tfsdk:"modified_on" query:"modified_on"`
	Priority          types.Float64     `tfsdk:"priority" query:"priority"`
	URISearch         types.String      `tfsdk:"uri_search" query:"uri_search"`
}
