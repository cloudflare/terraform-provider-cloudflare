// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_hold

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneHoldResultEnvelope struct {
	Result ZoneHoldModel `json:"result,computed"`
}

type ZoneHoldModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
