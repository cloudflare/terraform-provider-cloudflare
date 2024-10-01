// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationResultEnvelope struct {
	Result APIShieldOperationModel `json:"result"`
}

type APIShieldOperationModel struct {
	ID          types.String `tfsdk:"id" json:"-,computed"`
	OperationID types.String `tfsdk:"operation_id" path:"operation_id,required"`
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id,required"`
	State       types.String `tfsdk:"state" json:"state,optional"`
}

func (m APIShieldOperationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m APIShieldOperationModel) MarshalJSONForUpdate(state APIShieldOperationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
