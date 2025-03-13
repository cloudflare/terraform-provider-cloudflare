// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/address_map"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAddressMapDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*address_map.AddressMapDataSourceModel)(nil)
  schema := address_map.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
