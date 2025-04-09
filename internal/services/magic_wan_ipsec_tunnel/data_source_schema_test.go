// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_ipsec_tunnel_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_wan_ipsec_tunnel"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMagicWanipsecTunnelDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*magic_wan_ipsec_tunnel.MagicWANIPSECTunnelDataSourceModel)(nil)
  schema := magic_wan_ipsec_tunnel.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
