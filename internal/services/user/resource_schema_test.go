// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/user"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestUserModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*user.UserModel)(nil)
  schema := user.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
