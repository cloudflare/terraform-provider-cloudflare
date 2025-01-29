// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/registrar_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestRegistrarDomainDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*registrar_domain.RegistrarDomainDataSourceModel)(nil)
	schema := registrar_domain.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
