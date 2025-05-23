// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_categories_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_categories"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustGatewayCategoriesListDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_gateway_categories.ZeroTrustGatewayCategoriesListDataSourceModel)(nil)
	schema := zero_trust_gateway_categories.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
