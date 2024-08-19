// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueueResultEnvelope struct {
	Result QueueModel `json:"result,computed"`
}

type QueueModel struct {
	ID                  types.String            `tfsdk:"id" json:"-,computed"`
	QueueID             types.String            `tfsdk:"queue_id" json:"queue_id,computed"`
	AccountID           types.String            `tfsdk:"account_id" path:"account_id"`
	QueueName           types.String            `tfsdk:"queue_name" json:"queue_name"`
	ConsumersTotalCount types.Float64           `tfsdk:"consumers_total_count" json:"consumers_total_count,computed"`
	CreatedOn           types.String            `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn          types.String            `tfsdk:"modified_on" json:"modified_on,computed"`
	ProducersTotalCount types.Float64           `tfsdk:"producers_total_count" json:"producers_total_count,computed"`
	Consumers           *[]*QueueConsumersModel `tfsdk:"consumers" json:"consumers,computed"`
	Producers           *[]*QueueProducersModel `tfsdk:"producers" json:"producers,computed"`
}

type QueueConsumersModel struct {
	CreatedOn   types.String                 `tfsdk:"created_on" json:"created_on,computed"`
	Environment types.String                 `tfsdk:"environment" json:"environment,computed"`
	QueueName   types.String                 `tfsdk:"queue_name" json:"queue_name,computed"`
	Service     types.String                 `tfsdk:"service" json:"service,computed"`
	Settings    *QueueConsumersSettingsModel `tfsdk:"settings" json:"settings"`
}

type QueueConsumersSettingsModel struct {
	BatchSize     types.Float64 `tfsdk:"batch_size" json:"batch_size"`
	MaxRetries    types.Float64 `tfsdk:"max_retries" json:"max_retries"`
	MaxWaitTimeMs types.Float64 `tfsdk:"max_wait_time_ms" json:"max_wait_time_ms"`
}

type QueueProducersModel struct {
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Service     types.String `tfsdk:"service" json:"service,computed"`
}
