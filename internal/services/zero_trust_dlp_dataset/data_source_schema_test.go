// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_dataset_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_dataset"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDLPDatasetDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*zero_trust_dlp_dataset.ZeroTrustDLPDatasetDataSourceModel)(nil)
  schema := zero_trust_dlp_dataset.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
