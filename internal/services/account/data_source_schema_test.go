// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/account"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAccountDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*account.AccountDataSourceModel)(nil)
  schema := account.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
