// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueueResultDataSourceEnvelope struct {
	Result QueueDataSourceModel `json:"result,computed"`
}

type QueueResultListDataSourceEnvelope struct {
	Result *[]*QueueDataSourceModel `json:"result,computed"`
}

type QueueDataSourceModel struct {
	AccountID           types.String                   `tfsdk:"account_id" path:"account_id"`
	QueueID             types.String                   `tfsdk:"queue_id" path:"queue_id"`
	ConsumersTotalCount types.Float64                  `tfsdk:"consumers_total_count" json:"consumers_total_count,computed"`
	CreatedOn           types.String                   `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn          types.String                   `tfsdk:"modified_on" json:"modified_on,computed"`
	ProducersTotalCount types.Float64                  `tfsdk:"producers_total_count" json:"producers_total_count,computed"`
	Consumers           jsontypes.Normalized           `tfsdk:"consumers" json:"consumers,computed"`
	Producers           jsontypes.Normalized           `tfsdk:"producers" json:"producers,computed"`
	QueueName           types.String                   `tfsdk:"queue_name" json:"queue_name"`
	Filter              *QueueFindOneByDataSourceModel `tfsdk:"filter"`
}

type QueueFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
