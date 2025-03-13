// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_domain"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPagesDomainModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*pages_domain.PagesDomainModel)(nil)
  schema := pages_domain.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
