// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation_schema_validation_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/api_gateway"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationSchemaValidationSettingsDataSourceModel struct {
	OperationID      types.String `tfsdk:"operation_id" path:"operation_id,required"`
	ZoneID           types.String `tfsdk:"zone_id" path:"zone_id,required"`
	MitigationAction types.String `tfsdk:"mitigation_action" json:"mitigation_action,computed"`
}

func (m *APIShieldOperationSchemaValidationSettingsDataSourceModel) toReadParams(_ context.Context) (params api_gateway.OperationSchemaValidationGetParams, diags diag.Diagnostics) {
	params = api_gateway.OperationSchemaValidationGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
