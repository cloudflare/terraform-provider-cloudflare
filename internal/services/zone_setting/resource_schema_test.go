// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_setting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZoneSettingModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zone_setting.ZoneSettingModel)(nil)
	schema := zone_setting.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
