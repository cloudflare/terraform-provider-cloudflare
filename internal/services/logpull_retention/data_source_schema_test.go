// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpull_retention_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/logpull_retention"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestLogpullRetentionDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*logpull_retention.LogpullRetentionDataSourceModel)(nil)
	schema := logpull_retention.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
