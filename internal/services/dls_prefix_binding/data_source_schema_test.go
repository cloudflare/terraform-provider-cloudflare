// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dls_prefix_binding_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dls_prefix_binding"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDLSPrefixBindingDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dls_prefix_binding.DLSPrefixBindingDataSourceModel)(nil)
	schema := dls_prefix_binding.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
