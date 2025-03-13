// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning_expression_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/content_scanning_expression"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestContentScanningExpressionsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*content_scanning_expression.ContentScanningExpressionsDataSourceModel)(nil)
  schema := content_scanning_expression.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
