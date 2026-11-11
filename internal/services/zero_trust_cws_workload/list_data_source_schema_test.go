// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_workload_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_cws_workload"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustCWSWorkloadsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_cws_workload.ZeroTrustCWSWorkloadsDataSourceModel)(nil)
	schema := zero_trust_cws_workload.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
