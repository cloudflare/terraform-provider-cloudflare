// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_dataset_field_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/logpush_dataset_field"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestLogpushDatasetFieldDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*logpush_dataset_field.LogpushDatasetFieldDataSourceModel)(nil)
	schema := logpush_dataset_field.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
