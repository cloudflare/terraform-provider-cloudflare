// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation_schema_validation_settings

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationSchemaValidationSettingsModel struct {
	ID               types.String `tfsdk:"id" json:"-,computed"`
	OperationID      types.String `tfsdk:"operation_id" path:"operation_id,required"`
	ZoneID           types.String `tfsdk:"zone_id" path:"zone_id,required"`
	MitigationAction types.String `tfsdk:"mitigation_action" json:"mitigation_action,optional"`
}
