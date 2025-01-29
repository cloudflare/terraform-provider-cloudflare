// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_page_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_custom_page"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustAccessCustomPagesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_access_custom_page.ZeroTrustAccessCustomPagesDataSourceModel)(nil)
	schema := zero_trust_access_custom_page.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
