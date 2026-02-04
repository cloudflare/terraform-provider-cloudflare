// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_token_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ai_search_token"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAISearchTokenModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*ai_search_token.AISearchTokenModel)(nil)
	schema := ai_search_token.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
