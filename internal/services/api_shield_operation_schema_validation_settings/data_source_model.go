// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation_schema_validation_settings

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/api_gateway"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationSchemaValidationSettingsDataSourceModel struct {
	OperationID      types.String `tfsdk:"operation_id" path:"operation_id"`
	ZoneID           types.String `tfsdk:"zone_id" path:"zone_id"`
	MitigationAction types.String `tfsdk:"mitigation_action" json:"mitigation_action"`
}

func (m *APIShieldOperationSchemaValidationSettingsDataSourceModel) toReadParams() (params api_gateway.OperationSchemaValidationGetParams, diags diag.Diagnostics) {
	params = api_gateway.OperationSchemaValidationGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
