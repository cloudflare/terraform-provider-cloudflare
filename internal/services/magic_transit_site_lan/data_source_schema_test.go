// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_lan_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/magic_transit_site_lan"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestMagicTransitSiteLANDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_transit_site_lan.MagicTransitSiteLANDataSourceModel)(nil)
	schema := magic_transit_site_lan.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
