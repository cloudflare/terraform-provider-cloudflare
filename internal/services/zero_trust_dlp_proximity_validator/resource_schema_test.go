// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_proximity_validator_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_proximity_validator"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDLPProximityValidatorModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_dlp_proximity_validator.ZeroTrustDLPProximityValidatorModel)(nil)
	schema := zero_trust_dlp_proximity_validator.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
