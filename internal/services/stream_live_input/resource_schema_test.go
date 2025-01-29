// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_live_input_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/stream_live_input"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestStreamLiveInputModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*stream_live_input.StreamLiveInputModel)(nil)
	schema := stream_live_input.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
