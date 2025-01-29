// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_gateway_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustGatewaySettingsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_gateway_settings.ZeroTrustGatewaySettingsDataSourceModel)(nil)
	schema := zero_trust_gateway_settings.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
