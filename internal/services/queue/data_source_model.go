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
	Consumers           jsontypes.Normalized           `tfsdk:"consumers" json:"consumers,computed"`
	ConsumersTotalCount jsontypes.Normalized           `tfsdk:"consumers_total_count" json:"consumers_total_count,computed"`
	CreatedOn           jsontypes.Normalized           `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn          jsontypes.Normalized           `tfsdk:"modified_on" json:"modified_on,computed"`
	Producers           jsontypes.Normalized           `tfsdk:"producers" json:"producers,computed"`
	ProducersTotalCount jsontypes.Normalized           `tfsdk:"producers_total_count" json:"producers_total_count,computed"`
	QueueName           types.String                   `tfsdk:"queue_name" json:"queue_name"`
	FindOneBy           *QueueFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type QueueFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
