// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneResultEnvelope struct {
	Result ZoneModel `json:"result,computed"`
}

type ZoneModel struct {
	ID      types.String      `tfsdk:"id" json:"id,computed"`
	Account *ZoneAccountModel `tfsdk:"account" json:"account"`
	Name    types.String      `tfsdk:"name" json:"name"`
	Type    types.String      `tfsdk:"type" json:"type"`
}

type ZoneAccountModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}
