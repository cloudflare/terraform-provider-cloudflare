// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_bgp_filter_profile_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_wan_bgp_filter_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMagicWanbgpFilterProfileModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_wan_bgp_filter_profile.MagicWANBGPFilterProfileModel)(nil)
	schema := magic_wan_bgp_filter_profile.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
