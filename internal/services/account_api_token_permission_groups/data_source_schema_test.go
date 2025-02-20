// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_api_token_permission_groups_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_api_token_permission_groups"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAccountAPITokenPermissionGroupsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account_api_token_permission_groups.AccountAPITokenPermissionGroupsDataSourceModel)(nil)
	schema := account_api_token_permission_groups.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
