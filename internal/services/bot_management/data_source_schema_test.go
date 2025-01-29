// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/bot_management"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestBotManagementDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*bot_management.BotManagementDataSourceModel)(nil)
	schema := bot_management.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
