// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_discovery_operation

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldDiscoveryOperationResultEnvelope struct {
	Result APIShieldDiscoveryOperationModel `json:"result"`
}

type APIShieldDiscoveryOperationModel struct {
	ID          types.String `tfsdk:"id" json:"-,computed"`
	OperationID types.String `tfsdk:"operation_id" path:"operation_id,required"`
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id,required"`
	State       types.String `tfsdk:"state" json:"state,optional"`
}

func (m APIShieldDiscoveryOperationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m APIShieldDiscoveryOperationModel) MarshalJSONForUpdate(state APIShieldDiscoveryOperationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
