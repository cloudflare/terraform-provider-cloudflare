// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/d1_database"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestD1DatabaseDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*d1_database.D1DatabaseDataSourceModel)(nil)
	schema := d1_database.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
