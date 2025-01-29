// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_key_configuration_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_key_configuration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustAccessKeyConfigurationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_access_key_configuration.ZeroTrustAccessKeyConfigurationDataSourceModel)(nil)
	schema := zero_trust_access_key_configuration.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
