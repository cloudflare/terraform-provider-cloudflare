// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/firewall_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestFirewallRuleModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*firewall_rule.FirewallRuleModel)(nil)
	schema := firewall_rule.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
