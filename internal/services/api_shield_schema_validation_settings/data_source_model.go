// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema_validation_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/api_gateway"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemaValidationSettingsDataSourceModel struct {
	ZoneID                             types.String `tfsdk:"zone_id" path:"zone_id,required"`
	ValidationDefaultMitigationAction  types.String `tfsdk:"validation_default_mitigation_action" json:"validation_default_mitigation_action,computed"`
	ValidationOverrideMitigationAction types.String `tfsdk:"validation_override_mitigation_action" json:"validation_override_mitigation_action,computed"`
}

func (m *APIShieldSchemaValidationSettingsDataSourceModel) toReadParams(_ context.Context) (params api_gateway.SettingSchemaValidationGetParams, diags diag.Diagnostics) {
	params = api_gateway.SettingSchemaValidationGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
