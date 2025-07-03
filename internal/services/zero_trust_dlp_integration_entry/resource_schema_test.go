// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_integration_entry_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_integration_entry"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDLPIntegrationEntryModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_dlp_integration_entry.ZeroTrustDLPIntegrationEntryModel)(nil)
	schema := zero_trust_dlp_integration_entry.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
