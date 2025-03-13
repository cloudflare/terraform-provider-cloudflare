// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_virtual_network_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_virtual_network"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustTunnelCloudflaredVirtualNetworkModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_tunnel_cloudflared_virtual_network.ZeroTrustTunnelCloudflaredVirtualNetworkModel)(nil)
	schema := zero_trust_tunnel_cloudflared_virtual_network.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
