// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_settings_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAuthenticatedOriginPullsSettingsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*authenticated_origin_pulls_settings.AuthenticatedOriginPullsSettingsModel)(nil)
	schema := authenticated_origin_pulls_settings.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
