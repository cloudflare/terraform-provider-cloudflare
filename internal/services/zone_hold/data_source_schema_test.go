// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_hold_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zone_hold"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZoneHoldDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zone_hold.ZoneHoldDataSourceModel)(nil)
	schema := zone_hold.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
