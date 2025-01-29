// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/snippet_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestSnippetRulesListDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*snippet_rules.SnippetRulesListDataSourceModel)(nil)
	schema := snippet_rules.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
