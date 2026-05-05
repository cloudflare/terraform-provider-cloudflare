// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_category_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_resource_library_category"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustResourceLibraryCategoryDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_resource_library_category.ZeroTrustResourceLibraryCategoryDataSourceModel)(nil)
	schema := zero_trust_resource_library_category.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
