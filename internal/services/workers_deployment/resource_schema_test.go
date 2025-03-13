// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_deployment_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_deployment"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersDeploymentModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*workers_deployment.WorkersDeploymentModel)(nil)
  schema := workers_deployment.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
