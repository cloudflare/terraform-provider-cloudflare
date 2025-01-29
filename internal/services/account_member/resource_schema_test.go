// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/account_member"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestAccountMemberModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account_member.AccountMemberModel)(nil)
	schema := account_member.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
