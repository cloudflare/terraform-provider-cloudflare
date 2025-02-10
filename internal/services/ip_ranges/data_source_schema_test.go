// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ip_ranges_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ip_ranges"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestIPRangesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*ip_ranges.IPRangesDataSourceModel)(nil)
	schema := ip_ranges.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
