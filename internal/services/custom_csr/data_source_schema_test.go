// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_csr_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_csr"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCustomCsrDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_csr.CustomCsrDataSourceModel)(nil)
	schema := custom_csr.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
