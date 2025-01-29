// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_outgoing_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/dns_zone_transfers_outgoing"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestDNSZoneTransfersOutgoingModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_zone_transfers_outgoing.DNSZoneTransfersOutgoingModel)(nil)
	schema := dns_zone_transfers_outgoing.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
