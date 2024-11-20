// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation_discovery

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationDiscoveryResultEnvelope struct {
	Result APIShieldOperationDiscoveryModel `json:"result"`
}

type APIShieldOperationDiscoveryModel struct {
	ID          types.String `tfsdk:"id" json:"-,computed"`
	OperationID types.String `tfsdk:"operation_id" path:"operation_id,required"`
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id,required"`
	State       types.String `tfsdk:"state" json:"state,optional"`
}

func (m APIShieldOperationDiscoveryModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m APIShieldOperationDiscoveryModel) MarshalJSONForUpdate(state APIShieldOperationDiscoveryModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
