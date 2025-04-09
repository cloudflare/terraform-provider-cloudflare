// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_list"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustListModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*zero_trust_list.ZeroTrustListModel)(nil)
  schema := zero_trust_list.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
