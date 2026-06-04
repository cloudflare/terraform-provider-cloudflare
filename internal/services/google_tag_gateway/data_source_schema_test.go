// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package google_tag_gateway_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/google_tag_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestGoogleTagGatewayDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*google_tag_gateway.GoogleTagGatewayDataSourceModel)(nil)
	schema := google_tag_gateway.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
