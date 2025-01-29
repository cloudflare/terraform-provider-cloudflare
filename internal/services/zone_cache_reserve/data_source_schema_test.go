// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_reserve_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zone_cache_reserve"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZoneCacheReserveDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zone_cache_reserve.ZoneCacheReserveDataSourceModel)(nil)
	schema := zone_cache_reserve.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
