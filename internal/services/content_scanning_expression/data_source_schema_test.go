// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning_expression_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/content_scanning_expression"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestContentScanningExpressionDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*content_scanning_expression.ContentScanningExpressionDataSourceModel)(nil)
	schema := content_scanning_expression.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}