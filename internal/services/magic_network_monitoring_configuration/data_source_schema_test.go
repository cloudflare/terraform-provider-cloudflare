// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_configuration_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/magic_network_monitoring_configuration"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestMagicNetworkMonitoringConfigurationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_network_monitoring_configuration.MagicNetworkMonitoringConfigurationDataSourceModel)(nil)
	schema := magic_network_monitoring_configuration.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
