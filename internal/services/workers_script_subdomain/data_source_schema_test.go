// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script_subdomain_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_script_subdomain"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersScriptSubdomainDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*workers_script_subdomain.WorkersScriptSubdomainDataSourceModel)(nil)
  schema := workers_script_subdomain.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
