// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueueResultEnvelope struct {
	Result QueueModel `json:"result,computed"`
}

type QueueModel struct {
	ID                  types.String         `tfsdk:"id" json:"-,computed"`
	QueueID             types.String         `tfsdk:"queue_id" json:"queue_id,computed"`
	AccountID           types.String         `tfsdk:"account_id" path:"account_id"`
	CreatedOn           jsontypes.Normalized `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn          jsontypes.Normalized `tfsdk:"modified_on" json:"modified_on,computed"`
	QueueName           types.String         `tfsdk:"queue_name" json:"queue_name,computed"`
	Consumers           jsontypes.Normalized `tfsdk:"consumers" json:"consumers,computed"`
	ConsumersTotalCount jsontypes.Normalized `tfsdk:"consumers_total_count" json:"consumers_total_count,computed"`
	Producers           jsontypes.Normalized `tfsdk:"producers" json:"producers,computed"`
	ProducersTotalCount jsontypes.Normalized `tfsdk:"producers_total_count" json:"producers_total_count,computed"`
}
