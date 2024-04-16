// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZonesResultEnvelope struct {
	Result ZonesModel `json:"result,computed"`
}

type ZonesModel struct {
	ID      types.String       `tfsdk:"id" json:"id,computed"`
	Account *ZonesAccountModel `tfsdk:"account" json:"account"`
	Name    types.String       `tfsdk:"name" json:"name"`
	Type    types.String       `tfsdk:"type" json:"type"`
}

type ZonesAccountModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}
