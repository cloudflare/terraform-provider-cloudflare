// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_cookies_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_shield_cookies"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPageShieldCookiesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*page_shield_cookies.PageShieldCookiesDataSourceModel)(nil)
	schema := page_shield_cookies.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
