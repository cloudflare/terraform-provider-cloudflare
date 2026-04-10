// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_job

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpushJobResultEnvelope struct {
	Result LogpushJobModel `json:"result"`
}

type LogpushJobModel struct {
	ID                       types.Int64                   `tfsdk:"id" json:"id,computed"`
	AccountID                types.String                  `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID                   types.String                  `tfsdk:"zone_id" path:"zone_id,optional"`
	Dataset                  types.String                  `tfsdk:"dataset" json:"dataset,computed_optional"`
	DestinationConf          types.String                  `tfsdk:"destination_conf" json:"destination_conf,required"`
	Filter                   types.String                  `tfsdk:"filter" json:"filter,optional,no_refresh"`
	LogpullOptions           types.String                  `tfsdk:"logpull_options" json:"logpull_options,optional"`
	MaxUploadBytes           types.Int64                   `tfsdk:"max_upload_bytes" json:"max_upload_bytes,optional"`
	MaxUploadIntervalSeconds types.Int64                   `tfsdk:"max_upload_interval_seconds" json:"max_upload_interval_seconds,optional"`
	MaxUploadRecords         types.Int64                   `tfsdk:"max_upload_records" json:"max_upload_records,optional"`
	Name                     types.String                  `tfsdk:"name" json:"name,optional"`
	OwnershipChallenge       types.String                  `tfsdk:"ownership_challenge" json:"ownership_challenge,optional,no_refresh"`
	OutputOptions            *LogpushJobOutputOptionsModel `tfsdk:"output_options" json:"output_options,optional"`
	Enabled                  types.Bool                    `tfsdk:"enabled" json:"enabled,computed_optional"`
	Frequency                types.String                  `tfsdk:"frequency" json:"frequency,computed_optional"`
	Kind                     types.String                  `tfsdk:"kind" json:"kind,computed_optional"`
	ErrorMessage             types.String                  `tfsdk:"error_message" json:"error_message,computed"`
	LastComplete             timetypes.RFC3339             `tfsdk:"last_complete" json:"last_complete,computed" format:"date-time"`
	LastError                timetypes.RFC3339             `tfsdk:"last_error" json:"last_error,computed" format:"date-time"`
}

func (m LogpushJobModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m LogpushJobModel) MarshalJSONForUpdate(state LogpushJobModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type LogpushJobOutputOptionsModel struct {
	BatchPrefix     types.String    `tfsdk:"batch_prefix" json:"batch_prefix,optional"`
	BatchSuffix     types.String    `tfsdk:"batch_suffix" json:"batch_suffix,optional"`
	Cve2021_44228   types.Bool      `tfsdk:"cve_2021_44228" json:"CVE-2021-44228,optional"`
	FieldDelimiter  types.String    `tfsdk:"field_delimiter" json:"field_delimiter,optional"`
	FieldNames      *[]types.String `tfsdk:"field_names" json:"field_names,optional"`
	OutputType      types.String    `tfsdk:"output_type" json:"output_type,optional"`
	RecordDelimiter types.String    `tfsdk:"record_delimiter" json:"record_delimiter,optional"`
	RecordPrefix    types.String    `tfsdk:"record_prefix" json:"record_prefix,optional"`
	RecordSuffix    types.String    `tfsdk:"record_suffix" json:"record_suffix,optional"`
	RecordTemplate  types.String    `tfsdk:"record_template" json:"record_template,optional"`
	SampleRate      types.Float64   `tfsdk:"sample_rate" json:"sample_rate,optional"`
	TimestampFormat types.String    `tfsdk:"timestamp_format" json:"timestamp_format,optional"`
}
