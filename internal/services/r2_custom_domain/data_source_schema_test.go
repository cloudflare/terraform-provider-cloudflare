// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_custom_domain_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_custom_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestR2CustomDomainDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*r2_custom_domain.R2CustomDomainDataSourceModel)(nil)
	schema := r2_custom_domain.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
