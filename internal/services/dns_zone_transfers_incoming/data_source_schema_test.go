// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_incoming_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfers_incoming"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSZoneTransfersIncomingDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*dns_zone_transfers_incoming.DNSZoneTransfersIncomingDataSourceModel)(nil)
  schema := dns_zone_transfers_incoming.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
