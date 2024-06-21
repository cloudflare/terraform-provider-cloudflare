// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemaResultEnvelope struct {
	Result APIShieldSchemaModel `json:"result,computed"`
}

type APIShieldSchemaModel struct {
	ZoneID            types.String `tfsdk:"zone_id" path:"zone_id"`
	SchemaID          types.String `tfsdk:"schema_id" path:"schema_id"`
	File              types.String `tfsdk:"file" json:"file"`
	Kind              types.String `tfsdk:"kind" json:"kind"`
	Name              types.String `tfsdk:"name" json:"name"`
	ValidationEnabled types.String `tfsdk:"validation_enabled" json:"validation_enabled"`
}
