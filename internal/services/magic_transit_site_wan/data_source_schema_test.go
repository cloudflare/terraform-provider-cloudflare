// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_wan_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_transit_site_wan"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMagicTransitSiteWANDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_transit_site_wan.MagicTransitSiteWANDataSourceModel)(nil)
	schema := magic_transit_site_wan.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
