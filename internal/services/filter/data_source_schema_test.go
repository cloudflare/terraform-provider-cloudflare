// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/filter"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestFilterDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*filter.FilterDataSourceModel)(nil)
  schema := filter.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
