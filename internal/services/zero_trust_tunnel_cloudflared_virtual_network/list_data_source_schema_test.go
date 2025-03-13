// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_virtual_network_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_virtual_network"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustTunnelCloudflaredVirtualNetworksDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*zero_trust_tunnel_cloudflared_virtual_network.ZeroTrustTunnelCloudflaredVirtualNetworksDataSourceModel)(nil)
  schema := zero_trust_tunnel_cloudflared_virtual_network.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
