// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_reserve_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_cache_reserve"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZoneCacheReserveModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*zone_cache_reserve.ZoneCacheReserveModel)(nil)
  schema := zone_cache_reserve.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
