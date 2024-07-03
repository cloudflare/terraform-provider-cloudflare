// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemaResultDataSourceEnvelope struct {
	Result APIShieldSchemaDataSourceModel `json:"result,computed"`
}

type APIShieldSchemaResultListDataSourceEnvelope struct {
	Result *[]*APIShieldSchemaDataSourceModel `json:"result,computed"`
}

type APIShieldSchemaDataSourceModel struct {
	ZoneID            types.String                             `tfsdk:"zone_id" path:"zone_id"`
	SchemaID          types.String                             `tfsdk:"schema_id" path:"schema_id"`
	OmitSource        types.Bool                               `tfsdk:"omit_source" query:"omit_source"`
	CreatedAt         types.String                             `tfsdk:"created_at" json:"created_at"`
	Kind              types.String                             `tfsdk:"kind" json:"kind"`
	Name              types.String                             `tfsdk:"name" json:"name"`
	Source            types.String                             `tfsdk:"source" json:"source"`
	ValidationEnabled types.Bool                               `tfsdk:"validation_enabled" json:"validation_enabled"`
	FindOneBy         *APIShieldSchemaFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type APIShieldSchemaFindOneByDataSourceModel struct {
	ZoneID            types.String `tfsdk:"zone_id" path:"zone_id"`
	OmitSource        types.Bool   `tfsdk:"omit_source" query:"omit_source"`
	Page              types.String `tfsdk:"page" query:"page"`
	PerPage           types.String `tfsdk:"per_page" query:"per_page"`
	ValidationEnabled types.Bool   `tfsdk:"validation_enabled" query:"validation_enabled"`
}
