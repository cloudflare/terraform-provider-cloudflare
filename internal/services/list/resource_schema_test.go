// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestListModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*list.ListModel)(nil)
	schema := list.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
