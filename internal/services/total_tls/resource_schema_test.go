// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package total_tls_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/total_tls"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestTotalTLSModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*total_tls.TotalTLSModel)(nil)
  schema := total_tls.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
