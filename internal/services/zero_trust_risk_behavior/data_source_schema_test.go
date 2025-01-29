// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_behavior_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_risk_behavior"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustRiskBehaviorDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_risk_behavior.ZeroTrustRiskBehaviorDataSourceModel)(nil)
	schema := zero_trust_risk_behavior.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
