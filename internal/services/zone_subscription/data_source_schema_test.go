// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_subscription"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZoneSubscriptionDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zone_subscription.ZoneSubscriptionDataSourceModel)(nil)
	schema := zone_subscription.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
