// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_auth_dmarc_report_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_auth_dmarc_report"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestEmailAuthDMARCReportModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_auth_dmarc_report.EmailAuthDMARCReportModel)(nil)
	schema := email_auth_dmarc_report.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
