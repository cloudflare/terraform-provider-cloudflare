// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_subdomain_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_subdomain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersSubdomainModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_subdomain.WorkersSubdomainModel)(nil)
	schema := workers_subdomain.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
