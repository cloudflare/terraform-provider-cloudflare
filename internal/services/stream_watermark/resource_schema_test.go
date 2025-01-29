// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_watermark_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/stream_watermark"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestStreamWatermarkModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*stream_watermark.StreamWatermarkModel)(nil)
	schema := stream_watermark.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
