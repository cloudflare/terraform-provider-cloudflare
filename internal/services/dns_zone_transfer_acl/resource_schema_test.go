// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfer_acl_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfer_acl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSZoneTransferACLModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_zone_transfer_acl.DNSZoneTransferACLModel)(nil)
	schema := dns_zone_transfer_acl.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
