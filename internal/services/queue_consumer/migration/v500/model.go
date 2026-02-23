package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceQueueConsumerModel represents the queue_consumer state at schema_version=0.
// This resource was introduced in v5 with no v4 predecessor; this model captures the
// state as it existed before explicit schema versioning was applied.
type SourceQueueConsumerModel struct {
	AccountID       types.String                      `tfsdk:"account_id"`
	QueueID         types.String                      `tfsdk:"queue_id"`
	ConsumerID      types.String                      `tfsdk:"consumer_id"`
	DeadLetterQueue types.String                      `tfsdk:"dead_letter_queue"`
	ScriptName      types.String                      `tfsdk:"script_name"`
	Type            types.String                      `tfsdk:"type"`
	Settings        *SourceQueueConsumerSettingsModel `tfsdk:"settings"`
	CreatedOn       types.String                      `tfsdk:"created_on"`
	Script          types.String                      `tfsdk:"script"`
}

// SourceQueueConsumerSettingsModel represents the settings nested object at schema_version=0.
type SourceQueueConsumerSettingsModel struct {
	BatchSize           types.Float64 `tfsdk:"batch_size"`
	MaxConcurrency      types.Float64 `tfsdk:"max_concurrency"`
	MaxRetries          types.Float64 `tfsdk:"max_retries"`
	MaxWaitTimeMs       types.Float64 `tfsdk:"max_wait_time_ms"`
	RetryDelay          types.Float64 `tfsdk:"retry_delay"`
	VisibilityTimeoutMs types.Float64 `tfsdk:"visibility_timeout_ms"`
}
