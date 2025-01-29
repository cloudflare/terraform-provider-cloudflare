// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/api_token"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestAPITokensDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_token.APITokensDataSourceModel)(nil)
	schema := api_token.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
