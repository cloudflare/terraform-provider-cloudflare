// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package call_app_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/call_app"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCallAppDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*call_app.CallAppDataSourceModel)(nil)
	schema := call_app.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
