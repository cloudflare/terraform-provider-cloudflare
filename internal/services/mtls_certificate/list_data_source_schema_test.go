// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/mtls_certificate"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMTLSCertificatesDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*mtls_certificate.MTLSCertificatesDataSourceModel)(nil)
  schema := mtls_certificate.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
