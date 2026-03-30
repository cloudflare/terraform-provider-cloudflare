// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_sink

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/pipelines"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PipelineSinkResultDataSourceEnvelope struct {
	Result PipelineSinkDataSourceModel `json:"result,computed"`
}

type PipelineSinkDataSourceModel struct {
	ID         types.String                                                `tfsdk:"id" path:"sink_id,computed"`
	SinkID     types.String                                                `tfsdk:"sink_id" path:"sink_id,optional"`
	AccountID  types.String                                                `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt  timetypes.RFC3339                                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt timetypes.RFC3339                                           `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Name       types.String                                                `tfsdk:"name" json:"name,computed"`
	Type       types.String                                                `tfsdk:"type" json:"type,computed"`
	Config     customfield.NestedObject[PipelineSinkConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	Format     customfield.NestedObject[PipelineSinkFormatDataSourceModel] `tfsdk:"format" json:"format,computed"`
	Schema     customfield.NestedObject[PipelineSinkSchemaDataSourceModel] `tfsdk:"schema" json:"schema,computed"`
	Filter     *PipelineSinkFindOneByDataSourceModel                       `tfsdk:"filter"`
}

func (m *PipelineSinkDataSourceModel) toReadParams(_ context.Context) (params pipelines.SinkGetParams, diags diag.Diagnostics) {
	params = pipelines.SinkGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *PipelineSinkDataSourceModel) toListParams(_ context.Context) (params pipelines.SinkListParams, diags diag.Diagnostics) {
	params = pipelines.SinkListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.PipelineID.IsNull() {
		params.PipelineID = cloudflare.F(m.Filter.PipelineID.ValueString())
	}

	return
}

type PipelineSinkConfigDataSourceModel struct {
	AccountID     types.String                                                             `tfsdk:"account_id" json:"account_id,computed"`
	Bucket        types.String                                                             `tfsdk:"bucket" json:"bucket,computed"`
	FileNaming    customfield.NestedObject[PipelineSinkConfigFileNamingDataSourceModel]    `tfsdk:"file_naming" json:"file_naming,computed"`
	Jurisdiction  types.String                                                             `tfsdk:"jurisdiction" json:"jurisdiction,computed"`
	Partitioning  customfield.NestedObject[PipelineSinkConfigPartitioningDataSourceModel]  `tfsdk:"partitioning" json:"partitioning,computed"`
	Path          types.String                                                             `tfsdk:"path" json:"path,computed"`
	RollingPolicy customfield.NestedObject[PipelineSinkConfigRollingPolicyDataSourceModel] `tfsdk:"rolling_policy" json:"rolling_policy,computed"`
	TableName     types.String                                                             `tfsdk:"table_name" json:"table_name,computed"`
	Namespace     types.String                                                             `tfsdk:"namespace" json:"namespace,computed"`
}

type PipelineSinkConfigFileNamingDataSourceModel struct {
	Prefix   types.String `tfsdk:"prefix" json:"prefix,computed"`
	Strategy types.String `tfsdk:"strategy" json:"strategy,computed"`
	Suffix   types.String `tfsdk:"suffix" json:"suffix,computed"`
}

type PipelineSinkConfigPartitioningDataSourceModel struct {
	TimePattern types.String `tfsdk:"time_pattern" json:"time_pattern,computed"`
}

type PipelineSinkConfigRollingPolicyDataSourceModel struct {
	FileSizeBytes     types.Int64 `tfsdk:"file_size_bytes" json:"file_size_bytes,computed"`
	InactivitySeconds types.Int64 `tfsdk:"inactivity_seconds" json:"inactivity_seconds,computed"`
	IntervalSeconds   types.Int64 `tfsdk:"interval_seconds" json:"interval_seconds,computed"`
}

type PipelineSinkFormatDataSourceModel struct {
	Type            types.String `tfsdk:"type" json:"type,computed"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,computed"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,computed"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,computed"`
	Compression     types.String `tfsdk:"compression" json:"compression,computed"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,computed"`
}

type PipelineSinkSchemaDataSourceModel struct {
	Fields   customfield.NestedObjectList[PipelineSinkSchemaFieldsDataSourceModel] `tfsdk:"fields" json:"fields,computed"`
	Format   customfield.NestedObject[PipelineSinkSchemaFormatDataSourceModel]     `tfsdk:"format" json:"format,computed"`
	Inferred types.Bool                                                            `tfsdk:"inferred" json:"inferred,computed"`
}

type PipelineSinkSchemaFieldsDataSourceModel struct {
	Type        types.String `tfsdk:"type" json:"type,computed"`
	MetadataKey types.String `tfsdk:"metadata_key" json:"metadata_key,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Required    types.Bool   `tfsdk:"required" json:"required,computed"`
	SqlName     types.String `tfsdk:"sql_name" json:"sql_name,computed"`
	Unit        types.String `tfsdk:"unit" json:"unit,computed"`
}

type PipelineSinkSchemaFormatDataSourceModel struct {
	Type            types.String `tfsdk:"type" json:"type,computed"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,computed"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,computed"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,computed"`
	Compression     types.String `tfsdk:"compression" json:"compression,computed"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,computed"`
}

type PipelineSinkFindOneByDataSourceModel struct {
	PipelineID types.String `tfsdk:"pipeline_id" query:"pipeline_id,optional"`
}
