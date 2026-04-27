package user_group_members

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/iam"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// removeAllMembers removes all members from a user group by calling PUT with
// an empty array. This is a single efficient API call that removes all members
// regardless of count. The user group itself is not affected.
//
// Used by the Delete implementation to support:
//
//	terraform destroy -target cloudflare_user_group_members.xxx
//
// This matches the pattern used by cloud_connector_rules (PR #5559).
func removeAllMembers(ctx context.Context, r *UserGroupMembersResource, accountID, userGroupID string, resp *resource.DeleteResponse) {
	_, err := r.client.IAM.UserGroups.Members.Update(
		ctx,
		userGroupID,
		iam.UserGroupMemberUpdateParams{
			AccountID: cloudflare.F(accountID),
		},
		option.WithRequestBody("application/json", []byte(`[]`)),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
}

// normalizeMembers ensures the members field is always an empty list rather than nil
// when the API returns no members. This prevents null-vs-empty-list drift where:
//   - User config: members = []
//   - API response: "result": null (or missing)
//   - Without fix: state = null, plan shows "+ members = []" → persistent drift
//   - With fix: state = [], plan shows no drift
//
// Must be called after every unmarshal in Create, Update, and Read.
func normalizeMembers(m *UserGroupMembersModel) {
	if m == nil {
		return
	}
	if m.Members == nil {
		empty := []*UserGroupMembersMembersModel{}
		m.Members = &empty
	}
}

// unmarshalMembersCustom wraps the standard unmarshal with normalization to handle
// the null-vs-empty-list case in API responses.
//
// Used by the Read method to ensure consistent state representation.
func unmarshalMembersCustom(data []byte, configuredModel *UserGroupMembersModel) (*UserGroupMembersModel, error) {
	env := UserGroupMembersResultEnvelope{configuredModel.Members}
	if err := apijson.Unmarshal(data, &env); err != nil {
		return nil, err
	}

	configuredModel.Members = env.Result
	normalizeMembers(configuredModel)
	return configuredModel, nil
}
