// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_role_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_role"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAccountRoleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account_role.AccountRoleDataSourceModel)(nil)
	schema := account_role.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
