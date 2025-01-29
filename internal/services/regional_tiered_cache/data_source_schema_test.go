// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_tiered_cache_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/regional_tiered_cache"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestRegionalTieredCacheDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*regional_tiered_cache.RegionalTieredCacheDataSourceModel)(nil)
	schema := regional_tiered_cache.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
