// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_group_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/permission_group"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPermissionGroupDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*permission_group.PermissionGroupDataSourceModel)(nil)
	schema := permission_group.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
