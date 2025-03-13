// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_ownership_challenge_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_ownership_challenge"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestLogpushOwnershipChallengeModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*logpush_ownership_challenge.LogpushOwnershipChallengeModel)(nil)
  schema := logpush_ownership_challenge.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
