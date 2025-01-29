// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/workers_custom_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestWorkersCustomDomainModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_custom_domain.WorkersCustomDomainModel)(nil)
	schema := workers_custom_domain.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
