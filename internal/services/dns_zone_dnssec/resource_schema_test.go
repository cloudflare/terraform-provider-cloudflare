// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_dnssec_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_dnssec"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSZoneDNSSECModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_zone_dnssec.DNSZoneDNSSECModel)(nil)
	schema := dns_zone_dnssec.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
