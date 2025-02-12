// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/api_gateway"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemaResultDataSourceEnvelope struct {
	Result APIShieldSchemaDataSourceModel `json:"result,computed"`
}

type APIShieldSchemaDataSourceModel struct {
	ZoneID            types.String `tfsdk:"zone_id" path:"zone_id,required"`
	SchemaID          types.String `tfsdk:"schema_id" path:"schema_id,computed"`
	OmitSource        types.Bool   `tfsdk:"omit_source" query:"omit_source,computed_optional"`
	CreatedAt         types.String `tfsdk:"created_at" json:"created_at,computed"`
	Kind              types.String `tfsdk:"kind" json:"kind,computed"`
	Name              types.String `tfsdk:"name" json:"name,computed"`
	Source            types.String `tfsdk:"source" json:"source,computed"`
	ValidationEnabled types.Bool   `tfsdk:"validation_enabled" json:"validation_enabled,computed"`
}

func (m *APIShieldSchemaDataSourceModel) toReadParams(_ context.Context) (params api_gateway.UserSchemaGetParams, diags diag.Diagnostics) {
	params = api_gateway.UserSchemaGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
