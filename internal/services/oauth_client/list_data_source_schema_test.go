// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oauth_client_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/oauth_client"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestOAuthClientsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*oauth_client.OAuthClientsDataSourceModel)(nil)
	schema := oauth_client.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
