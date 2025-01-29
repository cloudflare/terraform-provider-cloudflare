// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_webhook_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/stream_webhook"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestStreamWebhookDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*stream_webhook.StreamWebhookDataSourceModel)(nil)
	schema := stream_webhook.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
