// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_schema"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAPIShieldSchemasDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_shield_schema.APIShieldSchemasDataSourceModel)(nil)
	schema := api_shield_schema.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
