// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_policy_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_allow_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailSecurityAllowPoliciesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_security_allow_policy.EmailSecurityAllowPoliciesDataSourceModel)(nil)
	schema := email_security_allow_policy.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
