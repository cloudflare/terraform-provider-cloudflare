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

type PipelineStreamsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PipelineStreamsResultDataSourceModel] `json:"result,computed"`
}

type PipelineStreamsDataSourceModel struct {
	AccountID  types.String                                                       `tfsdk:"account_id" path:"account_id,required"`
	PipelineID types.String                                                       `tfsdk:"pipeline_id" query:"pipeline_id,optional"`
	MaxItems   types.Int64                                                        `tfsdk:"max_items"`
	Result     customfield.NestedObjectList[PipelineStreamsResultDataSourceModel] `tfsdk:"result"`
}

func (m *PipelineStreamsDataSourceModel) toListParams(_ context.Context) (params pipelines.StreamListParams, diags diag.Diagnostics) {
	params = pipelines.StreamListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.PipelineID.IsNull() {
		params.PipelineID = cloudflare.F(m.PipelineID.ValueString())
	}

	return
}

type PipelineStreamsResultDataSourceModel struct {
	ID            types.String                                                          `tfsdk:"id" json:"id,computed"`
	CreatedAt     timetypes.RFC3339                                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	HTTP          customfield.NestedObject[PipelineStreamsHTTPDataSourceModel]          `tfsdk:"http" json:"http,computed"`
	ModifiedAt    timetypes.RFC3339                                                     `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Name          types.String                                                          `tfsdk:"name" json:"name,computed"`
	Version       types.Int64                                                           `tfsdk:"version" json:"version,computed"`
	WorkerBinding customfield.NestedObject[PipelineStreamsWorkerBindingDataSourceModel] `tfsdk:"worker_binding" json:"worker_binding,computed"`
	Endpoint      types.String                                                          `tfsdk:"endpoint" json:"endpoint,computed"`
	Format        customfield.NestedObject[PipelineStreamsFormatDataSourceModel]        `tfsdk:"format" json:"format,computed"`
	Schema        customfield.NestedObject[PipelineStreamsSchemaDataSourceModel]        `tfsdk:"schema" json:"schema,computed"`
}

type PipelineStreamsHTTPDataSourceModel struct {
	Authentication types.Bool                                                       `tfsdk:"authentication" json:"authentication,computed"`
	Enabled        types.Bool                                                       `tfsdk:"enabled" json:"enabled,computed"`
	CORS           customfield.NestedObject[PipelineStreamsHTTPCORSDataSourceModel] `tfsdk:"cors" json:"cors,computed"`
}

type PipelineStreamsHTTPCORSDataSourceModel struct {
	Origins customfield.List[types.String] `tfsdk:"origins" json:"origins,computed"`
}

type PipelineStreamsWorkerBindingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type PipelineStreamsFormatDataSourceModel struct {
	Type            types.String `tfsdk:"type" json:"type,computed"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,computed"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,computed"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,computed"`
	Compression     types.String `tfsdk:"compression" json:"compression,computed"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,computed"`
}

type PipelineStreamsSchemaDataSourceModel struct {
	Fields   customfield.NestedObjectList[PipelineStreamsSchemaFieldsDataSourceModel] `tfsdk:"fields" json:"fields,computed"`
	Format   customfield.NestedObject[PipelineStreamsSchemaFormatDataSourceModel]     `tfsdk:"format" json:"format,computed"`
	Inferred types.Bool                                                               `tfsdk:"inferred" json:"inferred,computed"`
}

type PipelineStreamsSchemaFieldsDataSourceModel struct {
	Type        types.String `tfsdk:"type" json:"type,computed"`
	MetadataKey types.String `tfsdk:"metadata_key" json:"metadata_key,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Required    types.Bool   `tfsdk:"required" json:"required,computed"`
	SqlName     types.String `tfsdk:"sql_name" json:"sql_name,computed"`
	Unit        types.String `tfsdk:"unit" json:"unit,computed"`
}

type PipelineStreamsSchemaFormatDataSourceModel struct {
	Type            types.String `tfsdk:"type" json:"type,computed"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,computed"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,computed"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,computed"`
	Compression     types.String `tfsdk:"compression" json:"compression,computed"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,computed"`
}
