// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemasResultListDataSourceEnvelope struct {
	Result *[]*APIShieldSchemasItemsDataSourceModel `json:"result,computed"`
}

type APIShieldSchemasDataSourceModel struct {
	ZoneID            types.String                             `tfsdk:"zone_id" path:"zone_id"`
	OmitSource        types.Bool                               `tfsdk:"omit_source" query:"omit_source"`
	Page              types.String                             `tfsdk:"page" query:"page"`
	PerPage           types.String                             `tfsdk:"per_page" query:"per_page"`
	ValidationEnabled types.Bool                               `tfsdk:"validation_enabled" query:"validation_enabled"`
	MaxItems          types.Int64                              `tfsdk:"max_items"`
	Items             *[]*APIShieldSchemasItemsDataSourceModel `tfsdk:"items"`
}

type APIShieldSchemasItemsDataSourceModel struct {
	CreatedAt         types.String `tfsdk:"created_at" json:"created_at,computed"`
	Kind              types.String `tfsdk:"kind" json:"kind,computed"`
	Name              types.String `tfsdk:"name" json:"name,computed"`
	SchemaID          types.String `tfsdk:"schema_id" json:"schema_id,computed"`
	Source            types.String `tfsdk:"source" json:"source"`
	ValidationEnabled types.Bool   `tfsdk:"validation_enabled" json:"validation_enabled"`
}
