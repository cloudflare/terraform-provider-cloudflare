// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneLockdownResultEnvelope struct {
	Result ZoneLockdownModel `json:"result"`
}

type ZoneLockdownModel struct {
	ID             types.String                        `tfsdk:"id" json:"id,computed"`
	ZoneID         types.String                        `tfsdk:"zone_id" path:"zone_id,required"`
	Description    types.String                        `tfsdk:"description" json:"description,optional"`
	Paused         types.Bool                          `tfsdk:"paused" json:"paused,optional"`
	Priority       types.Float64                       `tfsdk:"priority" json:"priority,optional,no_refresh"`
	URLs           *[]types.String                     `tfsdk:"urls" json:"urls,required"`
	Configurations *[]*ZoneLockdownConfigurationsModel `tfsdk:"configurations" json:"configurations,required"`
	CreatedOn      timetypes.RFC3339                   `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn     timetypes.RFC3339                   `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m ZoneLockdownModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZoneLockdownModel) MarshalJSONForUpdate(state ZoneLockdownModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZoneLockdownConfigurationsModel struct {
	Target types.String `tfsdk:"target" json:"target,optional"`
	Value  types.String `tfsdk:"value" json:"value,optional"`
}
