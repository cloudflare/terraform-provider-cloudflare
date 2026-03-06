package v500

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TestTransformV4toV500(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name          string
		source        SourceV4AccountMemberModel
		expectedEmail string
		expectedRoles []string
		expectError   bool
	}{
		{
			name: "basic transformation",
			source: SourceV4AccountMemberModel{
				ID:           types.StringValue("test-member-id"),
				AccountID:    types.StringValue("test-account-id"),
				EmailAddress: types.StringValue("user@example.com"),
				RoleIDs:      mustCreateStringSet(ctx, []string{"role-1", "role-2"}),
				Status:       types.StringValue("accepted"),
			},
			expectedEmail: "user@example.com",
			expectedRoles: []string{"role-1", "role-2"},
			expectError:   false,
		},
		{
			name: "single role",
			source: SourceV4AccountMemberModel{
				ID:           types.StringValue("test-member-id-2"),
				AccountID:    types.StringValue("test-account-id-2"),
				EmailAddress: types.StringValue("admin@example.com"),
				RoleIDs:      mustCreateStringSet(ctx, []string{"admin-role"}),
				Status:       types.StringValue("pending"),
			},
			expectedEmail: "admin@example.com",
			expectedRoles: []string{"admin-role"},
			expectError:   false,
		},
		{
			name: "null role_ids",
			source: SourceV4AccountMemberModel{
				ID:           types.StringValue("test-member-id-3"),
				AccountID:    types.StringValue("test-account-id-3"),
				EmailAddress: types.StringValue("noone@example.com"),
				RoleIDs:      types.SetNull(types.StringType),
				Status:       types.StringValue("accepted"),
			},
			expectedEmail: "noone@example.com",
			expectedRoles: nil, // null set
			expectError:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, diags := TransformV4toV500(ctx, tc.source)

			if tc.expectError {
				if !diags.HasError() {
					t.Errorf("expected error but got none")
				}
				return
			}

			if diags.HasError() {
				t.Errorf("unexpected errors: %v", diags.Errors())
				return
			}

			if result == nil {
				t.Errorf("expected non-nil result")
				return
			}

			// Verify email transformation (email_address → email)
			if result.Email.ValueString() != tc.expectedEmail {
				t.Errorf("email: got %q, want %q", result.Email.ValueString(), tc.expectedEmail)
			}

			// Verify ID preserved
			if result.ID.ValueString() != tc.source.ID.ValueString() {
				t.Errorf("ID: got %q, want %q", result.ID.ValueString(), tc.source.ID.ValueString())
			}

			// Verify AccountID preserved
			if result.AccountID.ValueString() != tc.source.AccountID.ValueString() {
				t.Errorf("AccountID: got %q, want %q", result.AccountID.ValueString(), tc.source.AccountID.ValueString())
			}

			// Verify Status preserved
			if result.Status.ValueString() != tc.source.Status.ValueString() {
				t.Errorf("Status: got %q, want %q", result.Status.ValueString(), tc.source.Status.ValueString())
			}

			// Verify roles transformation (role_ids → roles)
			if tc.expectedRoles == nil {
				if !result.Roles.IsNull() {
					t.Errorf("Roles: expected null, got %v", result.Roles)
				}
			} else {
				if result.Roles.IsNull() {
					t.Errorf("Roles: expected non-null, got null")
				} else {
					var resultRoles []types.String
					diags := result.Roles.ElementsAs(ctx, &resultRoles, false)
					if diags.HasError() {
						t.Errorf("failed to extract roles: %v", diags.Errors())
					} else {
						if len(resultRoles) != len(tc.expectedRoles) {
							t.Errorf("Roles count: got %d, want %d", len(resultRoles), len(tc.expectedRoles))
						}
						// Note: Sets are unordered, so we check membership
						gotRoles := make(map[string]bool)
						for _, r := range resultRoles {
							gotRoles[r.ValueString()] = true
						}
						for _, expected := range tc.expectedRoles {
							if !gotRoles[expected] {
								t.Errorf("Roles: missing expected role %q", expected)
							}
						}
					}
				}
			}

			// Verify policies is null (not in v4)
			if !result.Policies.IsNull() {
				t.Errorf("Policies: expected null, got %v", result.Policies)
			}

			// Verify user is null (not in v4)
			if !result.User.IsNull() {
				t.Errorf("User: expected null, got %v", result.User)
			}
		})
	}
}

// mustCreateStringSet creates a types.Set containing the given strings.
// Panics if set creation fails.
func mustCreateStringSet(ctx context.Context, values []string) types.Set {
	elements := make([]attr.Value, len(values))
	for i, v := range values {
		elements[i] = types.StringValue(v)
	}
	set, diags := basetypes.NewSetValue(types.StringType, elements)
	if diags.HasError() {
		panic("failed to create set: " + diags.Errors()[0].Summary())
	}
	return set
}
