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

type PipelineSinksResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PipelineSinksResultDataSourceModel] `json:"result,computed"`
}

type PipelineSinksDataSourceModel struct {
	AccountID  types.String                                                     `tfsdk:"account_id" path:"account_id,required"`
	PipelineID types.String                                                     `tfsdk:"pipeline_id" query:"pipeline_id,optional"`
	MaxItems   types.Int64                                                      `tfsdk:"max_items"`
	Result     customfield.NestedObjectList[PipelineSinksResultDataSourceModel] `tfsdk:"result"`
}

func (m *PipelineSinksDataSourceModel) toListParams(_ context.Context) (params pipelines.SinkListParams, diags diag.Diagnostics) {
	params = pipelines.SinkListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.PipelineID.IsNull() {
		params.PipelineID = cloudflare.F(m.PipelineID.ValueString())
	}

	return
}

type PipelineSinksResultDataSourceModel struct {
	ID         types.String                                                 `tfsdk:"id" json:"id,computed"`
	CreatedAt  timetypes.RFC3339                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt timetypes.RFC3339                                            `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Name       types.String                                                 `tfsdk:"name" json:"name,computed"`
	Type       types.String                                                 `tfsdk:"type" json:"type,computed"`
	Config     customfield.NestedObject[PipelineSinksConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	Format     customfield.NestedObject[PipelineSinksFormatDataSourceModel] `tfsdk:"format" json:"format,computed"`
	Schema     customfield.NestedObject[PipelineSinksSchemaDataSourceModel] `tfsdk:"schema" json:"schema,computed"`
}

type PipelineSinksConfigDataSourceModel struct {
	AccountID     types.String                                                              `tfsdk:"account_id" json:"account_id,computed"`
	Bucket        types.String                                                              `tfsdk:"bucket" json:"bucket,computed"`
	FileNaming    customfield.NestedObject[PipelineSinksConfigFileNamingDataSourceModel]    `tfsdk:"file_naming" json:"file_naming,computed"`
	Jurisdiction  types.String                                                              `tfsdk:"jurisdiction" json:"jurisdiction,computed"`
	Partitioning  customfield.NestedObject[PipelineSinksConfigPartitioningDataSourceModel]  `tfsdk:"partitioning" json:"partitioning,computed"`
	Path          types.String                                                              `tfsdk:"path" json:"path,computed"`
	RollingPolicy customfield.NestedObject[PipelineSinksConfigRollingPolicyDataSourceModel] `tfsdk:"rolling_policy" json:"rolling_policy,computed"`
	TableName     types.String                                                              `tfsdk:"table_name" json:"table_name,computed"`
	Namespace     types.String                                                              `tfsdk:"namespace" json:"namespace,computed"`
}

type PipelineSinksConfigFileNamingDataSourceModel struct {
	Prefix   types.String `tfsdk:"prefix" json:"prefix,computed"`
	Strategy types.String `tfsdk:"strategy" json:"strategy,computed"`
	Suffix   types.String `tfsdk:"suffix" json:"suffix,computed"`
}

type PipelineSinksConfigPartitioningDataSourceModel struct {
	TimePattern types.String `tfsdk:"time_pattern" json:"time_pattern,computed"`
}

type PipelineSinksConfigRollingPolicyDataSourceModel struct {
	FileSizeBytes     types.Int64 `tfsdk:"file_size_bytes" json:"file_size_bytes,computed"`
	InactivitySeconds types.Int64 `tfsdk:"inactivity_seconds" json:"inactivity_seconds,computed"`
	IntervalSeconds   types.Int64 `tfsdk:"interval_seconds" json:"interval_seconds,computed"`
}

type PipelineSinksFormatDataSourceModel struct {
	Type            types.String `tfsdk:"type" json:"type,computed"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,computed"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,computed"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,computed"`
	Compression     types.String `tfsdk:"compression" json:"compression,computed"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,computed"`
}

type PipelineSinksSchemaDataSourceModel struct {
	Fields   customfield.NestedObjectList[PipelineSinksSchemaFieldsDataSourceModel] `tfsdk:"fields" json:"fields,computed"`
	Format   customfield.NestedObject[PipelineSinksSchemaFormatDataSourceModel]     `tfsdk:"format" json:"format,computed"`
	Inferred types.Bool                                                             `tfsdk:"inferred" json:"inferred,computed"`
}

type PipelineSinksSchemaFieldsDataSourceModel struct {
	Type        types.String `tfsdk:"type" json:"type,computed"`
	MetadataKey types.String `tfsdk:"metadata_key" json:"metadata_key,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Required    types.Bool   `tfsdk:"required" json:"required,computed"`
	SqlName     types.String `tfsdk:"sql_name" json:"sql_name,computed"`
	Unit        types.String `tfsdk:"unit" json:"unit,computed"`
}

type PipelineSinksSchemaFormatDataSourceModel struct {
	Type            types.String `tfsdk:"type" json:"type,computed"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,computed"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,computed"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,computed"`
	Compression     types.String `tfsdk:"compression" json:"compression,computed"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,computed"`
}
