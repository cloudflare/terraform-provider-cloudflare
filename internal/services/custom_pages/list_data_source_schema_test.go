// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCustomPagesListDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_pages.CustomPagesListDataSourceModel)(nil)
	schema := custom_pages.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
