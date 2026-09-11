// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending_subdomain_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_sending_subdomain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailSendingSubdomainModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_sending_subdomain.EmailSendingSubdomainModel)(nil)
	schema := email_sending_subdomain.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
