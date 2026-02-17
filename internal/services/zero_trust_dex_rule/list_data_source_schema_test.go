// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_rule_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dex_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDEXRulesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_dex_rule.ZeroTrustDEXRulesDataSourceModel)(nil)
	schema := zero_trust_dex_rule.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
