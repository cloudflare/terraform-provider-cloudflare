// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_operation_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/schema_validation"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SchemaValidationOperationSettingsResultDataSourceEnvelope struct {
	Result SchemaValidationOperationSettingsDataSourceModel `json:"result,computed"`
}

type SchemaValidationOperationSettingsDataSourceModel struct {
	OperationID      types.String `tfsdk:"operation_id" path:"operation_id,required"`
	ZoneID           types.String `tfsdk:"zone_id" path:"zone_id,required"`
	MitigationAction types.String `tfsdk:"mitigation_action" json:"mitigation_action,computed"`
}

func (m *SchemaValidationOperationSettingsDataSourceModel) toReadParams(_ context.Context) (params schema_validation.SettingOperationGetParams, diags diag.Diagnostics) {
	params = schema_validation.SettingOperationGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
