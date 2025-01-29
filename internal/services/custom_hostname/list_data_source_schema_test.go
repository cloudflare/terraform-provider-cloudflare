// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/custom_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestCustomHostnamesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_hostname.CustomHostnamesDataSourceModel)(nil)
	schema := custom_hostname.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
