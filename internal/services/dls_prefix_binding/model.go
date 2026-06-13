// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dls_prefix_binding

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DLSPrefixBindingResultEnvelope struct {
	Result DLSPrefixBindingModel `json:"result"`
}

type DLSPrefixBindingModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	CIDR      types.String `tfsdk:"cidr" json:"cidr,required"`
	PrefixID  types.String `tfsdk:"prefix_id" json:"prefix_id,required"`
	RegionKey types.String `tfsdk:"region_key" json:"region_key,required"`
}

func (m DLSPrefixBindingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DLSPrefixBindingModel) MarshalJSONForUpdate(state DLSPrefixBindingModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
