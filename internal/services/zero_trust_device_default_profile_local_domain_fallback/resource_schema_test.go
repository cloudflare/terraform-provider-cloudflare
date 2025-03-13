// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile_local_domain_fallback_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_default_profile_local_domain_fallback"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDeviceDefaultProfileLocalDomainFallbackModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*zero_trust_device_default_profile_local_domain_fallback.ZeroTrustDeviceDefaultProfileLocalDomainFallbackModel)(nil)
  schema := zero_trust_device_default_profile_local_domain_fallback.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
