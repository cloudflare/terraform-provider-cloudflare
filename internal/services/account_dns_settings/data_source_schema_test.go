// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_dns_settings_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_dns_settings"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAccountDNSSettingsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*account_dns_settings.AccountDNSSettingsDataSourceModel)(nil)
  schema := account_dns_settings.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
