// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/snippets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSnippetsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*snippets.SnippetsModel)(nil)
	schema := snippets.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
