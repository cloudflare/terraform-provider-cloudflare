// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image_variant_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/image_variant"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestImageVariantDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*image_variant.ImageVariantDataSourceModel)(nil)
  schema := image_variant.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
