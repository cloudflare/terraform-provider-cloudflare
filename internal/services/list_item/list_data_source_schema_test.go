// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestListItemsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*list_item.ListItemsDataSourceModel)(nil)
	schema := list_item.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
