// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname_fallback_origin_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_hostname_fallback_origin"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCustomHostnameFallbackOriginDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_hostname_fallback_origin.CustomHostnameFallbackOriginDataSourceModel)(nil)
	schema := custom_hostname_fallback_origin.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
