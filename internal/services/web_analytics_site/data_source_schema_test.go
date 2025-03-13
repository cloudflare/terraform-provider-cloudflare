// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/web_analytics_site"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWebAnalyticsSiteDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*web_analytics_site.WebAnalyticsSiteDataSourceModel)(nil)
  schema := web_analytics_site.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
