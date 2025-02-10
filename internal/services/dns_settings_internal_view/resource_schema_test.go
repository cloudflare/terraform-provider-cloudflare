// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_settings_internal_view_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_settings_internal_view"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSSettingsInternalViewModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_settings_internal_view.DNSSettingsInternalViewModel)(nil)
	schema := dns_settings_internal_view.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
