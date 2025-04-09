// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_turn_app_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/calls_turn_app"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCallsTURNAppModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*calls_turn_app.CallsTURNAppModel)(nil)
  schema := calls_turn_app.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
