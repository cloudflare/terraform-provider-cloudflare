// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_stream

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PipelineStreamResultEnvelope struct {
	Result PipelineStreamModel `json:"result"`
}

type PipelineStreamModel struct {
	ID            types.String                                               `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                               `tfsdk:"account_id" path:"account_id,required"`
	Name          types.String                                               `tfsdk:"name" json:"name,required"`
	Format        *PipelineStreamFormatModel                                 `tfsdk:"format" json:"format,optional"`
	Schema        *PipelineStreamSchemaModel                                 `tfsdk:"schema" json:"schema,optional"`
	HTTP          customfield.NestedObject[PipelineStreamHTTPModel]          `tfsdk:"http" json:"http,computed_optional"`
	WorkerBinding customfield.NestedObject[PipelineStreamWorkerBindingModel] `tfsdk:"worker_binding" json:"worker_binding,computed_optional"`
	CreatedAt     timetypes.RFC3339                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Endpoint      types.String                                               `tfsdk:"endpoint" json:"endpoint,computed"`
	ModifiedAt    timetypes.RFC3339                                          `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Version       types.Int64                                                `tfsdk:"version" json:"version,computed"`
}

func (m PipelineStreamModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PipelineStreamModel) MarshalJSONForUpdate(state PipelineStreamModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type PipelineStreamFormatModel struct {
	Type            types.String `tfsdk:"type" json:"type,required"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,optional"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,optional"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,optional"`
	Compression     types.String `tfsdk:"compression" json:"compression,optional"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,optional"`
}

type PipelineStreamSchemaModel struct {
	Fields   *[]*PipelineStreamSchemaFieldsModel `tfsdk:"fields" json:"fields,optional"`
	Format   *PipelineStreamSchemaFormatModel    `tfsdk:"format" json:"format,optional"`
	Inferred types.Bool                          `tfsdk:"inferred" json:"inferred,optional"`
}

type PipelineStreamSchemaFieldsModel struct {
	Type        types.String `tfsdk:"type" json:"type,required"`
	MetadataKey types.String `tfsdk:"metadata_key" json:"metadata_key,optional"`
	Name        types.String `tfsdk:"name" json:"name,optional"`
	Required    types.Bool   `tfsdk:"required" json:"required,optional"`
	SqlName     types.String `tfsdk:"sql_name" json:"sql_name,optional"`
	Unit        types.String `tfsdk:"unit" json:"unit,optional"`
}

type PipelineStreamSchemaFormatModel struct {
	Type            types.String `tfsdk:"type" json:"type,required"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,optional"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,optional"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,optional"`
	Compression     types.String `tfsdk:"compression" json:"compression,optional"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,optional"`
}

type PipelineStreamHTTPModel struct {
	Authentication types.Bool                   `tfsdk:"authentication" json:"authentication,required"`
	Enabled        types.Bool                   `tfsdk:"enabled" json:"enabled,required"`
	CORS           *PipelineStreamHTTPCORSModel `tfsdk:"cors" json:"cors,optional"`
}

type PipelineStreamHTTPCORSModel struct {
	Origins *[]types.String `tfsdk:"origins" json:"origins,optional"`
}

type PipelineStreamWorkerBindingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,required"`
}
