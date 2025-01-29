// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_acl_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/dns_zone_transfers_acl"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestDNSZoneTransfersACLModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_zone_transfers_acl.DNSZoneTransfersACLModel)(nil)
	schema := dns_zone_transfers_acl.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
