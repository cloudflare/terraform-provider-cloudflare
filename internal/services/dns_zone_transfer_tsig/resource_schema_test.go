// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfer_tsig_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfer_tsig"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSZoneTransferTSIGModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_zone_transfer_tsig.DNSZoneTransferTSIGModel)(nil)
	schema := dns_zone_transfer_tsig.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
