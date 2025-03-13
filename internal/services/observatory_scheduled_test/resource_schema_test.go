// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/observatory_scheduled_test"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestObservatoryScheduledTestModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*observatory_scheduled_test.ObservatoryScheduledTestModel)(nil)
  schema := observatory_scheduled_test.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
