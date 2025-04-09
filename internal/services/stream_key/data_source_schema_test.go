// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_key_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_key"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestStreamKeyDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*stream_key.StreamKeyDataSourceModel)(nil)
  schema := stream_key.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
