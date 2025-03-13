// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_member"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAccountMembersDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*account_member.AccountMembersDataSourceModel)(nil)
  schema := account_member.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
