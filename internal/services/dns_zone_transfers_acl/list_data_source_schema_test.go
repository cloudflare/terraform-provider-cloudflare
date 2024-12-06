// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_acl_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfers_acl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSZoneTransfersACLsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_zone_transfers_acl.DNSZoneTransfersACLsDataSourceModel)(nil)
	schema := dns_zone_transfers_acl.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
