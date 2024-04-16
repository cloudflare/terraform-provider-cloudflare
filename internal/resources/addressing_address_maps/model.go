// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing_address_maps

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddressingAddressMapsResultEnvelope struct {
	Result AddressingAddressMapsModel `json:"result,computed"`
}

type AddressingAddressMapsModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	Description types.String `tfsdk:"description" json:"description"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled"`
}
