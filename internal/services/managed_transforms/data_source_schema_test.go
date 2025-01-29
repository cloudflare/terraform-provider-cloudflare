// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/managed_transforms"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestManagedTransformsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*managed_transforms.ManagedTransformsDataSourceModel)(nil)
	schema := managed_transforms.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
