// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_operation_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SchemaValidationOperationSettingsResultEnvelope struct {
	Result SchemaValidationOperationSettingsModel `json:"result"`
}

type SchemaValidationOperationSettingsModel struct {
	OperationID      types.String `tfsdk:"operation_id" path:"operation_id,required"`
	ZoneID           types.String `tfsdk:"zone_id" path:"zone_id,required"`
	MitigationAction types.String `tfsdk:"mitigation_action" json:"mitigation_action,required"`
}

func (m SchemaValidationOperationSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SchemaValidationOperationSettingsModel) MarshalJSONForUpdate(state SchemaValidationOperationSettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
