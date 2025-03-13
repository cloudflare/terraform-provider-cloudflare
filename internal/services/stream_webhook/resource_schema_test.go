// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_webhook_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_webhook"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestStreamWebhookModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*stream_webhook.StreamWebhookModel)(nil)
  schema := stream_webhook.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
