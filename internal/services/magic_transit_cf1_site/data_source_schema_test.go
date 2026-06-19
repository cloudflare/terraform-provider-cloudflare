// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_cf1_site_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_transit_cf1_site"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMagicTransitCf1SiteDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_transit_cf1_site.MagicTransitCf1SiteDataSourceModel)(nil)
	schema := magic_transit_cf1_site.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
