// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemasResultListDataSourceEnvelope struct {
	Result *[]*APIShieldSchemasResultDataSourceModel `json:"result,computed"`
}

type APIShieldSchemasDataSourceModel struct {
	ZoneID            types.String                              `tfsdk:"zone_id" path:"zone_id"`
	ValidationEnabled types.Bool                                `tfsdk:"validation_enabled" query:"validation_enabled"`
	OmitSource        types.Bool                                `tfsdk:"omit_source" query:"omit_source"`
	MaxItems          types.Int64                               `tfsdk:"max_items"`
	Result            *[]*APIShieldSchemasResultDataSourceModel `tfsdk:"result"`
}

type APIShieldSchemasResultDataSourceModel struct {
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Kind              types.String      `tfsdk:"kind" json:"kind,computed"`
	Name              types.String      `tfsdk:"name" json:"name,computed"`
	SchemaID          types.String      `tfsdk:"schema_id" json:"schema_id,computed"`
	Source            types.String      `tfsdk:"source" json:"source"`
	ValidationEnabled types.Bool        `tfsdk:"validation_enabled" json:"validation_enabled"`
}
