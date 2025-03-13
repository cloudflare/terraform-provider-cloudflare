// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_catch_all_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_catch_all"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailRoutingCatchAllDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*email_routing_catch_all.EmailRoutingCatchAllDataSourceModel)(nil)
  schema := email_routing_catch_all.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
