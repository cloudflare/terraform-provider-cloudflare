// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_sippy_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_sippy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestR2BucketSippyDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*r2_bucket_sippy.R2BucketSippyDataSourceModel)(nil)
	schema := r2_bucket_sippy.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
