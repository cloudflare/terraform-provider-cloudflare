// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_route_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_route"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustTunnelCloudflaredRouteDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_tunnel_cloudflared_route.ZeroTrustTunnelCloudflaredRouteDataSourceModel)(nil)
	schema := zero_trust_tunnel_cloudflared_route.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
