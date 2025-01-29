// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_rule_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/magic_network_monitoring_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestMagicNetworkMonitoringRuleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_network_monitoring_rule.MagicNetworkMonitoringRuleDataSourceModel)(nil)
	schema := magic_network_monitoring_rule.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
