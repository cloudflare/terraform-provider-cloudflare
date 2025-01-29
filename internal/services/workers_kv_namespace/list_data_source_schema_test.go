// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/workers_kv_namespace"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestWorkersKVNamespacesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_kv_namespace.WorkersKVNamespacesDataSourceModel)(nil)
	schema := workers_kv_namespace.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
