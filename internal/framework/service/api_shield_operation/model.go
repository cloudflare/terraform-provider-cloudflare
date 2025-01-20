package apishieldoperation

import "github.com/hashicorp/terraform-plugin-framework/types"

type APIShieldOperationModel struct {
	ID       types.String  `tfsdk:"id"`
	ZoneID   types.String  `tfsdk:"zone_id"`
	Method   types.String  `tfsdk:"method"`
	Host     types.String  `tfsdk:"host"`
	Endpoint EndpointValue `tfsdk:"endpoint"`
}
