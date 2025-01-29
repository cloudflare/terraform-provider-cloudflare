// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_dns_location"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustDNSLocationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_dns_location.ZeroTrustDNSLocationModel)(nil)
	schema := zero_trust_dns_location.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
