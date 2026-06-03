// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oauth_client_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/oauth_client"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestOAuthClientModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*oauth_client.OAuthClientModel)(nil)
	schema := oauth_client.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
