// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_cloud_region_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/origin_cloud_region"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestOriginCloudRegionsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*origin_cloud_region.OriginCloudRegionsDataSourceModel)(nil)
	schema := origin_cloud_region.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
