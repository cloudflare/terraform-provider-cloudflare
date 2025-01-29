// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_route_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/workers_route"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestWorkersRouteModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_route.WorkersRouteModel)(nil)
	schema := workers_route.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
