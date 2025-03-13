// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/resource_group"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestResourceGroupsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*resource_group.ResourceGroupsDataSourceModel)(nil)
	schema := resource_group.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
