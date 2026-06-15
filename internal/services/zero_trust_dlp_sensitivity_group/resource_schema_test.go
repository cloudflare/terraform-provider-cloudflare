// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_sensitivity_group_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_sensitivity_group"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDLPSensitivityGroupModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_dlp_sensitivity_group.ZeroTrustDLPSensitivityGroupModel)(nil)
	schema := zero_trust_dlp_sensitivity_group.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
