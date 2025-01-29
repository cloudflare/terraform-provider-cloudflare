// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_token_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/account_token"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestAccountTokensDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account_token.AccountTokensDataSourceModel)(nil)
	schema := account_token.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
