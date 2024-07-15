// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
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
	CreatedAt         timetypes.RFC3339                        `tfsdk:"created_at" json:"created_at,computed"`
	Kind              types.String                             `tfsdk:"kind" json:"kind,computed"`
	Name              types.String                             `tfsdk:"name" json:"name,computed"`
	Source            types.String                             `tfsdk:"source" json:"source"`
	ValidationEnabled types.Bool                               `tfsdk:"validation_enabled" json:"validation_enabled"`
	FindOneBy         *APIShieldSchemaFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type APIShieldSchemaFindOneByDataSourceModel struct {
	ZoneID            types.String         `tfsdk:"zone_id" path:"zone_id"`
	OmitSource        types.Bool           `tfsdk:"omit_source" query:"omit_source"`
	Page              jsontypes.Normalized `tfsdk:"page" query:"page"`
	PerPage           jsontypes.Normalized `tfsdk:"per_page" query:"per_page"`
	ValidationEnabled types.Bool           `tfsdk:"validation_enabled" query:"validation_enabled"`
}
