// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_smart_routing

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ArgoSmartRoutingResultEnvelope struct {
	Result ArgoSmartRoutingModel `json:"result"`
}

type ArgoSmartRoutingModel struct {
	ID     types.String `tfsdk:"id" json:"-,computed"`
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Value  types.String `tfsdk:"value" json:"value,required,no_refresh"`
}

func (m ArgoSmartRoutingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ArgoSmartRoutingModel) MarshalJSONForUpdate(state ArgoSmartRoutingModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
