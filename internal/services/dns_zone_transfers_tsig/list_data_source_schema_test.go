// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_tsig_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/dns_zone_transfers_tsig"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestDNSZoneTransfersTSIGsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_zone_transfers_tsig.DNSZoneTransfersTSIGsDataSourceModel)(nil)
	schema := dns_zone_transfers_tsig.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
