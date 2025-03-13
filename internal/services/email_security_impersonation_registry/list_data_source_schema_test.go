// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_impersonation_registry_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_impersonation_registry"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailSecurityImpersonationRegistriesDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*email_security_impersonation_registry.EmailSecurityImpersonationRegistriesDataSourceModel)(nil)
  schema := email_security_impersonation_registry.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
