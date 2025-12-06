package ruleset_rule_test

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

// TODO Figreout what is Stainless

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestRulesetRuleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*ruleset_rule.RulesetRuleDataSourceModel)(nil)
	schema := ruleset_rule.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
