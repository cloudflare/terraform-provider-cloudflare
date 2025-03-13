// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_job

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpushJobResultEnvelope struct {
Result LogpushJobModel `json:"result"`
}

type LogpushJobModel struct {
ID types.Int64 `tfsdk:"id" json:"id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,optional"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,optional"`
Dataset types.String `tfsdk:"dataset" json:"dataset,optional"`
DestinationConf types.String `tfsdk:"destination_conf" json:"destination_conf,required"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
Kind types.String `tfsdk:"kind" json:"kind,optional"`
LogpullOptions types.String `tfsdk:"logpull_options" json:"logpull_options,optional"`
MaxUploadBytes types.Int64 `tfsdk:"max_upload_bytes" json:"max_upload_bytes,optional"`
Name types.String `tfsdk:"name" json:"name,optional"`
OwnershipChallenge types.String `tfsdk:"ownership_challenge" json:"ownership_challenge,optional"`
Frequency types.String `tfsdk:"frequency" json:"frequency,computed_optional"`
MaxUploadIntervalSeconds types.Int64 `tfsdk:"max_upload_interval_seconds" json:"max_upload_interval_seconds,computed_optional"`
MaxUploadRecords types.Int64 `tfsdk:"max_upload_records" json:"max_upload_records,computed_optional"`
OutputOptions customfield.NestedObject[LogpushJobOutputOptionsModel] `tfsdk:"output_options" json:"output_options,computed_optional"`
ErrorMessage timetypes.RFC3339 `tfsdk:"error_message" json:"error_message,computed" format:"date-time"`
LastComplete timetypes.RFC3339 `tfsdk:"last_complete" json:"last_complete,computed" format:"date-time"`
LastError timetypes.RFC3339 `tfsdk:"last_error" json:"last_error,computed" format:"date-time"`
}

func (m LogpushJobModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m LogpushJobModel) MarshalJSONForUpdate(state LogpushJobModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

type LogpushJobOutputOptionsModel struct {
BatchPrefix types.String `tfsdk:"batch_prefix" json:"batch_prefix,computed_optional"`
BatchSuffix types.String `tfsdk:"batch_suffix" json:"batch_suffix,computed_optional"`
Cve2021_44228 types.Bool `tfsdk:"cve_2021_44228" json:"CVE-2021-44228,computed_optional"`
FieldDelimiter types.String `tfsdk:"field_delimiter" json:"field_delimiter,computed_optional"`
FieldNames *[]types.String `tfsdk:"field_names" json:"field_names,optional"`
OutputType types.String `tfsdk:"output_type" json:"output_type,computed_optional"`
RecordDelimiter types.String `tfsdk:"record_delimiter" json:"record_delimiter,computed_optional"`
RecordPrefix types.String `tfsdk:"record_prefix" json:"record_prefix,computed_optional"`
RecordSuffix types.String `tfsdk:"record_suffix" json:"record_suffix,computed_optional"`
RecordTemplate types.String `tfsdk:"record_template" json:"record_template,computed_optional"`
SampleRate types.Float64 `tfsdk:"sample_rate" json:"sample_rate,computed_optional"`
TimestampFormat types.String `tfsdk:"timestamp_format" json:"timestamp_format,computed_optional"`
}
