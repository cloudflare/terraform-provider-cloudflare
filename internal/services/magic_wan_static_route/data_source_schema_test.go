// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_static_route_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/magic_wan_static_route"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestMagicWANStaticRouteDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_wan_static_route.MagicWANStaticRouteDataSourceModel)(nil)
	schema := magic_wan_static_route.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
