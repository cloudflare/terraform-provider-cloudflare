// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/pages_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestPagesDomainDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pages_domain.PagesDomainDataSourceModel)(nil)
	schema := pages_domain.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
