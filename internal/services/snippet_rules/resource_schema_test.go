// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/snippet_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSnippetRulesModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*snippet_rules.SnippetRulesModel)(nil)
	schema := snippet_rules.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
