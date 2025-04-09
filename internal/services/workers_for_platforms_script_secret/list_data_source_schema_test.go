// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_script_secret_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_for_platforms_script_secret"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersForPlatformsScriptSecretsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*workers_for_platforms_script_secret.WorkersForPlatformsScriptSecretsDataSourceModel)(nil)
  schema := workers_for_platforms_script_secret.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
