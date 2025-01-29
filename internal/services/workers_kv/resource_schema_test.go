// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/workers_kv"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestWorkersKVModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_kv.WorkersKVModel)(nil)
	schema := workers_kv.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
