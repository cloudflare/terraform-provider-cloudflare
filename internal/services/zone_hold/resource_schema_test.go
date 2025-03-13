// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_hold_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_hold"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZoneHoldModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*zone_hold.ZoneHoldModel)(nil)
  schema := zone_hold.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
