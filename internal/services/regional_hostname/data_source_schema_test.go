// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/regional_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestRegionalHostnameDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*regional_hostname.RegionalHostnameDataSourceModel)(nil)
	schema := regional_hostname.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
