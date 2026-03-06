// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_sink_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pipeline_sink"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPipelineSinkModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pipeline_sink.PipelineSinkModel)(nil)
	schema := pipeline_sink.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
