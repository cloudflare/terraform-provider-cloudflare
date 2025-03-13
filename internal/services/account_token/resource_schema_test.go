// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_token_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_token"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAccountTokenModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account_token.AccountTokenModel)(nil)
	schema := account_token.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
