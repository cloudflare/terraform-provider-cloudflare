// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/registrar_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestRegistrarDomainsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*registrar_domain.RegistrarDomainsDataSourceModel)(nil)
	schema := registrar_domain.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
