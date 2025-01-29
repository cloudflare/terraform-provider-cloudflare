// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning_expression_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/content_scanning_expression"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestContentScanningExpressionModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*content_scanning_expression.ContentScanningExpressionModel)(nil)
	schema := content_scanning_expression.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
