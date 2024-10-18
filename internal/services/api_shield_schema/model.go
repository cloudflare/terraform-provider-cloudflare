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
	ZoneID            types.String                                                `tfsdk:"zone_id" path:"zone_id,required"`
	SchemaID          types.String                                                `tfsdk:"schema_id" path:"schema_id,optional"`
	File              types.String                                                `tfsdk:"file" json:"file,required"`
	Kind              types.String                                                `tfsdk:"kind" json:"kind,required"`
	Name              types.String                                                `tfsdk:"name" json:"name,optional"`
	ValidationEnabled types.String                                                `tfsdk:"validation_enabled" json:"validation_enabled,optional"`
	CreatedAt         timetypes.RFC3339                                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Source            types.String                                                `tfsdk:"source" json:"source,computed"`
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

type APIShieldSchemaSchemaModel struct {
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Kind              types.String      `tfsdk:"kind" json:"kind,computed"`
	Name              types.String      `tfsdk:"name" json:"name,computed"`
	SchemaID          types.String      `tfsdk:"schema_id" json:"schema_id,computed"`
	Source            types.String      `tfsdk:"source" json:"source,computed"`
	ValidationEnabled types.Bool        `tfsdk:"validation_enabled" json:"validation_enabled,computed"`
}

type APIShieldSchemaUploadDetailsModel struct {
	Warnings customfield.NestedObjectList[APIShieldSchemaUploadDetailsWarningsModel] `tfsdk:"warnings" json:"warnings,computed"`
}

type APIShieldSchemaUploadDetailsWarningsModel struct {
	Code      types.Int64                    `tfsdk:"code" json:"code,computed"`
	Locations customfield.List[types.String] `tfsdk:"locations" json:"locations,computed"`
	Message   types.String                   `tfsdk:"message" json:"message,computed"`
}
