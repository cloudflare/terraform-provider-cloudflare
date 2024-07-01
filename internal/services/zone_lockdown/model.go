// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneLockdownResultEnvelope struct {
	Result ZoneLockdownModel `json:"result,computed"`
}

type ZoneLockdownResultDataSourceEnvelope struct {
	Result ZoneLockdownDataSourceModel `json:"result,computed"`
}

type ZoneLockdownsResultDataSourceEnvelope struct {
	Result ZoneLockdownsDataSourceModel `json:"result,computed"`
}

type ZoneLockdownModel struct {
	ID             types.String    `tfsdk:"id" json:"id,computed"`
	ZoneIdentifier types.String    `tfsdk:"zone_identifier" path:"zone_identifier"`
	CreatedOn      types.String    `tfsdk:"created_on" json:"created_on,computed"`
	Description    types.String    `tfsdk:"description" json:"description,computed"`
	ModifiedOn     types.String    `tfsdk:"modified_on" json:"modified_on,computed"`
	Paused         types.Bool      `tfsdk:"paused" json:"paused,computed"`
	URLs           *[]types.String `tfsdk:"urls" json:"urls,computed"`
}

type ZoneLockdownConfigurationsModel struct {
	Target types.String `tfsdk:"target" json:"target"`
	Value  types.String `tfsdk:"value" json:"value"`
}

type ZoneLockdownDataSourceModel struct {
}

type ZoneLockdownsDataSourceModel struct {
}
