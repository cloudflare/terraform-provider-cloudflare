// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_script"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersScriptModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*workers_script.WorkersScriptModel)(nil)
  schema := workers_script.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
