// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/queues"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueuesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[QueuesResultDataSourceModel] `json:"result,computed"`
}

type QueuesDataSourceModel struct {
	AccountID types.String                                              `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                               `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[QueuesResultDataSourceModel] `tfsdk:"result"`
}

func (m *QueuesDataSourceModel) toListParams(_ context.Context) (params queues.QueueListParams, diags diag.Diagnostics) {
	params = queues.QueueListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type QueuesResultDataSourceModel struct {
	Consumers           customfield.NestedObjectList[QueuesConsumersDataSourceModel] `tfsdk:"consumers" json:"consumers,computed"`
	ConsumersTotalCount types.Float64                                                `tfsdk:"consumers_total_count" json:"consumers_total_count,computed"`
	CreatedOn           types.String                                                 `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn          types.String                                                 `tfsdk:"modified_on" json:"modified_on,computed"`
	Producers           customfield.NestedObjectList[QueuesProducersDataSourceModel] `tfsdk:"producers" json:"producers,computed"`
	ProducersTotalCount types.Float64                                                `tfsdk:"producers_total_count" json:"producers_total_count,computed"`
	QueueID             types.String                                                 `tfsdk:"queue_id" json:"queue_id,computed"`
	QueueName           types.String                                                 `tfsdk:"queue_name" json:"queue_name,computed"`
}

type QueuesConsumersDataSourceModel struct {
	CreatedOn   types.String                                                     `tfsdk:"created_on" json:"created_on,computed"`
	Environment types.String                                                     `tfsdk:"environment" json:"environment,computed"`
	QueueName   types.String                                                     `tfsdk:"queue_name" json:"queue_name,computed"`
	Service     types.String                                                     `tfsdk:"service" json:"service,computed"`
	Settings    customfield.NestedObject[QueuesConsumersSettingsDataSourceModel] `tfsdk:"settings" json:"settings,computed"`
}

type QueuesConsumersSettingsDataSourceModel struct {
	BatchSize     types.Float64 `tfsdk:"batch_size" json:"batch_size,computed"`
	MaxRetries    types.Float64 `tfsdk:"max_retries" json:"max_retries,computed"`
	MaxWaitTimeMs types.Float64 `tfsdk:"max_wait_time_ms" json:"max_wait_time_ms,computed"`
}

type QueuesProducersDataSourceModel struct {
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Service     types.String `tfsdk:"service" json:"service,computed"`
}
