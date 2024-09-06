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

type QueueResultDataSourceEnvelope struct {
	Result QueueDataSourceModel `json:"result,computed"`
}

type QueueResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[QueueDataSourceModel] `json:"result,computed"`
}

type QueueDataSourceModel struct {
	AccountID           types.String                                                `tfsdk:"account_id" path:"account_id,optional"`
	QueueID             types.String                                                `tfsdk:"queue_id" path:"queue_id,computed_optional"`
	ConsumersTotalCount types.Float64                                               `tfsdk:"consumers_total_count" json:"consumers_total_count,computed"`
	CreatedOn           types.String                                                `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn          types.String                                                `tfsdk:"modified_on" json:"modified_on,computed"`
	ProducersTotalCount types.Float64                                               `tfsdk:"producers_total_count" json:"producers_total_count,computed"`
	QueueName           types.String                                                `tfsdk:"queue_name" json:"queue_name,computed"`
	Consumers           customfield.NestedObjectList[QueueConsumersDataSourceModel] `tfsdk:"consumers" json:"consumers,computed"`
	Producers           customfield.NestedObjectList[QueueProducersDataSourceModel] `tfsdk:"producers" json:"producers,computed"`
	Filter              *QueueFindOneByDataSourceModel                              `tfsdk:"filter"`
}

func (m *QueueDataSourceModel) toReadParams(_ context.Context) (params queues.QueueGetParams, diags diag.Diagnostics) {
	params = queues.QueueGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *QueueDataSourceModel) toListParams(_ context.Context) (params queues.QueueListParams, diags diag.Diagnostics) {
	params = queues.QueueListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type QueueConsumersDataSourceModel struct {
	CreatedOn   types.String                                                    `tfsdk:"created_on" json:"created_on,computed"`
	Environment types.String                                                    `tfsdk:"environment" json:"environment,computed"`
	QueueName   types.String                                                    `tfsdk:"queue_name" json:"queue_name,computed"`
	Service     types.String                                                    `tfsdk:"service" json:"service,computed"`
	Settings    customfield.NestedObject[QueueConsumersSettingsDataSourceModel] `tfsdk:"settings" json:"settings,computed"`
}

type QueueConsumersSettingsDataSourceModel struct {
	BatchSize     types.Float64 `tfsdk:"batch_size" json:"batch_size,computed"`
	MaxRetries    types.Float64 `tfsdk:"max_retries" json:"max_retries,computed"`
	MaxWaitTimeMs types.Float64 `tfsdk:"max_wait_time_ms" json:"max_wait_time_ms,computed"`
}

type QueueProducersDataSourceModel struct {
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Service     types.String `tfsdk:"service" json:"service,computed"`
}

type QueueFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
