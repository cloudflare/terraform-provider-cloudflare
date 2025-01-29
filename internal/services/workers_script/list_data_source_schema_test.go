// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/workers_script"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestWorkersScriptsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_script.WorkersScriptsDataSourceModel)(nil)
	schema := workers_script.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
