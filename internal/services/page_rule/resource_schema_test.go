// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestPageRuleModelSchemaParity(t *testing.T) {
	t.Skip("too much custom code to have model parity")
	t.Parallel()
	model := (*page_rule.PageRuleModel)(nil)
	schema := page_rule.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
