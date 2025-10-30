package queue

import "github.com/hashicorp/terraform-plugin-framework/types"

func FixInconsistentCRUDResponses(data *QueueModel) {
	if data.ConsumersTotalCount.IsNull() {
		data.ConsumersTotalCount = types.Float64Value(0)
	}
	if data.ProducersTotalCount.IsNull() {
		data.ProducersTotalCount = types.Float64Value(0)
	}
}
