// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset_test

import (
  "context"
  "testing"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestRulesetDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*ruleset.RulesetDataSourceModel)(nil)
  schema := ruleset.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
