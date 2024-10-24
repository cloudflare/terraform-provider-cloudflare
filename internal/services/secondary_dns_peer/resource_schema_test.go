// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_peer_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/secondary_dns_peer"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSecondaryDNSPeerModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*secondary_dns_peer.SecondaryDNSPeerModel)(nil)
	schema := secondary_dns_peer.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
