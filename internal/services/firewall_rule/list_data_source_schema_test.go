// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/firewall_rule"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestFirewallRulesDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*firewall_rule.FirewallRulesDataSourceModel)(nil)
  schema := firewall_rule.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
