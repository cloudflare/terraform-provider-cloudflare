// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package total_tls_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/total_tls"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestTotalTLSDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*total_tls.TotalTLSDataSourceModel)(nil)
	schema := total_tls.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
