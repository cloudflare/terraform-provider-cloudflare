// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/pipeline"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPipelineDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*pipeline.PipelineDataSourceModel)(nil)
  schema := pipeline.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
