// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hostname_tls_setting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestHostnameTLSSettingDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*hostname_tls_setting.HostnameTLSSettingDataSourceModel)(nil)
	schema := hostname_tls_setting.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
