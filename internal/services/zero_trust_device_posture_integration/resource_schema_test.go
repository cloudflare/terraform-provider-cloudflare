// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_integration_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_device_posture_integration"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustDevicePostureIntegrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_device_posture_integration.ZeroTrustDevicePostureIntegrationModel)(nil)
	schema := zero_trust_device_posture_integration.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
