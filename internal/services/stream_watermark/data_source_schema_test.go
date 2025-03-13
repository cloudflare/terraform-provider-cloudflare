// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_watermark_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_watermark"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestStreamWatermarkDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*stream_watermark.StreamWatermarkDataSourceModel)(nil)
	schema := stream_watermark.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
