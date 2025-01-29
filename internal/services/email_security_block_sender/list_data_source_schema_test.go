// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_block_sender_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/email_security_block_sender"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestEmailSecurityBlockSendersDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_security_block_sender.EmailSecurityBlockSendersDataSourceModel)(nil)
	schema := email_security_block_sender.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
