// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_secret_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_secret"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersSecretsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*workers_secret.WorkersSecretsDataSourceModel)(nil)
  schema := workers_secret.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
