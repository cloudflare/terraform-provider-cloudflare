// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_stream_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pipeline_stream"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPipelineStreamDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pipeline_stream.PipelineStreamDataSourceModel)(nil)
	schema := pipeline_stream.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
