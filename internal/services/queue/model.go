// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueueResultEnvelope struct {
	Result QueueModel `json:"result"`
}

type QueueModel struct {
	ID                  types.String                                      `tfsdk:"id" json:"-,computed"`
	QueueID             types.String                                      `tfsdk:"queue_id" json:"queue_id,computed"`
	AccountID           types.String                                      `tfsdk:"account_id" path:"account_id,required"`
	QueueName           types.String                                      `tfsdk:"queue_name" json:"queue_name,required"`
	Settings            customfield.NestedObject[QueueSettingsModel]      `tfsdk:"settings" json:"settings,computed_optional"`
	ConsumersTotalCount types.Float64                                     `tfsdk:"consumers_total_count" json:"consumers_total_count,computed"`
	CreatedOn           types.String                                      `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn          types.String                                      `tfsdk:"modified_on" json:"modified_on,computed"`
	ProducersTotalCount types.Float64                                     `tfsdk:"producers_total_count" json:"producers_total_count,computed"`
	Consumers           customfield.NestedObjectList[QueueConsumersModel] `tfsdk:"consumers" json:"consumers,computed"`
	Producers           customfield.NestedObjectList[QueueProducersModel] `tfsdk:"producers" json:"producers,computed"`
}

func (m QueueModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m QueueModel) MarshalJSONForUpdate(state QueueModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type QueueSettingsModel struct {
	DeliveryDelay          types.Float64 `tfsdk:"delivery_delay" json:"delivery_delay,optional"`
	MessageRetentionPeriod types.Float64 `tfsdk:"message_retention_period" json:"message_retention_period,optional"`
}

type QueueConsumersModel struct {
	ConsumerID types.String                                          `tfsdk:"consumer_id" json:"consumer_id,computed"`
	CreatedOn  types.String                                          `tfsdk:"created_on" json:"created_on,computed"`
	QueueID    types.String                                          `tfsdk:"queue_id" json:"queue_id,computed"`
	Script     types.String                                          `tfsdk:"script" json:"script,computed"`
	ScriptName types.String                                          `tfsdk:"script_name" json:"script_name,computed"`
	Settings   customfield.NestedObject[QueueConsumersSettingsModel] `tfsdk:"settings" json:"settings,computed"`
	Type       types.String                                          `tfsdk:"type" json:"type,computed"`
}

type QueueConsumersSettingsModel struct {
	BatchSize           types.Float64 `tfsdk:"batch_size" json:"batch_size,computed"`
	MaxConcurrency      types.Float64 `tfsdk:"max_concurrency" json:"max_concurrency,computed"`
	MaxRetries          types.Float64 `tfsdk:"max_retries" json:"max_retries,computed"`
	MaxWaitTimeMs       types.Float64 `tfsdk:"max_wait_time_ms" json:"max_wait_time_ms,computed"`
	RetryDelay          types.Float64 `tfsdk:"retry_delay" json:"retry_delay,computed"`
	VisibilityTimeoutMs types.Float64 `tfsdk:"visibility_timeout_ms" json:"visibility_timeout_ms,computed"`
}

type QueueProducersModel struct {
	Script     types.String `tfsdk:"script" json:"script,computed"`
	Type       types.String `tfsdk:"type" json:"type,computed"`
	BucketName types.String `tfsdk:"bucket_name" json:"bucket_name,computed"`
}
