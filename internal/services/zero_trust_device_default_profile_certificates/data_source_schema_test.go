// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile_certificates_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_default_profile_certificates"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDeviceDefaultProfileCertificatesDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*zero_trust_device_default_profile_certificates.ZeroTrustDeviceDefaultProfileCertificatesDataSourceModel)(nil)
  schema := zero_trust_device_default_profile_certificates.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
