// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_policy_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_shield_policy"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPageShieldPolicyModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*page_shield_policy.PageShieldPolicyModel)(nil)
  schema := page_shield_policy.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
