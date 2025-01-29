// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/api_shield_schema"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestAPIShieldSchemaDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_shield_schema.APIShieldSchemaDataSourceModel)(nil)
	schema := api_shield_schema.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
