// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/web_analytics_site"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestWebAnalyticsSiteModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*web_analytics_site.WebAnalyticsSiteModel)(nil)
	schema := web_analytics_site.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
