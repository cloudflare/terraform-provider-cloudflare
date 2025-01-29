// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/firewall_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestFirewallRuleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*firewall_rule.FirewallRuleDataSourceModel)(nil)
	schema := firewall_rule.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
