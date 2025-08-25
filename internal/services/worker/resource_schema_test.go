// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkerModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*worker.WorkerModel)(nil)
	schema := worker.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
