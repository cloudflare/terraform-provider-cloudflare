package queue_consumer

import "github.com/hashicorp/terraform-plugin-framework/types"

func FixInconsistentCRUDResponses(data *QueueConsumerModel) {
	// The API returns "script" in the response, but we expose "script_name" in the schema.
	// Copy the script value to script_name to ensure state consistency.
	if !data.Script.IsNull() && !data.Script.IsUnknown() && data.Script.ValueString() != "" {
		data.ScriptName = data.Script
	}

	// Ensure settings fields that are marked as Computed have known values after apply.
	// The API may not return all settings fields, so we need to set them to null if they're unknown.
	if data.Settings != nil {
		if data.Settings.BatchSize.IsUnknown() {
			data.Settings.BatchSize = types.Float64Null()
		}
		if data.Settings.MaxConcurrency.IsUnknown() {
			data.Settings.MaxConcurrency = types.Float64Null()
		}
		if data.Settings.MaxRetries.IsUnknown() {
			data.Settings.MaxRetries = types.Float64Null()
		}
		if data.Settings.MaxWaitTimeMs.IsUnknown() {
			data.Settings.MaxWaitTimeMs = types.Float64Null()
		}
		if data.Settings.RetryDelay.IsUnknown() {
			data.Settings.RetryDelay = types.Float64Null()
		}
		if data.Settings.VisibilityTimeoutMs.IsUnknown() {
			data.Settings.VisibilityTimeoutMs = types.Float64Null()
		}
	}
}
