// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestQueueModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*queue.QueueModel)(nil)
	schema := queue.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
