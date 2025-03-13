// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_route_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_route"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersRoutesDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*workers_route.WorkersRoutesDataSourceModel)(nil)
  schema := workers_route.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
