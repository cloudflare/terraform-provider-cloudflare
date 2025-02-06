// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_settings_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSSettingsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_settings.DNSSettingsDataSourceModel)(nil)
	schema := dns_settings.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
