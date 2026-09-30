// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_integration_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_casb_integration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustCasbIntegrationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_casb_integration.ZeroTrustCasbIntegrationDataSourceModel)(nil)
	schema := zero_trust_casb_integration.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
