// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_rule_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_network_monitoring_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMagicNetworkMonitoringRuleModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_network_monitoring_rule.MagicNetworkMonitoringRuleModel)(nil)
	schema := magic_network_monitoring_rule.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
