// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueueConsumerResultEnvelope struct {
	Result QueueConsumerModel `json:"result"`
}

type QueueConsumerModel struct {
	AccountID       types.String                                         `tfsdk:"account_id" path:"account_id,required"`
	QueueID         types.String                                         `tfsdk:"queue_id" path:"queue_id,required"`
	ConsumerID      types.String                                         `tfsdk:"consumer_id" json:"consumer_id,computed"`
	Type            types.String                                         `tfsdk:"type" json:"type,required"`
	DeadLetterQueue types.String                                         `tfsdk:"dead_letter_queue" json:"dead_letter_queue,optional,no_refresh"`
	ScriptName      types.String                                         `tfsdk:"script_name" json:"script_name,computed_optional"`
	Script          types.String                                         `tfsdk:"-" json:"script,computed"` // API returns "script" but we expose "script_name"
	Settings        customfield.NestedObject[QueueConsumerSettingsModel] `tfsdk:"settings" json:"settings,computed_optional"`
	CreatedOn       timetypes.RFC3339                                    `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	QueueName       types.String                                         `tfsdk:"queue_name" json:"queue_name,computed"`
}

func (m QueueConsumerModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m QueueConsumerModel) MarshalJSONForUpdate(state QueueConsumerModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type QueueConsumerSettingsModel struct {
	BatchSize           types.Float64 `tfsdk:"batch_size" json:"batch_size,computed_optional"`
	MaxConcurrency      types.Float64 `tfsdk:"max_concurrency" json:"max_concurrency,computed_optional"`
	MaxRetries          types.Float64 `tfsdk:"max_retries" json:"max_retries,computed_optional"`
	MaxWaitTimeMs       types.Float64 `tfsdk:"max_wait_time_ms" json:"max_wait_time_ms,computed_optional"`
	RetryDelay          types.Float64 `tfsdk:"retry_delay" json:"retry_delay,computed_optional"`
	VisibilityTimeoutMs types.Float64 `tfsdk:"visibility_timeout_ms" json:"visibility_timeout_ms,computed_optional"`
}
