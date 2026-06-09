// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_app_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/flagship_app"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestFlagshipAppDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*flagship_app.FlagshipAppDataSourceModel)(nil)
	schema := flagship_app.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
