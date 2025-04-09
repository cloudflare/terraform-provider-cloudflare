// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCustomPagesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_pages.CustomPagesDataSourceModel)(nil)
	schema := custom_pages.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
