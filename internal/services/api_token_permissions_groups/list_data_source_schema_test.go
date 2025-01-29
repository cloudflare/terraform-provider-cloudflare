// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token_permissions_groups_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/api_token_permissions_groups"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestAPITokenPermissionsGroupsListDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_token_permissions_groups.APITokenPermissionsGroupsListDataSourceModel)(nil)
	schema := api_token_permissions_groups.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
