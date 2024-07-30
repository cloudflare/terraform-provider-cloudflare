// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueuesResultListDataSourceEnvelope struct {
	Result *[]*QueuesResultDataSourceModel `json:"result,computed"`
}

type QueuesDataSourceModel struct {
	AccountID types.String                    `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                     `tfsdk:"max_items"`
	Result    *[]*QueuesResultDataSourceModel `tfsdk:"result"`
}

type QueuesResultDataSourceModel struct {
	Consumers           jsontypes.Normalized `tfsdk:"consumers" json:"consumers,computed"`
	ConsumersTotalCount jsontypes.Normalized `tfsdk:"consumers_total_count" json:"consumers_total_count,computed"`
	CreatedOn           jsontypes.Normalized `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn          jsontypes.Normalized `tfsdk:"modified_on" json:"modified_on,computed"`
	Producers           jsontypes.Normalized `tfsdk:"producers" json:"producers,computed"`
	ProducersTotalCount jsontypes.Normalized `tfsdk:"producers_total_count" json:"producers_total_count,computed"`
	QueueID             types.String         `tfsdk:"queue_id" json:"queue_id,computed"`
	QueueName           types.String         `tfsdk:"queue_name" json:"queue_name"`
}
