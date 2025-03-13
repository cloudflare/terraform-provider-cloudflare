// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_caption_language_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_caption_language"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestStreamCaptionLanguageDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*stream_caption_language.StreamCaptionLanguageDataSourceModel)(nil)
  schema := stream_caption_language.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
