// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_job

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpushJobResultEnvelope struct {
	Result LogpushJobModel `json:"result,computed"`
}

type LogpushJobModel struct {
	ID                       types.Int64                   `tfsdk:"id" json:"id,computed"`
	AccountID                types.String                  `tfsdk:"account_id" path:"account_id"`
	ZoneID                   types.String                  `tfsdk:"zone_id" path:"zone_id"`
	DestinationConf          types.String                  `tfsdk:"destination_conf" json:"destination_conf"`
	Dataset                  types.String                  `tfsdk:"dataset" json:"dataset"`
	Enabled                  types.Bool                    `tfsdk:"enabled" json:"enabled"`
	Frequency                types.String                  `tfsdk:"frequency" json:"frequency"`
	Kind                     types.String                  `tfsdk:"kind" json:"kind"`
	LogpullOptions           types.String                  `tfsdk:"logpull_options" json:"logpull_options"`
	MaxUploadBytes           types.Int64                   `tfsdk:"max_upload_bytes" json:"max_upload_bytes"`
	MaxUploadIntervalSeconds types.Int64                   `tfsdk:"max_upload_interval_seconds" json:"max_upload_interval_seconds"`
	MaxUploadRecords         types.Int64                   `tfsdk:"max_upload_records" json:"max_upload_records"`
	Name                     types.String                  `tfsdk:"name" json:"name"`
	OutputOptions            *LogpushJobOutputOptionsModel `tfsdk:"output_options" json:"output_options"`
	OwnershipChallenge       types.String                  `tfsdk:"ownership_challenge" json:"ownership_challenge"`
}

type LogpushJobOutputOptionsModel struct {
	BatchPrefix     types.String    `tfsdk:"batch_prefix" json:"batch_prefix"`
	BatchSuffix     types.String    `tfsdk:"batch_suffix" json:"batch_suffix"`
	Cve2021_4428    types.Bool      `tfsdk:"cve_2021_4428" json:"CVE-2021-4428"`
	FieldDelimiter  types.String    `tfsdk:"field_delimiter" json:"field_delimiter"`
	FieldNames      *[]types.String `tfsdk:"field_names" json:"field_names"`
	OutputType      types.String    `tfsdk:"output_type" json:"output_type"`
	RecordDelimiter types.String    `tfsdk:"record_delimiter" json:"record_delimiter"`
	RecordPrefix    types.String    `tfsdk:"record_prefix" json:"record_prefix"`
	RecordSuffix    types.String    `tfsdk:"record_suffix" json:"record_suffix"`
	RecordTemplate  types.String    `tfsdk:"record_template" json:"record_template"`
	SampleRate      types.Float64   `tfsdk:"sample_rate" json:"sample_rate"`
	TimestampFormat types.String    `tfsdk:"timestamp_format" json:"timestamp_format"`
}
