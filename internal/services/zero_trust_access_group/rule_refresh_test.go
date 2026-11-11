package zero_trust_access_group

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"
)

const groupInclude = `[{"email_domain":{"domain":"example.com"}}]`

func groupRules(field, rules string) string {
	if field == "include" {
		return `"include":` + rules
	}
	return `"include":` + groupInclude + `,"` + field + `":` + rules
}

func groupState(t *testing.T, ctx context.Context, field, rules string) tfsdk.State {
	t.Helper()
	var model ZeroTrustAccessGroupModel
	body := `{"id":"group-id","name":"Group",` + groupRules(field, rules) + `}`
	require.NoError(t, apijson.Unmarshal([]byte(body), &model))
	model.AccountID = types.StringValue("account-id")
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	diags := state.Set(ctx, &model)
	require.False(t, diags.HasError(), "%v", diags)
	return state
}

func groupRuleList(t *testing.T, ctx context.Context, state tfsdk.State, field string) any {
	t.Helper()
	var model ZeroTrustAccessGroupModel
	diags := state.Get(ctx, &model)
	require.False(t, diags.HasError(), "%v", diags)
	switch field {
	case "include":
		return model.Include
	case "exclude":
		return model.Exclude
	default:
		return model.Require
	}
}

// The API returns a group's rules in its own order, which can differ from the
// order of the list in state. Refresh must store the API's rules as they are:
// a rule at a new position is a reorder, not a selector merged into the rule
// that used to sit there.
func TestAccessGroupRefreshDoesNotMergeSelectorsAcrossRules(t *testing.T) {
	tests := []struct {
		name  string
		state string
		api   string
		want  string
	}{
		{
			name:  "ip and device_posture",
			state: `[{"ip":{"ip":"192.0.2.10/32"}},{"device_posture":{"integration_uid":"posture-id"}}]`,
			api:   `[{"device_posture":{"integration_uid":"posture-id"}},{"ip":{"ip":"192.0.2.10/32"}}]`,
		},
		{
			name:  "email and group",
			state: `[{"email":{"email":"user@example.com"}},{"group":{"id":"other-group-id"}}]`,
			api:   `[{"group":{"id":"other-group-id"}},{"email":{"email":"user@example.com"}}]`,
		},
		{
			name:  "everyone and certificate",
			state: `[{"everyone":{}},{"certificate":{}}]`,
			api:   `[{"certificate":{}},{"everyone":{}}]`,
		},
		{
			name:  "service token and any valid service token",
			state: `[{"service_token":{"token_id":"token-id"}},{"any_valid_service_token":{}}]`,
			api:   `[{"any_valid_service_token":{}},{"service_token":{"token_id":"token-id"}}]`,
		},
		{
			name:  "ip and user risk score",
			state: `[{"ip":{"ip":"192.0.2.10/32"}},{"user_risk_score":{"user_risk_score":["high"]}}]`,
			api:   `[{"user_risk_score":{"user_risk_score":["high"]}},{"ip":{"ip":"192.0.2.10/32"}}]`,
		},
		{
			name:  "email and cloudflare account member",
			state: `[{"email":{"email":"user@example.com"}},{"cloudflare_account_member":{"account_id":"member-account-id"}}]`,
			api:   `[{"cloudflare_account_member":{"account_id":"member-account-id"}},{"email":{"email":"user@example.com"}}]`,
		},
		{
			name:  "a different rule at a position",
			state: `[{"everyone":{}},{"ip":{"ip":"192.0.2.10/32"}}]`,
			api:   `[{"certificate":{}},{"ip":{"ip":"192.0.2.10/32"}}]`,
		},
		{
			name:  "state holds a merged rule",
			state: `[{"ip":{"ip":"192.0.2.10/32"},"everyone":{}}]`,
			api:   `[{"ip":{"ip":"192.0.2.10/32"}}]`,
		},
		{
			name:  "api returns an empty selector",
			state: `[{"ip":{"ip":"192.0.2.10/32"}},{"common_name":{"common_name":"device.example.com"}}]`,
			api:   `[{"common_name":{"common_name":"device.example.com"}},{"ip":{"ip":"192.0.2.10/32"},"common_name":{"common_name":null}}]`,
			want:  `[{"common_name":{"common_name":"device.example.com"}},{"ip":{"ip":"192.0.2.10/32"}}]`,
		},
	}
	for _, field := range []string{"include", "exclude", "require"} {
		for _, tt := range tests {
			t.Run(field+"/"+tt.name, func(t *testing.T) {
				ctx := context.Background()
				state := readGroup(t, ctx, groupState(t, ctx, field, tt.state), field, tt.api)
				want := tt.want
				if want == "" {
					want = tt.api
				}
				expected := groupRuleList(t, ctx, groupState(t, ctx, field, want), field)
				require.Equal(t, expected, groupRuleList(t, ctx, state, field))
			})
		}
	}
}

// An empty list from the API still refreshes to null when state has no rules.
func TestAccessGroupRefreshKeepsNullRules(t *testing.T) {
	for _, field := range []string{"exclude", "require"} {
		t.Run(field, func(t *testing.T) {
			ctx := context.Background()
			state := readGroup(t, ctx, groupState(t, ctx, "include", groupInclude), field, `[]`)
			require.Nil(t, groupRuleList(t, ctx, state, field))
		})
	}
}

// readGroup refreshes state against a mock API that returns the group with the
// given rules in field.
func readGroup(t *testing.T, ctx context.Context, state tfsdk.State, field, rules string) tfsdk.State {
	t.Helper()
	group := map[string]any{
		"id": "group-id", "name": "Group",
		"created_at": "2026-01-01T00:00:00Z", "updated_at": "2026-01-01T00:00:00Z",
		"include": json.RawMessage(groupInclude), "exclude": []any{}, "require": []any{},
	}
	group[field] = json.RawMessage(rules)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet || req.URL.Path != "/accounts/account-id/access/groups/group-id" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
			http.Error(w, "unexpected request", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true, "errors": []any{}, "messages": []any{}, "result": group,
		})
	}))
	defer server.Close()
	r := &ZeroTrustAccessGroupResource{client: cloudflare.NewClient(option.WithBaseURL(server.URL), option.WithAPIToken("test-token"))}
	response := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, response)
	require.False(t, response.Diagnostics.HasError(), "%v", response.Diagnostics)
	return response.State
}
