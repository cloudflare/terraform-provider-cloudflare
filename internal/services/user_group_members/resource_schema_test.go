// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_group_members_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/user_group_members"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestUserGroupMembersModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*user_group_members.UserGroupMembersModel)(nil)
	schema := user_group_members.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
