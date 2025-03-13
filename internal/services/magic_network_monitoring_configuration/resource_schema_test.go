// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_configuration_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_network_monitoring_configuration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestMagicNetworkMonitoringConfigurationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_network_monitoring_configuration.MagicNetworkMonitoringConfigurationModel)(nil)
	schema := magic_network_monitoring_configuration.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
