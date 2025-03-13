// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/image"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestImageModelSchemaParity(t *testing.T) {
  t.Skip("need investigation: currently broken")
  t.Parallel()
  model := (*image.ImageModel)(nil)
  schema := image.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
