// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_wan_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/magic_transit_site_wan"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestMagicTransitSiteWANsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_transit_site_wan.MagicTransitSiteWANsDataSourceModel)(nil)
	schema := magic_transit_site_wan.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
