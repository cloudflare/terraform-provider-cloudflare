// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneLockdownsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZoneLockdownsResultDataSourceModel] `json:"result,computed"`
}

type ZoneLockdownsDataSourceModel struct {
	ZoneIdentifier    types.String                                                     `tfsdk:"zone_identifier" path:"zone_identifier,required"`
	CreatedOn         timetypes.RFC3339                                                `tfsdk:"created_on" query:"created_on,optional" format:"date-time"`
	Description       types.String                                                     `tfsdk:"description" query:"description,optional"`
	DescriptionSearch types.String                                                     `tfsdk:"description_search" query:"description_search,optional"`
	IP                types.String                                                     `tfsdk:"ip" query:"ip,optional"`
	IPRangeSearch     types.String                                                     `tfsdk:"ip_range_search" query:"ip_range_search,optional"`
	IPSearch          types.String                                                     `tfsdk:"ip_search" query:"ip_search,optional"`
	ModifiedOn        timetypes.RFC3339                                                `tfsdk:"modified_on" query:"modified_on,optional" format:"date-time"`
	Priority          types.Float64                                                    `tfsdk:"priority" query:"priority,optional"`
	URISearch         types.String                                                     `tfsdk:"uri_search" query:"uri_search,optional"`
	MaxItems          types.Int64                                                      `tfsdk:"max_items"`
	Result            customfield.NestedObjectList[ZoneLockdownsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZoneLockdownsDataSourceModel) toListParams(_ context.Context) (params firewall.LockdownListParams, diags diag.Diagnostics) {
	mCreatedOn, errs := m.CreatedOn.ValueRFC3339Time()
	diags.Append(errs...)
	mModifiedOn, errs := m.ModifiedOn.ValueRFC3339Time()
	diags.Append(errs...)

	params = firewall.LockdownListParams{}

	if !m.CreatedOn.IsNull() {
		params.CreatedOn = cloudflare.F(mCreatedOn)
	}
	if !m.Description.IsNull() {
		params.Description = cloudflare.F(m.Description.ValueString())
	}
	if !m.DescriptionSearch.IsNull() {
		params.DescriptionSearch = cloudflare.F(m.DescriptionSearch.ValueString())
	}
	if !m.IP.IsNull() {
		params.IP = cloudflare.F(m.IP.ValueString())
	}
	if !m.IPRangeSearch.IsNull() {
		params.IPRangeSearch = cloudflare.F(m.IPRangeSearch.ValueString())
	}
	if !m.IPSearch.IsNull() {
		params.IPSearch = cloudflare.F(m.IPSearch.ValueString())
	}
	if !m.ModifiedOn.IsNull() {
		params.ModifiedOn = cloudflare.F(mModifiedOn)
	}
	if !m.Priority.IsNull() {
		params.Priority = cloudflare.F(m.Priority.ValueFloat64())
	}
	if !m.URISearch.IsNull() {
		params.URISearch = cloudflare.F(m.URISearch.ValueString())
	}

	return
}

type ZoneLockdownsResultDataSourceModel struct {
	ID             types.String                                                         `tfsdk:"id" json:"id,computed"`
	Configurations customfield.NestedObject[ZoneLockdownsConfigurationsDataSourceModel] `tfsdk:"configurations" json:"configurations,computed"`
	CreatedOn      timetypes.RFC3339                                                    `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description    types.String                                                         `tfsdk:"description" json:"description,computed"`
	ModifiedOn     timetypes.RFC3339                                                    `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Paused         types.Bool                                                           `tfsdk:"paused" json:"paused,computed"`
	URLs           customfield.List[types.String]                                       `tfsdk:"urls" json:"urls,computed"`
}

type ZoneLockdownsConfigurationsDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target,computed"`
	Value  types.String `tfsdk:"value" json:"value,computed"`
}
