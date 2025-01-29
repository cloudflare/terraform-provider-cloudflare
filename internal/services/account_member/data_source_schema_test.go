// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_member"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAccountMemberDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account_member.AccountMemberDataSourceModel)(nil)
	schema := account_member.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
