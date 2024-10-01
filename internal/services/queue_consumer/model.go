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
	AccountID       types.String                                         `tfsdk:"account_id" path:"account_id,required"`
	QueueID         types.String                                         `tfsdk:"queue_id" path:"queue_id,required"`
	ConsumerID      types.String                                         `tfsdk:"consumer_id" path:"consumer_id,optional"`
	CreatedOn       types.String                                         `tfsdk:"created_on" json:"created_on,computed"`
	DeadLetterQueue types.String                                         `tfsdk:"dead_letter_queue" json:"dead_letter_queue,computed"`
	Environment     types.String                                         `tfsdk:"environment" json:"environment,computed"`
	QueueName       types.String                                         `tfsdk:"queue_name" json:"queue_name,computed"`
	ScriptName      types.String                                         `tfsdk:"script_name" json:"script_name,computed"`
	Settings        customfield.NestedObject[QueueConsumerSettingsModel] `tfsdk:"settings" json:"settings,computed"`
}

func (m QueueConsumerModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m QueueConsumerModel) MarshalJSONForUpdate(state QueueConsumerModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type QueueConsumerSettingsModel struct {
	BatchSize     types.Float64 `tfsdk:"batch_size" json:"batch_size,computed"`
	MaxRetries    types.Float64 `tfsdk:"max_retries" json:"max_retries,computed"`
	MaxWaitTimeMs types.Float64 `tfsdk:"max_wait_time_ms" json:"max_wait_time_ms,computed"`
}
