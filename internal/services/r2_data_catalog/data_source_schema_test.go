// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_data_catalog_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_data_catalog"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestR2DataCatalogDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*r2_data_catalog.R2DataCatalogDataSourceModel)(nil)
	schema := r2_data_catalog.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
