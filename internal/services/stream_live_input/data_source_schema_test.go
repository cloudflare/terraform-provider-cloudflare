// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_live_input_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_live_input"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestStreamLiveInputDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*stream_live_input.StreamLiveInputDataSourceModel)(nil)
	schema := stream_live_input.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
