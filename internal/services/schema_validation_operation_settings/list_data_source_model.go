// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_operation_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/schema_validation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SchemaValidationOperationSettingsListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SchemaValidationOperationSettingsListResultDataSourceModel] `json:"result,computed"`
}

type SchemaValidationOperationSettingsListDataSourceModel struct {
	ZoneID   types.String                                                                             `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                                                              `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[SchemaValidationOperationSettingsListResultDataSourceModel] `tfsdk:"result"`
}

func (m *SchemaValidationOperationSettingsListDataSourceModel) toListParams(_ context.Context) (params schema_validation.SettingOperationListParams, diags diag.Diagnostics) {
	params = schema_validation.SettingOperationListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type SchemaValidationOperationSettingsListResultDataSourceModel struct {
	MitigationAction types.String `tfsdk:"mitigation_action" json:"mitigation_action,computed"`
	OperationID      types.String `tfsdk:"operation_id" json:"operation_id,computed"`
}
