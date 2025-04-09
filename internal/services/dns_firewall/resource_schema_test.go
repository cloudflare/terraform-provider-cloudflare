// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_firewall"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSFirewallModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*dns_firewall.DNSFirewallModel)(nil)
  schema := dns_firewall.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
