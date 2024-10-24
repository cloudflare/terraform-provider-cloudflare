// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_project"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPagesProjectDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pages_project.PagesProjectDataSourceModel)(nil)
	schema := pages_project.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
