// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_trusted_domains_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_trusted_domains"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailSecurityTrustedDomainsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_security_trusted_domains.EmailSecurityTrustedDomainsModel)(nil)
	schema := email_security_trusted_domains.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
