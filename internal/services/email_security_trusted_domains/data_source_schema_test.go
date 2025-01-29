// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_trusted_domains_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/email_security_trusted_domains"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestEmailSecurityTrustedDomainsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_security_trusted_domains.EmailSecurityTrustedDomainsDataSourceModel)(nil)
	schema := email_security_trusted_domains.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
