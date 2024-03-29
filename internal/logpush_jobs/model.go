// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_jobs

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpushJobsResultEnvelope struct {
	Result LogpushJobsModel `json:"result,computed"`
}

type LogpushJobsModel struct {
	AccountID          types.String                   `tfsdk:"account_id" path:"account_id"`
	ZoneID             types.String                   `tfsdk:"zone_id" path:"zone_id"`
	JobID              types.Int64                    `tfsdk:"job_id" path:"job_id"`
	DestinationConf    types.String                   `tfsdk:"destination_conf" json:"destination_conf"`
	Dataset            types.String                   `tfsdk:"dataset" json:"dataset"`
	Enabled            types.Bool                     `tfsdk:"enabled" json:"enabled"`
	Frequency          types.String                   `tfsdk:"frequency" json:"frequency"`
	LogpullOptions     types.String                   `tfsdk:"logpull_options" json:"logpull_options"`
	Name               types.String                   `tfsdk:"name" json:"name"`
	OutputOptions      *LogpushJobsOutputOptionsModel `tfsdk:"output_options" json:"output_options"`
	OwnershipChallenge types.String                   `tfsdk:"ownership_challenge" json:"ownership_challenge"`
	ID                 types.Int64                    `tfsdk:"id" json:"id"`
	ErrorMessage       types.String                   `tfsdk:"error_message" json:"error_message"`
	LastComplete       types.String                   `tfsdk:"last_complete" json:"last_complete"`
	LastError          types.String                   `tfsdk:"last_error" json:"last_error"`
}

type LogpushJobsOutputOptionsModel struct {
	BatchPrefix     types.String   `tfsdk:"batch_prefix" json:"batch_prefix"`
	BatchSuffix     types.String   `tfsdk:"batch_suffix" json:"batch_suffix"`
	Cve2021_4428    types.Bool     `tfsdk:"cve_2021_4428" json:"CVE-2021-4428"`
	FieldDelimiter  types.String   `tfsdk:"field_delimiter" json:"field_delimiter"`
	FieldNames      []types.String `tfsdk:"field_names" json:"field_names"`
	OutputType      types.String   `tfsdk:"output_type" json:"output_type"`
	RecordDelimiter types.String   `tfsdk:"record_delimiter" json:"record_delimiter"`
	RecordPrefix    types.String   `tfsdk:"record_prefix" json:"record_prefix"`
	RecordSuffix    types.String   `tfsdk:"record_suffix" json:"record_suffix"`
	RecordTemplate  types.String   `tfsdk:"record_template" json:"record_template"`
	SampleRate      types.Float64  `tfsdk:"sample_rate" json:"sample_rate"`
	TimestampFormat types.String   `tfsdk:"timestamp_format" json:"timestamp_format"`
}
