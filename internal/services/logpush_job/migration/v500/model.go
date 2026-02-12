package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x SDKv2)
// ============================================================================

// SourceCloudflareLogpushJobModel represents the source cloudflare_logpush_job state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
// Used by UpgradeFromLegacyV0 to parse legacy state.
type SourceCloudflareLogpushJobModel struct {
	ID                       types.String                            `tfsdk:"id"`
	AccountID                types.String                            `tfsdk:"account_id"`
	ZoneID                   types.String                            `tfsdk:"zone_id"`
	Dataset                  types.String                            `tfsdk:"dataset"`
	DestinationConf          types.String                            `tfsdk:"destination_conf"`
	Enabled                  types.Bool                              `tfsdk:"enabled"`
	Filter                   types.String                            `tfsdk:"filter"`
	Frequency                types.String                            `tfsdk:"frequency"`
	Kind                     types.String                            `tfsdk:"kind"`
	LogpullOptions           types.String                            `tfsdk:"logpull_options"`
	MaxUploadBytes           types.Int64                             `tfsdk:"max_upload_bytes"`
	MaxUploadIntervalSeconds types.Int64                             `tfsdk:"max_upload_interval_seconds"`
	MaxUploadRecords         types.Int64                             `tfsdk:"max_upload_records"`
	Name                     types.String                            `tfsdk:"name"`
	OwnershipChallenge       types.String                            `tfsdk:"ownership_challenge"`
	OutputOptions            []SourceLogpushJobOutputOptionsModel    `tfsdk:"output_options"` // SDKv2 TypeList MaxItems:1 stored as array
}

// SourceLogpushJobOutputOptionsModel represents the source output_options nested structure.
// In the legacy provider, this is stored as a TypeList with MaxItems: 1, hence the array in parent model.
type SourceLogpushJobOutputOptionsModel struct {
	BatchPrefix     types.String    `tfsdk:"batch_prefix"`
	BatchSuffix     types.String    `tfsdk:"batch_suffix"`
	Cve20214428     types.Bool      `tfsdk:"cve20214428"` // Renamed to cve_2021_44228 in target
	FieldDelimiter  types.String    `tfsdk:"field_delimiter"`
	FieldNames      *[]types.String `tfsdk:"field_names"`
	OutputType      types.String    `tfsdk:"output_type"`
	RecordDelimiter types.String    `tfsdk:"record_delimiter"`
	RecordPrefix    types.String    `tfsdk:"record_prefix"`
	RecordSuffix    types.String    `tfsdk:"record_suffix"`
	RecordTemplate  types.String    `tfsdk:"record_template"`
	SampleRate      types.Float64   `tfsdk:"sample_rate"`
	TimestampFormat types.String    `tfsdk:"timestamp_format"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+ Plugin Framework)
// ============================================================================

// TargetLogpushJobModel represents the target cloudflare_logpush_job state structure (v500).
// This matches the model in the parent package (internal/services/logpush_job/model.go).
type TargetLogpushJobModel struct {
	ID                       types.Int64                             `tfsdk:"id"`
	AccountID                types.String                            `tfsdk:"account_id"`
	ZoneID                   types.String                            `tfsdk:"zone_id"`
	Dataset                  types.String                            `tfsdk:"dataset"`
	DestinationConf          types.String                            `tfsdk:"destination_conf"`
	Filter                   types.String                            `tfsdk:"filter"`
	LogpullOptions           types.String                            `tfsdk:"logpull_options"`
	MaxUploadBytes           types.Int64                             `tfsdk:"max_upload_bytes"`
	MaxUploadIntervalSeconds types.Int64                             `tfsdk:"max_upload_interval_seconds"`
	MaxUploadRecords         types.Int64                             `tfsdk:"max_upload_records"`
	Name                     types.String                            `tfsdk:"name"`
	OwnershipChallenge       types.String                            `tfsdk:"ownership_challenge"`
	OutputOptions            *TargetLogpushJobOutputOptionsModel     `tfsdk:"output_options"` // SingleNestedAttribute in v5
	Enabled                  types.Bool                              `tfsdk:"enabled"`
	Frequency                types.String                            `tfsdk:"frequency"`
	Kind                     types.String                            `tfsdk:"kind"`
	// Note: error_message, last_complete, last_error are computed-only and not included in migration
}

// TargetLogpushJobOutputOptionsModel represents the target output_options nested structure.
// In v5, this is a SingleNestedAttribute (object), not a list.
type TargetLogpushJobOutputOptionsModel struct {
	BatchPrefix     types.String    `tfsdk:"batch_prefix"`
	BatchSuffix     types.String    `tfsdk:"batch_suffix"`
	Cve2021_44228   types.Bool      `tfsdk:"cve_2021_44228"` // Renamed from cve20214428
	FieldDelimiter  types.String    `tfsdk:"field_delimiter"`
	FieldNames      *[]types.String `tfsdk:"field_names"`
	OutputType      types.String    `tfsdk:"output_type"`
	RecordDelimiter types.String    `tfsdk:"record_delimiter"`
	RecordPrefix    types.String    `tfsdk:"record_prefix"`
	RecordSuffix    types.String    `tfsdk:"record_suffix"`
	RecordTemplate  types.String    `tfsdk:"record_template"`
	SampleRate      types.Float64   `tfsdk:"sample_rate"`
	TimestampFormat types.String    `tfsdk:"timestamp_format"`
}
