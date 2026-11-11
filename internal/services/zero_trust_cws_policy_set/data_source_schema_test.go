// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_policy_set_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_cws_policy_set"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustCWSPolicySetDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_cws_policy_set.ZeroTrustCWSPolicySetDataSourceModel)(nil)
	schema := zero_trust_cws_policy_set.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
