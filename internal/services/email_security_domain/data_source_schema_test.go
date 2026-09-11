// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_domain_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailSecurityDomainDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_security_domain.EmailSecurityDomainDataSourceModel)(nil)
	schema := email_security_domain.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
