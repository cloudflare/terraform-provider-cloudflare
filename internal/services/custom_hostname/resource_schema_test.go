// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_hostname"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCustomHostnameModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*custom_hostname.CustomHostnameModel)(nil)
  schema := custom_hostname.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
