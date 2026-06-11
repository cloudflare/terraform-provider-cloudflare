// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share_recipient_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/share_recipient"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestShareRecipientsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*share_recipient.ShareRecipientsDataSourceModel)(nil)
	schema := share_recipient.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
