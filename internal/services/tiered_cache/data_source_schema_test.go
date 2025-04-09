// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tiered_cache_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/tiered_cache"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestTieredCacheDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*tiered_cache.TieredCacheDataSourceModel)(nil)
  schema := tiered_cache.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
