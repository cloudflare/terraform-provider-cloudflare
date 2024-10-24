// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/image"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestImageDataSourceModelSchemaParity(t *testing.T) {
	t.Skip("need investigation: currently broken")
	t.Parallel()
	model := (*image.ImageDataSourceModel)(nil)
	schema := image.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
