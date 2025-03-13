// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_behavior_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_risk_behavior"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustRiskBehaviorModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*zero_trust_risk_behavior.ZeroTrustRiskBehaviorModel)(nil)
  schema := zero_trust_risk_behavior.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
