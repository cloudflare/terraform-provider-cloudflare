// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dex_test"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDEXTestsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_dex_test.ZeroTrustDEXTestsDataSourceModel)(nil)
	schema := zero_trust_dex_test.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
