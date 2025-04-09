// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAPIShieldModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*api_shield.APIShieldModel)(nil)
  schema := api_shield.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
