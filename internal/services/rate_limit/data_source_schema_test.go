// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/rate_limit"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestRateLimitDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*rate_limit.RateLimitDataSourceModel)(nil)
	schema := rate_limit.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
