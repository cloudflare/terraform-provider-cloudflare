// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_managed_domain_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_managed_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestR2ManagedDomainModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*r2_managed_domain.R2ManagedDomainModel)(nil)
	schema := r2_managed_domain.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
