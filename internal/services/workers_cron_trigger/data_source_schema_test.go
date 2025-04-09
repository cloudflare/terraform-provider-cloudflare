// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_cron_trigger_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_cron_trigger"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersCronTriggerDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*workers_cron_trigger.WorkersCronTriggerDataSourceModel)(nil)
  schema := workers_cron_trigger.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
