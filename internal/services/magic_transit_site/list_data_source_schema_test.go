// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_transit_site"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMagicTransitSitesDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*magic_transit_site.MagicTransitSitesDataSourceModel)(nil)
  schema := magic_transit_site.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
