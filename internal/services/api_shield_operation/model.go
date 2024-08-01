// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationResultEnvelope struct {
	Result APIShieldOperationModel `json:"result,computed"`
}

type APIShieldOperationModel struct {
	ID          types.String `tfsdk:"id" json:"-,computed"`
	OperationID types.String `tfsdk:"operation_id" path:"operation_id"`
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id"`
	State       types.String `tfsdk:"state" json:"state"`
}
