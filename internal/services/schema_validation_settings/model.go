// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SchemaValidationSettingsResultEnvelope struct {
	Result SchemaValidationSettingsModel `json:"result"`
}

type SchemaValidationSettingsModel struct {
	ZoneID                             types.String `tfsdk:"zone_id" path:"zone_id,required"`
	ValidationDefaultMitigationAction  types.String `tfsdk:"validation_default_mitigation_action" json:"validation_default_mitigation_action,required"`
	ValidationOverrideMitigationAction types.String `tfsdk:"validation_override_mitigation_action" json:"validation_override_mitigation_action,optional"`
}

func (m SchemaValidationSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SchemaValidationSettingsModel) MarshalJSONForUpdate(state SchemaValidationSettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
