// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_device_default_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustDeviceDefaultProfileModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_device_default_profile.ZeroTrustDeviceDefaultProfileModel)(nil)
	schema := zero_trust_device_default_profile.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
