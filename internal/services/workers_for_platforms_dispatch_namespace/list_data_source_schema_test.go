// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/workers_for_platforms_dispatch_namespace"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestWorkersForPlatformsDispatchNamespacesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_for_platforms_dispatch_namespace.WorkersForPlatformsDispatchNamespacesDataSourceModel)(nil)
	schema := workers_for_platforms_dispatch_namespace.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
