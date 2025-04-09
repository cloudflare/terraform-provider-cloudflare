// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_dataset_job_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_dataset_job"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestLogpushDatasetJobDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*logpush_dataset_job.LogpushDatasetJobDataSourceModel)(nil)
  schema := logpush_dataset_job.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
