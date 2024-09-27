// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneLockdownResultEnvelope struct {
	Result ZoneLockdownModel `json:"result"`
}

type ZoneLockdownModel struct {
	ID             types.String                     `tfsdk:"id" json:"id,computed"`
	ZoneIdentifier types.String                     `tfsdk:"zone_identifier" path:"zone_identifier,required"`
	URLs           *[]types.String                  `tfsdk:"urls" json:"urls,required"`
	Configurations *ZoneLockdownConfigurationsModel `tfsdk:"configurations" json:"configurations,required"`
	CreatedOn      timetypes.RFC3339                `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description    types.String                     `tfsdk:"description" json:"description,computed"`
	ModifiedOn     timetypes.RFC3339                `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Paused         types.Bool                       `tfsdk:"paused" json:"paused,computed"`
}

type ZoneLockdownConfigurationsModel struct {
	Target types.String `tfsdk:"target" json:"target,optional"`
	Value  types.String `tfsdk:"value" json:"value,optional"`
}
