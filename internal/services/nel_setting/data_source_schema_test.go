// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nel_setting_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/nel_setting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestNELSettingDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*nel_setting.NELSettingDataSourceModel)(nil)
	schema := nel_setting.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
