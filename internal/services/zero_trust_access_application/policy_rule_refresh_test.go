package zero_trust_access_application

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"
)

const inlinePolicyInclude = `[{"email_domain":{"domain":"example.com"}}]`

func inlinePolicyRules(field, rules string) string {
	if field == "include" {
		return `"include":` + rules
	}
	return `"include":` + inlinePolicyInclude + `,"` + field + `":` + rules
}

func inlinePolicyState(t *testing.T, ctx context.Context, field, rules string) tfsdk.State {
	t.Helper()
	var model ZeroTrustAccessApplicationModel
	body := `{"id":"app-id","type":"self_hosted","name":"App","domain":"app.example.com",` +
		`"policies":[{"precedence":1,"decision":"allow","name":"Inline",` + inlinePolicyRules(field, rules) + `}]}`
	require.NoError(t, apijson.Unmarshal([]byte(body), &model))
	model.AccountID = types.StringValue("account-id")
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	diags := state.Set(ctx, &model)
	require.False(t, diags.HasError(), "%v", diags)
	return state
}

func inlinePolicyRuleSet(t *testing.T, ctx context.Context, state tfsdk.State, field string) attr.Value {
	t.Helper()
	var model ZeroTrustAccessApplicationModel
	diags := state.Get(ctx, &model)
	require.False(t, diags.HasError(), "%v", diags)
	require.NotNil(t, model.Policies)
	require.Len(t, *model.Policies, 1)
	policy := (*model.Policies)[0]
	switch field {
	case "include":
		return policy.Include
	case "exclude":
		return policy.Exclude
	default:
		return policy.Require
	}
}

// The API returns the rules of an inline policy in its own order, which can
// differ from the order of the set in state. Refresh must not carry a selector
// from one rule into another, either as an empty phantom object or as a merged
// selector.
func TestAccessApplicationRefreshDoesNotMergeInlinePolicySelectorsAcrossRules(t *testing.T) {
	tests := []struct {
		name  string
		state string
		api   string
	}{
		{
			name:  "ip and device_posture",
			state: `[{"ip":{"ip":"192.0.2.10/32"}},{"device_posture":{"integration_uid":"posture-id"}}]`,
			api:   `[{"device_posture":{"integration_uid":"posture-id"}},{"ip":{"ip":"192.0.2.10/32"}}]`,
		},
		{
			name:  "email and oidc",
			state: `[{"email":{"email":"user@example.com"}},{"oidc":{"claim_name":"groups","claim_value":"admins","identity_provider_id":"idp-id"}}]`,
			api:   `[{"oidc":{"claim_name":"groups","claim_value":"admins","identity_provider_id":"idp-id"}},{"email":{"email":"user@example.com"}}]`,
		},
		{
			name:  "group and service token",
			state: `[{"group":{"id":"group-id"}},{"any_valid_service_token":{}}]`,
			api:   `[{"any_valid_service_token":{}},{"group":{"id":"group-id"}}]`,
		},
		{
			name:  "everyone and certificate",
			state: `[{"everyone":{}},{"certificate":{}}]`,
			api:   `[{"certificate":{}},{"everyone":{}}]`,
		},
		{
			name:  "state holds a phantom selector",
			state: `[{"ip":{"ip":"192.0.2.10/32"},"device_posture":{"integration_uid":null}}]`,
			api:   `[{"ip":{"ip":"192.0.2.10/32"}}]`,
		},
	}
	for _, field := range []string{"include", "exclude", "require"} {
		for _, tt := range tests {
			t.Run(field+"/"+tt.name, func(t *testing.T) {
				ctx := context.Background()
				policy := map[string]any{
					"id": "policy-id", "precedence": 1, "decision": "allow", "name": "Inline",
					"created_at": "2026-01-01T00:00:00Z", "updated_at": "2026-01-01T00:00:00Z",
					"session_duration": "24h", "include": json.RawMessage(inlinePolicyInclude),
					"exclude": []any{}, "require": []any{},
				}
				policy[field] = json.RawMessage(tt.api)
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					if req.Method != http.MethodGet || req.URL.Path != "/accounts/account-id/access/apps/app-id" {
						t.Errorf("unexpected request: %s %s", req.Method, req.URL)
						http.Error(w, "unexpected request", http.StatusBadRequest)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					_ = json.NewEncoder(w).Encode(map[string]any{
						"success": true,
						"result": map[string]any{
							"id": "app-id", "type": "self_hosted", "name": "App", "domain": "app.example.com",
							"aud": "aud-tag", "created_at": "2026-01-01T00:00:00Z", "updated_at": "2026-01-01T00:00:00Z",
							"session_duration": "24h", "policies": []any{policy},
						},
					})
				}))
				defer server.Close()
				r := &ZeroTrustAccessApplicationResource{client: cloudflare.NewClient(option.WithBaseURL(server.URL), option.WithAPIToken("test-token"))}
				state := inlinePolicyState(t, ctx, field, tt.state)
				response := &resource.ReadResponse{State: state}
				r.Read(ctx, resource.ReadRequest{State: state}, response)
				require.False(t, response.Diagnostics.HasError(), "%v", response.Diagnostics)
				expected := inlinePolicyRuleSet(t, ctx, inlinePolicyState(t, ctx, field, tt.api), field)
				actual := inlinePolicyRuleSet(t, ctx, response.State, field)
				require.True(t, expected.Equal(actual), "expected %s, got %s", expected, actual)
			})
		}
	}
}
