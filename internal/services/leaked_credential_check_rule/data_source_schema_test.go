// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_check_rule_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/leaked_credential_check_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestLeakedCredentialCheckRuleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*leaked_credential_check_rule.LeakedCredentialCheckRuleDataSourceModel)(nil)
	schema := leaked_credential_check_rule.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
