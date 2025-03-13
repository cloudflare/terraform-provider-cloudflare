// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_short_lived_certificate"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustAccessShortLivedCertificateDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*zero_trust_access_short_lived_certificate.ZeroTrustAccessShortLivedCertificateDataSourceModel)(nil)
  schema := zero_trust_access_short_lived_certificate.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
