// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_connections_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_shield_connections"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPageShieldConnectionsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*page_shield_connections.PageShieldConnectionsDataSourceModel)(nil)
  schema := page_shield_connections.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
