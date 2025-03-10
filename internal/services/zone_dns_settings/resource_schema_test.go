// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dns_settings_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_dns_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZoneDNSSettingsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zone_dns_settings.ZoneDNSSettingsModel)(nil)
	schema := zone_dns_settings.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
