// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_for_platforms_dispatch_namespace"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersForPlatformsDispatchNamespaceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*workers_for_platforms_dispatch_namespace.WorkersForPlatformsDispatchNamespaceModel)(nil)
  schema := workers_for_platforms_dispatch_namespace.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
