// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_scripts_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/page_shield_scripts"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestPageShieldScriptsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*page_shield_scripts.PageShieldScriptsDataSourceModel)(nil)
	schema := page_shield_scripts.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
