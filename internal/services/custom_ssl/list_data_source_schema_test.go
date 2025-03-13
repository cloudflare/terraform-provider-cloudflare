// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_ssl"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCustomSSLsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*custom_ssl.CustomSSLsDataSourceModel)(nil)
  schema := custom_ssl.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
