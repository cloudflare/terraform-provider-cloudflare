// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue_consumer"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestQueueConsumerModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*queue_consumer.QueueConsumerModel)(nil)
  schema := queue_consumer.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
