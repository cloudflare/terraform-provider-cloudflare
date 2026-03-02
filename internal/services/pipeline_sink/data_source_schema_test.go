// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_sink_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pipeline_sink"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPipelineSinkDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pipeline_sink.PipelineSinkDataSourceModel)(nil)
	schema := pipeline_sink.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
