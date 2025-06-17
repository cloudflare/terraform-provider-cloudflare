// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_schemas

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/schema_validation"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SchemaValidationSchemasResultDataSourceEnvelope struct {
	Result SchemaValidationSchemasDataSourceModel `json:"result,computed"`
}

type SchemaValidationSchemasDataSourceModel struct {
	ID                types.String                                     `tfsdk:"id" path:"schema_id,computed"`
	SchemaID          types.String                                     `tfsdk:"schema_id" path:"schema_id,computed_optional"`
	ZoneID            types.String                                     `tfsdk:"zone_id" path:"zone_id,required"`
	OmitSource        types.Bool                                       `tfsdk:"omit_source" query:"omit_source,computed_optional"`
	CreatedAt         timetypes.RFC3339                                `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Kind              types.String                                     `tfsdk:"kind" json:"kind,computed"`
	Name              types.String                                     `tfsdk:"name" json:"name,computed"`
	Source            types.String                                     `tfsdk:"source" json:"source,computed"`
	ValidationEnabled types.Bool                                       `tfsdk:"validation_enabled" json:"validation_enabled,computed"`
	Filter            *SchemaValidationSchemasFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *SchemaValidationSchemasDataSourceModel) toReadParams(_ context.Context) (params schema_validation.SchemaGetParams, diags diag.Diagnostics) {
	params = schema_validation.SchemaGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *SchemaValidationSchemasDataSourceModel) toListParams(_ context.Context) (params schema_validation.SchemaListParams, diags diag.Diagnostics) {
	params = schema_validation.SchemaListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Filter.OmitSource.IsNull() {
		params.OmitSource = cloudflare.F(m.Filter.OmitSource.ValueBool())
	}
	if !m.Filter.ValidationEnabled.IsNull() {
		params.ValidationEnabled = cloudflare.F(m.Filter.ValidationEnabled.ValueBool())
	}

	return
}

type SchemaValidationSchemasFindOneByDataSourceModel struct {
	OmitSource        types.Bool `tfsdk:"omit_source" query:"omit_source,computed_optional"`
	ValidationEnabled types.Bool `tfsdk:"validation_enabled" query:"validation_enabled,optional"`
}
