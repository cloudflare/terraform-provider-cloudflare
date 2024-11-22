// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueueConsumerResultEnvelope struct {
	Result QueueConsumerModel `json:"result"`
}

type QueueConsumerModel struct {
	AccountID  types.String                                         `tfsdk:"account_id" path:"account_id,required"`
	QueueID    types.String                                         `tfsdk:"queue_id" path:"queue_id,required"`
	ConsumerID types.String                                         `tfsdk:"consumer_id" path:"consumer_id,optional"`
	ScriptName types.String                                         `tfsdk:"script_name" json:"script_name,optional"`
	Type       types.String                                         `tfsdk:"type" json:"type,optional"`
	Settings   customfield.NestedObject[QueueConsumerSettingsModel] `tfsdk:"settings" json:"settings,computed_optional"`
	CreatedOn  types.String                                         `tfsdk:"created_on" json:"created_on,computed"`
	Script     types.String                                         `tfsdk:"script" json:"script,computed"`
}

func (m QueueConsumerModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m QueueConsumerModel) MarshalJSONForUpdate(state QueueConsumerModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type QueueConsumerSettingsModel struct {
	BatchSize           types.Float64 `tfsdk:"batch_size" json:"batch_size,optional"`
	MaxConcurrency      types.Float64 `tfsdk:"max_concurrency" json:"max_concurrency,optional"`
	MaxRetries          types.Float64 `tfsdk:"max_retries" json:"max_retries,optional"`
	MaxWaitTimeMs       types.Float64 `tfsdk:"max_wait_time_ms" json:"max_wait_time_ms,optional"`
	RetryDelay          types.Float64 `tfsdk:"retry_delay" json:"retry_delay,optional"`
	VisibilityTimeoutMs types.Float64 `tfsdk:"visibility_timeout_ms" json:"visibility_timeout_ms,optional"`
}
