// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_stream

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/pipelines"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PipelineStreamResultDataSourceEnvelope struct {
	Result PipelineStreamDataSourceModel `json:"result,computed"`
}

type PipelineStreamDataSourceModel struct {
	ID            types.String                                                         `tfsdk:"id" path:"stream_id,computed"`
	StreamID      types.String                                                         `tfsdk:"stream_id" path:"stream_id,optional"`
	AccountID     types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt     timetypes.RFC3339                                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Endpoint      types.String                                                         `tfsdk:"endpoint" json:"endpoint,computed"`
	ModifiedAt    timetypes.RFC3339                                                    `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Name          types.String                                                         `tfsdk:"name" json:"name,computed"`
	Version       types.Int64                                                          `tfsdk:"version" json:"version,computed"`
	Format        customfield.NestedObject[PipelineStreamFormatDataSourceModel]        `tfsdk:"format" json:"format,computed"`
	HTTP          customfield.NestedObject[PipelineStreamHTTPDataSourceModel]          `tfsdk:"http" json:"http,computed"`
	Schema        customfield.NestedObject[PipelineStreamSchemaDataSourceModel]        `tfsdk:"schema" json:"schema,computed"`
	WorkerBinding customfield.NestedObject[PipelineStreamWorkerBindingDataSourceModel] `tfsdk:"worker_binding" json:"worker_binding,computed"`
	Filter        *PipelineStreamFindOneByDataSourceModel                              `tfsdk:"filter"`
}

func (m *PipelineStreamDataSourceModel) toReadParams(_ context.Context) (params pipelines.StreamGetParams, diags diag.Diagnostics) {
	params = pipelines.StreamGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *PipelineStreamDataSourceModel) toListParams(_ context.Context) (params pipelines.StreamListParams, diags diag.Diagnostics) {
	params = pipelines.StreamListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.PipelineID.IsNull() {
		params.PipelineID = cloudflare.F(m.Filter.PipelineID.ValueString())
	}

	return
}

type PipelineStreamFormatDataSourceModel struct {
	Type            types.String `tfsdk:"type" json:"type,computed"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,computed"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,computed"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,computed"`
	Compression     types.String `tfsdk:"compression" json:"compression,computed"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,computed"`
}

type PipelineStreamHTTPDataSourceModel struct {
	Authentication types.Bool                                                      `tfsdk:"authentication" json:"authentication,computed"`
	Enabled        types.Bool                                                      `tfsdk:"enabled" json:"enabled,computed"`
	CORS           customfield.NestedObject[PipelineStreamHTTPCORSDataSourceModel] `tfsdk:"cors" json:"cors,computed"`
}

type PipelineStreamHTTPCORSDataSourceModel struct {
	Origins customfield.List[types.String] `tfsdk:"origins" json:"origins,computed"`
}

type PipelineStreamSchemaDataSourceModel struct {
	Fields   customfield.NestedObjectList[PipelineStreamSchemaFieldsDataSourceModel] `tfsdk:"fields" json:"fields,computed"`
	Format   customfield.NestedObject[PipelineStreamSchemaFormatDataSourceModel]     `tfsdk:"format" json:"format,computed"`
	Inferred types.Bool                                                              `tfsdk:"inferred" json:"inferred,computed"`
}

type PipelineStreamSchemaFieldsDataSourceModel struct {
	Type        types.String `tfsdk:"type" json:"type,computed"`
	MetadataKey types.String `tfsdk:"metadata_key" json:"metadata_key,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Required    types.Bool   `tfsdk:"required" json:"required,computed"`
	SqlName     types.String `tfsdk:"sql_name" json:"sql_name,computed"`
	Unit        types.String `tfsdk:"unit" json:"unit,computed"`
}

type PipelineStreamSchemaFormatDataSourceModel struct {
	Type            types.String `tfsdk:"type" json:"type,computed"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,computed"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,computed"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,computed"`
	Compression     types.String `tfsdk:"compression" json:"compression,computed"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,computed"`
}

type PipelineStreamWorkerBindingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type PipelineStreamFindOneByDataSourceModel struct {
	PipelineID types.String `tfsdk:"pipeline_id" query:"pipeline_id,optional"`
}
