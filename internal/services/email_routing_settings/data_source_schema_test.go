// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_settings_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_settings"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailRoutingSettingsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*email_routing_settings.EmailRoutingSettingsDataSourceModel)(nil)
  schema := email_routing_settings.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
