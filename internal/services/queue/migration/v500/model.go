package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareQueueModel represents the source cloudflare_queue state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
// Used by UpgradeFromLegacyV0 to parse legacy state.
type SourceCloudflareQueueModel struct {
	ID        types.String `tfsdk:"id"`
	AccountID types.String `tfsdk:"account_id"`
	Name      types.String `tfsdk:"name"`
}

// TargetQueueModel represents the target cloudflare_queue state structure (v500).
// Must match the v5 QueueModel structure exactly.
type TargetQueueModel struct {
	ID                  types.String                                                  `tfsdk:"id"`
	QueueID             types.String                                                  `tfsdk:"queue_id"`
	AccountID           types.String                                                  `tfsdk:"account_id"`
	QueueName           types.String                                                  `tfsdk:"queue_name"`
	Settings            customfield.NestedObject[TargetQueueSettingsModel]            `tfsdk:"settings"`
	ConsumersTotalCount types.Float64                                                 `tfsdk:"consumers_total_count"`
	CreatedOn           types.String                                                  `tfsdk:"created_on"`
	ModifiedOn          types.String                                                  `tfsdk:"modified_on"`
	ProducersTotalCount types.Float64                                                 `tfsdk:"producers_total_count"`
	Consumers           customfield.NestedObjectList[TargetQueueConsumersModel]       `tfsdk:"consumers"`
	Producers           customfield.NestedObjectList[TargetQueueProducersModel]       `tfsdk:"producers"`
}

// TargetQueueSettingsModel represents the target settings nested object (v500).
// Must match QueueSettingsModel structure exactly.
type TargetQueueSettingsModel struct {
	DeliveryDelay          types.Float64 `tfsdk:"delivery_delay"`
	DeliveryPaused         types.Bool    `tfsdk:"delivery_paused"`
	MessageRetentionPeriod types.Float64 `tfsdk:"message_retention_period"`
}

// TargetQueueConsumersModel represents a consumer in the target consumers list (v500).
// Must match QueueConsumersModel structure exactly.
type TargetQueueConsumersModel struct {
	ConsumerID      types.String                                                 `tfsdk:"consumer_id"`
	CreatedOn       types.String                                                 `tfsdk:"created_on"`
	DeadLetterQueue types.String                                                 `tfsdk:"dead_letter_queue"`
	QueueName       types.String                                                 `tfsdk:"queue_name"`
	ScriptName      types.String                                                 `tfsdk:"script_name"`
	Settings        customfield.NestedObject[TargetQueueConsumersSettingsModel]  `tfsdk:"settings"`
	Type            types.String                                                 `tfsdk:"type"`
}

// TargetQueueConsumersSettingsModel represents the consumer settings nested object (v500).
// Must match QueueConsumersSettingsModel structure exactly.
type TargetQueueConsumersSettingsModel struct {
	BatchSize           types.Float64 `tfsdk:"batch_size"`
	MaxConcurrency      types.Float64 `tfsdk:"max_concurrency"`
	MaxRetries          types.Float64 `tfsdk:"max_retries"`
	MaxWaitTimeMs       types.Float64 `tfsdk:"max_wait_time_ms"`
	RetryDelay          types.Float64 `tfsdk:"retry_delay"`
	VisibilityTimeoutMs types.Float64 `tfsdk:"visibility_timeout_ms"`
}

// TargetQueueProducersModel represents a producer in the target producers list (v500).
// Must match QueueProducersModel structure exactly.
type TargetQueueProducersModel struct {
	Script     types.String `tfsdk:"script"`
	Type       types.String `tfsdk:"type"`
	BucketName types.String `tfsdk:"bucket_name"`
}
