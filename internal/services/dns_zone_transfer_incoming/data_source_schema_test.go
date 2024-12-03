// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfer_incoming_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfer_incoming"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSZoneTransferIncomingDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_zone_transfer_incoming.DNSZoneTransferIncomingDataSourceModel)(nil)
	schema := dns_zone_transfer_incoming.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
