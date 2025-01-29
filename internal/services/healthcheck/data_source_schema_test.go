// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/healthcheck"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestHealthcheckDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*healthcheck.HealthcheckDataSourceModel)(nil)
	schema := healthcheck.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
