// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_infrastructure_target_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_access_infrastructure_target"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustAccessInfrastructureTargetDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_access_infrastructure_target.ZeroTrustAccessInfrastructureTargetDataSourceModel)(nil)
	schema := zero_trust_access_infrastructure_target.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
