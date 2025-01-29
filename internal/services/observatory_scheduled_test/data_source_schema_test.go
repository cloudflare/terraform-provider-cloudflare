// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/observatory_scheduled_test"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestObservatoryScheduledTestDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*observatory_scheduled_test.ObservatoryScheduledTestDataSourceModel)(nil)
	schema := observatory_scheduled_test.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
