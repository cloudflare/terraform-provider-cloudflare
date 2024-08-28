// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneLockdownResultDataSourceEnvelope struct {
	Result ZoneLockdownDataSourceModel `json:"result,computed"`
}

type ZoneLockdownResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZoneLockdownDataSourceModel] `json:"result,computed"`
}

type ZoneLockdownDataSourceModel struct {
	ZoneIdentifier types.String                                                        `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                                                        `tfsdk:"id" path:"id,computed_optional"`
	CreatedOn      timetypes.RFC3339                                                   `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description    types.String                                                        `tfsdk:"description" json:"description,computed"`
	ModifiedOn     timetypes.RFC3339                                                   `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Paused         types.Bool                                                          `tfsdk:"paused" json:"paused,computed"`
	URLs           types.List                                                          `tfsdk:"urls" json:"urls,computed"`
	Configurations customfield.NestedObject[ZoneLockdownConfigurationsDataSourceModel] `tfsdk:"configurations" json:"configurations,computed"`
	Filter         *ZoneLockdownFindOneByDataSourceModel                               `tfsdk:"filter"`
}

func (m *ZoneLockdownDataSourceModel) toListParams() (params firewall.LockdownListParams, diags diag.Diagnostics) {
	mFilterCreatedOn, errs := m.Filter.CreatedOn.ValueRFC3339Time()
	diags.Append(errs...)
	mFilterModifiedOn, errs := m.Filter.ModifiedOn.ValueRFC3339Time()
	diags.Append(errs...)

	params = firewall.LockdownListParams{}

	if !m.Filter.CreatedOn.IsNull() {
		params.CreatedOn = cloudflare.F(mFilterCreatedOn)
	}
	if !m.Filter.Description.IsNull() {
		params.Description = cloudflare.F(m.Filter.Description.ValueString())
	}
	if !m.Filter.DescriptionSearch.IsNull() {
		params.DescriptionSearch = cloudflare.F(m.Filter.DescriptionSearch.ValueString())
	}
	if !m.Filter.IP.IsNull() {
		params.IP = cloudflare.F(m.Filter.IP.ValueString())
	}
	if !m.Filter.IPRangeSearch.IsNull() {
		params.IPRangeSearch = cloudflare.F(m.Filter.IPRangeSearch.ValueString())
	}
	if !m.Filter.IPSearch.IsNull() {
		params.IPSearch = cloudflare.F(m.Filter.IPSearch.ValueString())
	}
	if !m.Filter.ModifiedOn.IsNull() {
		params.ModifiedOn = cloudflare.F(mFilterModifiedOn)
	}
	if !m.Filter.Priority.IsNull() {
		params.Priority = cloudflare.F(m.Filter.Priority.ValueFloat64())
	}
	if !m.Filter.URISearch.IsNull() {
		params.URISearch = cloudflare.F(m.Filter.URISearch.ValueString())
	}

	return
}

type ZoneLockdownConfigurationsDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target,computed"`
	Value  types.String `tfsdk:"value" json:"value,computed"`
}

type ZoneLockdownFindOneByDataSourceModel struct {
	ZoneIdentifier    types.String      `tfsdk:"zone_identifier" path:"zone_identifier"`
	CreatedOn         timetypes.RFC3339 `tfsdk:"created_on" query:"created_on" format:"date-time"`
	Description       types.String      `tfsdk:"description" query:"description"`
	DescriptionSearch types.String      `tfsdk:"description_search" query:"description_search"`
	IP                types.String      `tfsdk:"ip" query:"ip"`
	IPRangeSearch     types.String      `tfsdk:"ip_range_search" query:"ip_range_search"`
	IPSearch          types.String      `tfsdk:"ip_search" query:"ip_search"`
	ModifiedOn        timetypes.RFC3339 `tfsdk:"modified_on" query:"modified_on" format:"date-time"`
	Priority          types.Float64     `tfsdk:"priority" query:"priority"`
	URISearch         types.String      `tfsdk:"uri_search" query:"uri_search"`
}
