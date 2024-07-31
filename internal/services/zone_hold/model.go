// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_hold

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneHoldResultEnvelope struct {
	Result ZoneHoldModel `json:"result,computed"`
}

type ZoneHoldModel struct {
	ID                types.String `tfsdk:"id" json:"-,computed"`
	ZoneID            types.String `tfsdk:"zone_id" path:"zone_id"`
	Hold              types.Bool   `tfsdk:"hold" json:"hold,computed"`
	HoldAfter         types.String `tfsdk:"hold_after" json:"hold_after,computed"`
	IncludeSubdomains types.String `tfsdk:"include_subdomains" json:"include_subdomains,computed"`
}
