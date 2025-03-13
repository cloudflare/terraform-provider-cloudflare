// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_key_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_key"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestStreamKeyModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*stream_key.StreamKeyModel)(nil)
  schema := stream_key.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
