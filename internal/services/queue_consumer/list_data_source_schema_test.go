// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue_consumer"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestQueueConsumersDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*queue_consumer.QueueConsumersDataSourceModel)(nil)
	schema := queue_consumer.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
