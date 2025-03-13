// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_custom_profile_local_domain_fallback_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_custom_profile_local_domain_fallback"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDeviceCustomProfileLocalDomainFallbackDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_device_custom_profile_local_domain_fallback.ZeroTrustDeviceCustomProfileLocalDomainFallbackDataSourceModel)(nil)
	schema := zero_trust_device_custom_profile_local_domain_fallback.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
