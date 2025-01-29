// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_app_types_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_gateway_app_types"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustGatewayAppTypesListDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_gateway_app_types.ZeroTrustGatewayAppTypesListDataSourceModel)(nil)
	schema := zero_trust_gateway_app_types.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
