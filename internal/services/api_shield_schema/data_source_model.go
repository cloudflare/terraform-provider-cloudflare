// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/api_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemaResultDataSourceEnvelope struct {
	Result APIShieldSchemaDataSourceModel `json:"result,computed"`
}

type APIShieldSchemaResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[APIShieldSchemaDataSourceModel] `json:"result,computed"`
}

type APIShieldSchemaDataSourceModel struct {
	ZoneID            types.String                             `tfsdk:"zone_id" path:"zone_id,optional"`
	SchemaID          types.String                             `tfsdk:"schema_id" path:"schema_id,computed_optional"`
	OmitSource        types.Bool                               `tfsdk:"omit_source" query:"omit_source,optional"`
	CreatedAt         timetypes.RFC3339                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Kind              types.String                             `tfsdk:"kind" json:"kind,computed"`
	Name              types.String                             `tfsdk:"name" json:"name,computed"`
	Source            types.String                             `tfsdk:"source" json:"source,computed"`
	ValidationEnabled types.Bool                               `tfsdk:"validation_enabled" json:"validation_enabled,computed"`
	Filter            *APIShieldSchemaFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *APIShieldSchemaDataSourceModel) toReadParams(_ context.Context) (params api_gateway.UserSchemaGetParams, diags diag.Diagnostics) {
	params = api_gateway.UserSchemaGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *APIShieldSchemaDataSourceModel) toListParams(_ context.Context) (params api_gateway.UserSchemaListParams, diags diag.Diagnostics) {
	params = api_gateway.UserSchemaListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	if !m.Filter.OmitSource.IsNull() {
		params.OmitSource = cloudflare.F(m.Filter.OmitSource.ValueBool())
	}
	if !m.Filter.ValidationEnabled.IsNull() {
		params.ValidationEnabled = cloudflare.F(m.Filter.ValidationEnabled.ValueBool())
	}

	return
}

type APIShieldSchemaFindOneByDataSourceModel struct {
	ZoneID            types.String `tfsdk:"zone_id" path:"zone_id,required"`
	OmitSource        types.Bool   `tfsdk:"omit_source" query:"omit_source,computed_optional"`
	ValidationEnabled types.Bool   `tfsdk:"validation_enabled" query:"validation_enabled,optional"`
}
