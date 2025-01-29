// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_proxy_endpoint_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_proxy_endpoint"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustGatewayProxyEndpointDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_gateway_proxy_endpoint.ZeroTrustGatewayProxyEndpointDataSourceModel)(nil)
	schema := zero_trust_gateway_proxy_endpoint.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
