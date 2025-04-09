// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestStreamsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*stream.StreamsDataSourceModel)(nil)
	schema := stream.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
