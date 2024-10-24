// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_pattern_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_allow_pattern"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailSecurityAllowPatternsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_security_allow_pattern.EmailSecurityAllowPatternsDataSourceModel)(nil)
	schema := email_security_allow_pattern.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
