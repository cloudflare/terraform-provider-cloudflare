// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oauth_scope_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/oauth_scope"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestOAuthScopesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*oauth_scope.OAuthScopesDataSourceModel)(nil)
	schema := oauth_scope.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
