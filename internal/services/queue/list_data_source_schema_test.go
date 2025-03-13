// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestQueuesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*queue.QueuesDataSourceModel)(nil)
	schema := queue.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
