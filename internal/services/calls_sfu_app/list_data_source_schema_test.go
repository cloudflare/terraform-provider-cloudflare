// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_sfu_app_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/calls_sfu_app"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCallsSFUAppsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*calls_sfu_app.CallsSFUAppsDataSourceModel)(nil)
  schema := calls_sfu_app.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
