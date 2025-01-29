// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/stream"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestStreamDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*stream.StreamDataSourceModel)(nil)
	schema := stream.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
