// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_tracing_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_tracing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZoneTracingDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zone_tracing.ZoneTracingDataSourceModel)(nil)
	schema := zone_tracing.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
