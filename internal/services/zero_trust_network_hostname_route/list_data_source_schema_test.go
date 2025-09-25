// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_network_hostname_route_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_network_hostname_route"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustNetworkHostnameRoutesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_network_hostname_route.ZeroTrustNetworkHostnameRoutesDataSourceModel)(nil)
	schema := zero_trust_network_hostname_route.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
