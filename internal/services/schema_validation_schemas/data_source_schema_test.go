// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_schemas_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/schema_validation_schemas"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSchemaValidationSchemasDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*schema_validation_schemas.SchemaValidationSchemasDataSourceModel)(nil)
	schema := schema_validation_schemas.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
