// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_tiered_cache_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/regional_tiered_cache"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestRegionalTieredCacheModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*regional_tiered_cache.RegionalTieredCacheModel)(nil)
  schema := regional_tiered_cache.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
