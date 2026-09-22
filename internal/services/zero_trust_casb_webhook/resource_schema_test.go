// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_webhook_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_casb_webhook"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustCasbWebhookModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_casb_webhook.ZeroTrustCasbWebhookModel)(nil)
	schema := zero_trust_casb_webhook.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
