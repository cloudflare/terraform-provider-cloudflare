// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneLockdownResultEnvelope struct {
	Result ZoneLockdownModel `json:"result,computed"`
}

type ZoneLockdownModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
}
