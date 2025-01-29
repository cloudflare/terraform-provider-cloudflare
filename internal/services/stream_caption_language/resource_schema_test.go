// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_caption_language_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/stream_caption_language"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestStreamCaptionLanguageModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*stream_caption_language.StreamCaptionLanguageModel)(nil)
	schema := stream_caption_language.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
