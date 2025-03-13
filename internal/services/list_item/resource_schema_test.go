// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestListItemModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*list_item.ListItemModel)(nil)
  schema := list_item.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
