// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_ai_controls_mcp_portal_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_ai_controls_mcp_portal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustAccessAIControlsMcpPortalModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_access_ai_controls_mcp_portal.ZeroTrustAccessAIControlsMcpPortalModel)(nil)
	schema := zero_trust_access_ai_controls_mcp_portal.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
