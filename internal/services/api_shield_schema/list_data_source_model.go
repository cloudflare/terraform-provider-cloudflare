// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/api_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemasResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[APIShieldSchemasResultDataSourceModel] `json:"result,computed"`
}

type APIShieldSchemasDataSourceModel struct {
	ZoneID            types.String                                                        `tfsdk:"zone_id" path:"zone_id,required"`
	ValidationEnabled types.Bool                                                          `tfsdk:"validation_enabled" query:"validation_enabled,optional"`
	OmitSource        types.Bool                                                          `tfsdk:"omit_source" query:"omit_source,computed_optional"`
	MaxItems          types.Int64                                                         `tfsdk:"max_items"`
	Result            customfield.NestedObjectList[APIShieldSchemasResultDataSourceModel] `tfsdk:"result"`
}

func (m *APIShieldSchemasDataSourceModel) toListParams(_ context.Context) (params api_gateway.UserSchemaListParams, diags diag.Diagnostics) {
	params = api_gateway.UserSchemaListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.OmitSource.IsNull() {
		params.OmitSource = cloudflare.F(m.OmitSource.ValueBool())
	}
	if !m.ValidationEnabled.IsNull() {
		params.ValidationEnabled = cloudflare.F(m.ValidationEnabled.ValueBool())
	}

	return
}

type APIShieldSchemasResultDataSourceModel struct {
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Kind              types.String      `tfsdk:"kind" json:"kind,computed"`
	Name              types.String      `tfsdk:"name" json:"name,computed"`
	SchemaID          types.String      `tfsdk:"schema_id" json:"schema_id,computed"`
	Source            types.String      `tfsdk:"source" json:"source,computed"`
	ValidationEnabled types.Bool        `tfsdk:"validation_enabled" json:"validation_enabled,computed"`
}
