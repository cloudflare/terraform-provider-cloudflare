// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZonesDataSourceModelSchemaParity(t *testing.T) {
  t.Skip("need investigation: currently broken")
  t.Parallel()
  model := (*zone.ZonesDataSourceModel)(nil)
  schema := zone.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
