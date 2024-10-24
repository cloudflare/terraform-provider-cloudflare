// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package call_app_turn_key_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/call_app_turn_key"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCallAppTURNKeyDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*call_app_turn_key.CallAppTURNKeyDataSourceModel)(nil)
	schema := call_app_turn_key.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
