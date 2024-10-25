// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestAccessRulesDataSourceModelSchemaParity(t *testing.T) {
	t.Skip("need investigation: currently broken")
	t.Parallel()
	model := (*access_rule.AccessRulesDataSourceModel)(nil)
	schema := access_rule.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}