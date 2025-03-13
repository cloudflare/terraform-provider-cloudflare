// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAPIShieldOperationModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*api_shield_operation.APIShieldOperationModel)(nil)
  schema := api_shield_operation.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
