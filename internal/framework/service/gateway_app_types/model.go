package gateway_app_types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GatewayAppTypesModel struct {
	AccountID types.String          `tfsdk:"account_id"`
	AppTypes  []GatewayAppTypeModel `tfsdk:"app_types"`
}

type GatewayAppTypeModel struct {
	ID                types.Int64  `tfsdk:"id"`
	ApplicationTypeID types.Int64  `tfsdk:"application_type_id"`
	Name              types.String `tfsdk:"name"`
	Description       types.String `tfsdk:"description"`
}
