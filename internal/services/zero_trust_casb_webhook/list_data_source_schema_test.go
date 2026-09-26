// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_webhook_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_casb_webhook"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustCasbWebhooksDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_casb_webhook.ZeroTrustCasbWebhooksDataSourceModel)(nil)
	schema := zero_trust_casb_webhook.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
