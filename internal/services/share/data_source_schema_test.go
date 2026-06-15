// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/share"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestShareDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*share.ShareDataSourceModel)(nil)
	schema := share.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
