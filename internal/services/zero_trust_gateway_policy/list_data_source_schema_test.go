// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustGatewayPoliciesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_gateway_policy.ZeroTrustGatewayPoliciesDataSourceModel)(nil)
	schema := zero_trust_gateway_policy.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
