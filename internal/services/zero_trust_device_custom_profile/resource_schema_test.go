// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_custom_profile_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_custom_profile"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDeviceCustomProfileModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*zero_trust_device_custom_profile.ZeroTrustDeviceCustomProfileModel)(nil)
  schema := zero_trust_device_custom_profile.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
