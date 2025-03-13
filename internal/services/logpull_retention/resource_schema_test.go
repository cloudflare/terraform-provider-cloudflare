// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpull_retention_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpull_retention"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestLogpullRetentionModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*logpull_retention.LogpullRetentionModel)(nil)
  schema := logpull_retention.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
