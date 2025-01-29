// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_virtual_network_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_tunnel_cloudflared_virtual_network"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustTunnelCloudflaredVirtualNetworkDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_tunnel_cloudflared_virtual_network.ZeroTrustTunnelCloudflaredVirtualNetworkDataSourceModel)(nil)
	schema := zero_trust_tunnel_cloudflared_virtual_network.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
