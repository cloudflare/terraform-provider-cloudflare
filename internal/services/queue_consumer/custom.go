package queue_consumer

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func FixInconsistentCRUDResponses(ctx context.Context, data *QueueConsumerModel) {
	// The API returns "script" in the response, but we expose "script_name" in the schema.
	// Copy the script value to script_name to ensure state consistency.
	if !data.Script.IsNull() && !data.Script.IsUnknown() && data.Script.ValueString() != "" {
		data.ScriptName = data.Script
	}

	// Ensure settings fields that are marked as Computed have known values after apply.
	// The API may not return all settings fields, so we need to set them to null if they're unknown.
	if !data.Settings.IsNull() && !data.Settings.IsUnknown() {
		settings, diags := data.Settings.Value(ctx)
		if !diags.HasError() && settings != nil {
			modified := false
			if settings.BatchSize.IsUnknown() {
				settings.BatchSize = types.Float64Null()
				modified = true
			}
			if settings.MaxConcurrency.IsUnknown() {
				settings.MaxConcurrency = types.Float64Null()
				modified = true
			}
			if settings.MaxRetries.IsUnknown() {
				settings.MaxRetries = types.Float64Null()
				modified = true
			}
			if settings.MaxWaitTimeMs.IsUnknown() {
				settings.MaxWaitTimeMs = types.Float64Null()
				modified = true
			}
			if settings.RetryDelay.IsUnknown() {
				settings.RetryDelay = types.Float64Null()
				modified = true
			}
			if settings.VisibilityTimeoutMs.IsUnknown() {
				settings.VisibilityTimeoutMs = types.Float64Null()
				modified = true
			}
			if modified {
				data.Settings = customfield.NewObjectMust(ctx, settings)
			}
		}
	}
}
