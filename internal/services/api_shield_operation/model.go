// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationResultEnvelope struct {
	Result APIShieldOperationModel `json:"result,computed"`
}

type APIShieldOperationsResultDataSourceEnvelope struct {
	Result APIShieldOperationsDataSourceModel `json:"result,computed"`
}

type APIShieldOperationModel struct {
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id"`
	OperationID types.String `tfsdk:"operation_id" path:"operation_id"`
	State       types.String `tfsdk:"state" json:"state"`
}

type APIShieldOperationsDataSourceModel struct {
}
