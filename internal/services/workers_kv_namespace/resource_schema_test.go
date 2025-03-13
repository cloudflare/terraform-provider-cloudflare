// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_kv_namespace"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersKVNamespaceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*workers_kv_namespace.WorkersKVNamespaceModel)(nil)
  schema := workers_kv_namespace.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
