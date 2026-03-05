// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_sink

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PipelineSinkResultEnvelope struct {
	Result PipelineSinkModel `json:"result"`
}

type PipelineSinkModel struct {
	ID         types.String             `tfsdk:"id" json:"id,computed"`
	AccountID  types.String             `tfsdk:"account_id" path:"account_id,required"`
	Name       types.String             `tfsdk:"name" json:"name,required"`
	Type       types.String             `tfsdk:"type" json:"type,required"`
	Config     *PipelineSinkConfigModel `tfsdk:"config" json:"config,optional"`
	Format     *PipelineSinkFormatModel `tfsdk:"format" json:"format,optional"`
	Schema     *PipelineSinkSchemaModel `tfsdk:"schema" json:"schema,optional"`
	CreatedAt  timetypes.RFC3339        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt timetypes.RFC3339        `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
}

func (m PipelineSinkModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PipelineSinkModel) MarshalJSONForUpdate(state PipelineSinkModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type PipelineSinkConfigModel struct {
	AccountID     types.String                          `tfsdk:"account_id" json:"account_id,required"`
	Bucket        types.String                          `tfsdk:"bucket" json:"bucket,required"`
	Credentials   *PipelineSinkConfigCredentialsModel   `tfsdk:"credentials" json:"credentials,optional,no_refresh"`
	FileNaming    *PipelineSinkConfigFileNamingModel    `tfsdk:"file_naming" json:"file_naming,optional"`
	Jurisdiction  types.String                          `tfsdk:"jurisdiction" json:"jurisdiction,optional"`
	Partitioning  *PipelineSinkConfigPartitioningModel  `tfsdk:"partitioning" json:"partitioning,optional"`
	Path          types.String                          `tfsdk:"path" json:"path,optional"`
	RollingPolicy *PipelineSinkConfigRollingPolicyModel `tfsdk:"rolling_policy" json:"rolling_policy,optional"`
	Token         types.String                          `tfsdk:"token" json:"token,optional,no_refresh"`
	TableName     types.String                          `tfsdk:"table_name" json:"table_name,optional"`
	Namespace     types.String                          `tfsdk:"namespace" json:"namespace,optional"`
}

type PipelineSinkConfigCredentialsModel struct {
	AccessKeyID     types.String `tfsdk:"access_key_id" json:"access_key_id,required"`
	SecretAccessKey types.String `tfsdk:"secret_access_key" json:"secret_access_key,required"`
}

type PipelineSinkConfigFileNamingModel struct {
	Prefix   types.String `tfsdk:"prefix" json:"prefix,optional"`
	Strategy types.String `tfsdk:"strategy" json:"strategy,optional"`
	Suffix   types.String `tfsdk:"suffix" json:"suffix,optional"`
}

type PipelineSinkConfigPartitioningModel struct {
	TimePattern types.String `tfsdk:"time_pattern" json:"time_pattern,optional"`
}

type PipelineSinkConfigRollingPolicyModel struct {
	FileSizeBytes     types.Int64 `tfsdk:"file_size_bytes" json:"file_size_bytes,optional"`
	InactivitySeconds types.Int64 `tfsdk:"inactivity_seconds" json:"inactivity_seconds,optional"`
	IntervalSeconds   types.Int64 `tfsdk:"interval_seconds" json:"interval_seconds,optional"`
}

type PipelineSinkFormatModel struct {
	Type            types.String `tfsdk:"type" json:"type,required"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,optional"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,optional"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,optional"`
	Compression     types.String `tfsdk:"compression" json:"compression,optional"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,optional"`
}

type PipelineSinkSchemaModel struct {
	Fields   *[]*PipelineSinkSchemaFieldsModel `tfsdk:"fields" json:"fields,optional"`
	Format   *PipelineSinkSchemaFormatModel    `tfsdk:"format" json:"format,optional"`
	Inferred types.Bool                        `tfsdk:"inferred" json:"inferred,optional"`
}

type PipelineSinkSchemaFieldsModel struct {
	Type        types.String `tfsdk:"type" json:"type,required"`
	MetadataKey types.String `tfsdk:"metadata_key" json:"metadata_key,optional"`
	Name        types.String `tfsdk:"name" json:"name,optional"`
	Required    types.Bool   `tfsdk:"required" json:"required,optional"`
	SqlName     types.String `tfsdk:"sql_name" json:"sql_name,optional"`
	Unit        types.String `tfsdk:"unit" json:"unit,optional"`
}

type PipelineSinkSchemaFormatModel struct {
	Type            types.String `tfsdk:"type" json:"type,required"`
	DecimalEncoding types.String `tfsdk:"decimal_encoding" json:"decimal_encoding,optional"`
	TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,optional"`
	Unstructured    types.Bool   `tfsdk:"unstructured" json:"unstructured,optional"`
	Compression     types.String `tfsdk:"compression" json:"compression,optional"`
	RowGroupBytes   types.Int64  `tfsdk:"row_group_bytes" json:"row_group_bytes,optional"`
}
