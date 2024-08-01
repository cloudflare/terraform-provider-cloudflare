// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemaResultEnvelope struct {
	Result APIShieldSchemaModel `json:"result,computed"`
}

type APIShieldSchemaModel struct {
	ZoneID            types.String                                                `tfsdk:"zone_id" path:"zone_id"`
	SchemaID          types.String                                                `tfsdk:"schema_id" path:"schema_id"`
	File              types.String                                                `tfsdk:"file" json:"file"`
	Kind              types.String                                                `tfsdk:"kind" json:"kind"`
	Name              types.String                                                `tfsdk:"name" json:"name"`
	ValidationEnabled types.String                                                `tfsdk:"validation_enabled" json:"validation_enabled"`
	CreatedAt         timetypes.RFC3339                                           `tfsdk:"created_at" json:"created_at,computed"`
	Source            types.String                                                `tfsdk:"source" json:"source,computed"`
	Schema            customfield.NestedObject[APIShieldSchemaSchemaModel]        `tfsdk:"schema" json:"schema,computed"`
	UploadDetails     customfield.NestedObject[APIShieldSchemaUploadDetailsModel] `tfsdk:"upload_details" json:"upload_details,computed"`
}

type APIShieldSchemaSchemaModel struct {
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Kind              types.String      `tfsdk:"kind" json:"kind"`
	Name              types.String      `tfsdk:"name" json:"name"`
	SchemaID          types.String      `tfsdk:"schema_id" json:"schema_id"`
	Source            types.String      `tfsdk:"source" json:"source"`
	ValidationEnabled types.Bool        `tfsdk:"validation_enabled" json:"validation_enabled"`
}

type APIShieldSchemaUploadDetailsModel struct {
	Warnings *[]*APIShieldSchemaUploadDetailsWarningsModel `tfsdk:"warnings" json:"warnings"`
}

type APIShieldSchemaUploadDetailsWarningsModel struct {
	Code      types.Int64     `tfsdk:"code" json:"code"`
	Locations *[]types.String `tfsdk:"locations" json:"locations"`
	Message   types.String    `tfsdk:"message" json:"message"`
}
