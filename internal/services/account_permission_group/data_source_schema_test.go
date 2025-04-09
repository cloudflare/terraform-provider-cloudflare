// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_permission_group_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_permission_group"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAccountPermissionGroupDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account_permission_group.AccountPermissionGroupDataSourceModel)(nil)
	schema := account_permission_group.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
