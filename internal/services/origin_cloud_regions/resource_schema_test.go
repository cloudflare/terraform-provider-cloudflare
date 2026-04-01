package origin_cloud_regions_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/origin_cloud_regions"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestOriginCloudRegionsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*origin_cloud_regions.OriginCloudRegionsModel)(nil)
	schema := origin_cloud_regions.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
