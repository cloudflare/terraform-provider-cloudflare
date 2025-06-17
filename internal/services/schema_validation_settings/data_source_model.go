// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/schema_validation"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SchemaValidationSettingsResultDataSourceEnvelope struct {
	Result SchemaValidationSettingsDataSourceModel `json:"result,computed"`
}

type SchemaValidationSettingsDataSourceModel struct {
	ZoneID                             types.String `tfsdk:"zone_id" path:"zone_id,required"`
	ValidationDefaultMitigationAction  types.String `tfsdk:"validation_default_mitigation_action" json:"validation_default_mitigation_action,computed"`
	ValidationOverrideMitigationAction types.String `tfsdk:"validation_override_mitigation_action" json:"validation_override_mitigation_action,computed"`
}

func (m *SchemaValidationSettingsDataSourceModel) toReadParams(_ context.Context) (params schema_validation.SettingGetParams, diags diag.Diagnostics) {
	params = schema_validation.SettingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
