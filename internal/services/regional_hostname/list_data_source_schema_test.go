// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/regional_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestRegionalHostnamesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*regional_hostname.RegionalHostnamesDataSourceModel)(nil)
	schema := regional_hostname.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
