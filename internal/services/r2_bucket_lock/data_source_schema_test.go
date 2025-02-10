// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_lock_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_lock"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestR2BucketLockDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*r2_bucket_lock.R2BucketLockDataSourceModel)(nil)
	schema := r2_bucket_lock.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
