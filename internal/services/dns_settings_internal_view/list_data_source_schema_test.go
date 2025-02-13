// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_settings_internal_view_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_settings_internal_view"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSSettingsInternalViewsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_settings_internal_view.DNSSettingsInternalViewsDataSourceModel)(nil)
	schema := dns_settings_internal_view.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
