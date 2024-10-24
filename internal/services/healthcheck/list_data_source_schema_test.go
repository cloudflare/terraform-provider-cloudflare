// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/healthcheck"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestHealthchecksDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*healthcheck.HealthchecksDataSourceModel)(nil)
	schema := healthcheck.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
