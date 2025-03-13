// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_tsig_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfers_tsig"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSZoneTransfersTSIGDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*dns_zone_transfers_tsig.DNSZoneTransfersTSIGDataSourceModel)(nil)
  schema := dns_zone_transfers_tsig.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
