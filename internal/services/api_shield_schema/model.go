// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"bytes"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemaResultEnvelope struct {
	Result APIShieldSchemaModel `json:"result"`
}

type APIShieldSchemaModel struct {
	ZoneID            types.String                                                `tfsdk:"zone_id" path:"zone_id"`
	SchemaID          types.String                                                `tfsdk:"schema_id" path:"schema_id"`
	File              types.String                                                `tfsdk:"file" json:"file"`
	Kind              types.String                                                `tfsdk:"kind" json:"kind"`
	Name              types.String                                                `tfsdk:"name" json:"name"`
	ValidationEnabled types.String                                                `tfsdk:"validation_enabled" json:"validation_enabled"`
	CreatedAt         timetypes.RFC3339                                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Source            types.String                                                `tfsdk:"source" json:"source,computed"`
	Success           types.Bool                                                  `tfsdk:"success" json:"success,computed"`
	Errors            customfield.NestedObjectList[APIShieldSchemaErrorsModel]    `tfsdk:"errors" json:"errors,computed"`
	Messages          customfield.NestedObjectList[APIShieldSchemaMessagesModel]  `tfsdk:"messages" json:"messages,computed"`
	Schema            customfield.NestedObject[APIShieldSchemaSchemaModel]        `tfsdk:"schema" json:"schema,computed"`
	UploadDetails     customfield.NestedObject[APIShieldSchemaUploadDetailsModel] `tfsdk:"upload_details" json:"upload_details,computed"`
}

func (r APIShieldSchemaModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type APIShieldSchemaErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type APIShieldSchemaMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type APIShieldSchemaSchemaModel struct {
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Kind              types.String      `tfsdk:"kind" json:"kind"`
	Name              types.String      `tfsdk:"name" json:"name"`
	SchemaID          types.String      `tfsdk:"schema_id" json:"schema_id"`
	Source            types.String      `tfsdk:"source" json:"source,computed_optional"`
	ValidationEnabled types.Bool        `tfsdk:"validation_enabled" json:"validation_enabled,computed_optional"`
}

type APIShieldSchemaUploadDetailsModel struct {
	Warnings customfield.NestedObjectList[APIShieldSchemaUploadDetailsWarningsModel] `tfsdk:"warnings" json:"warnings,computed_optional"`
}

type APIShieldSchemaUploadDetailsWarningsModel struct {
	Code      types.Int64  `tfsdk:"code" json:"code"`
	Locations types.List   `tfsdk:"locations" json:"locations,computed_optional"`
	Message   types.String `tfsdk:"message" json:"message,computed_optional"`
}
