package v500

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function handles all field transformations from v4 SDKv2 to v5 Plugin Framework.
func Transform(ctx context.Context, source SourceCloudflareLogpushJobModel) (*TargetLogpushJobModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.Dataset.IsNull() || source.Dataset.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"dataset is required for logpush_job migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.DestinationConf.IsNull() || source.DestinationConf.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"destination_conf is required for logpush_job migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Convert ID from String (v4) to Int64 (v5)
	var targetID types.Int64
	if !source.ID.IsNull() && !source.ID.IsUnknown() {
		idStr := source.ID.ValueString()
		idInt, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			diags.AddError(
				"Invalid ID",
				"Failed to convert ID from string to int64: "+err.Error(),
			)
			return nil, diags
		}
		targetID = types.Int64Value(idInt)
	} else {
		targetID = types.Int64Null()
	}

	// Step 3: Initialize target with direct copies
	target := &TargetLogpushJobModel{
		ID:                 targetID,
		AccountID:          source.AccountID,
		ZoneID:             source.ZoneID,
		Dataset:            source.Dataset,
		DestinationConf:    source.DestinationConf,
		Enabled:            source.Enabled,
		Frequency:          source.Frequency,
		OwnershipChallenge: source.OwnershipChallenge,
	}

	// Step 4: Handle empty string → null transformations
	// v4 sets these to "" when not configured, v5 uses null
	target.Filter = emptyStringToNull(source.Filter)
	target.LogpullOptions = emptyStringToNull(source.LogpullOptions)
	target.Name = emptyStringToNull(source.Name)

	// Step 5: Handle kind field - "instant-logs" no longer valid in v5
	if !source.Kind.IsNull() && source.Kind.ValueString() == "instant-logs" {
		// Remove "instant-logs" value (set to null)
		target.Kind = types.StringNull()
	} else {
		target.Kind = source.Kind
	}

	// Step 6: Handle integer fields - convert 0 to null (0 means "not set" in v5)
	target.MaxUploadBytes = zeroInt64ToNull(source.MaxUploadBytes)
	target.MaxUploadRecords = zeroInt64ToNull(source.MaxUploadRecords)
	target.MaxUploadIntervalSeconds = zeroInt64ToNull(source.MaxUploadIntervalSeconds)

	// Step 7: Handle output_options
	// v4 SDKv2 TypeList MaxItems:1 is stored as a JSON array [{...}].
	// TransformState is a no-op for this resource, so we receive the raw v4 array.
	// Extract the first element (if present) and transform to v5 object format.
	if len(source.OutputOptions) > 0 {
		targetOutputOptions, err := transformOutputOptions(ctx, source.OutputOptions[0])
		if err.HasError() {
			diags.Append(err...)
			return nil, diags
		}
		target.OutputOptions = targetOutputOptions
	} else {
		target.OutputOptions = nil
	}

	// Step 8: Computed-only fields are NOT migrated (error_message, last_complete, last_error)
	// These will be populated by API on next refresh

	return target, diags
}

// transformOutputOptions transforms a single output_options element from v4 to v5 format.
// Handles field rename (cve20214428 → cve_2021_44228) and preserves v4 defaults.
func transformOutputOptions(ctx context.Context, source SourceLogpushJobOutputOptionsModel) (*TargetLogpushJobOutputOptionsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetLogpushJobOutputOptionsModel{
		BatchPrefix:     source.BatchPrefix,
		BatchSuffix:     source.BatchSuffix,
		FieldNames:      source.FieldNames,
		OutputType:      source.OutputType,
		RecordDelimiter: source.RecordDelimiter,
		RecordTemplate:  source.RecordTemplate,
		SampleRate:      source.SampleRate,
	}

	// Field rename: cve20214428 → cve_2021_44228
	target.Cve2021_44228 = source.Cve20214428

	// Preserve v4 schema defaults if not present in source
	// v5 schema has NO defaults for these fields, so we must add them to prevent drift
	target.FieldDelimiter = preserveV4Default(source.FieldDelimiter, ",")
	target.RecordPrefix = preserveV4Default(source.RecordPrefix, "{")
	target.RecordSuffix = preserveV4Default(source.RecordSuffix, "}\n")
	target.TimestampFormat = preserveV4Default(source.TimestampFormat, "unixnano")

	// Set default for merge_subrequests (not present in v4 state, default is false)
	target.MergeSubrequests = types.BoolValue(false)

	// cve_2021_44228 has default false in both v4 and v5, so no special handling needed
	// sample_rate has default 1.0 in v4, but we preserve the actual value from state

	return target, diags
}

// Helper: Convert empty string to null
func emptyStringToNull(val types.String) types.String {
	if !val.IsNull() && val.ValueString() == "" {
		return types.StringNull()
	}
	return val
}

// Helper: Convert 0 to null for Int64 fields
func zeroInt64ToNull(val types.Int64) types.Int64 {
	if !val.IsNull() && val.ValueInt64() == 0 {
		return types.Int64Null()
	}
	return val
}

// Helper: Preserve v4 default value if field is null or empty
func preserveV4Default(val types.String, defaultValue string) types.String {
	if val.IsNull() || (!val.IsNull() && val.ValueString() == "") {
		return types.StringValue(defaultValue)
	}
	return val
}
