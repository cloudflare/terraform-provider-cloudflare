// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_gre_tunnel_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_wan_gre_tunnel"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMagicWangreTunnelModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_wan_gre_tunnel.MagicWANGRETunnelModel)(nil)
	schema := magic_wan_gre_tunnel.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}